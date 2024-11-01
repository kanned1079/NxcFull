package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/document/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetAllDocuments(context *gin.Context) {
	lang := context.Query("lang")
	find := context.Query("find")

	resp, err := grpcClient.DocumentServiceClient.GetDocuments(sysContext.Background(), &pb.GetDocumentsRequest{
		Language: lang,
		Find:     find,
	})
	//if err != nil {
	//	log.Println(err)
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	//if resp == nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "rpc方法没有返回值",
	//	})
	//	return
	//}
	if err := failOnRpcError(err, lang); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	var docMap []map[string]interface{}

	if err := json.Unmarshal(resp.Documents, &docMap); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Unmarshal failure" + err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"doc_list": docMap,
	})
}

func HandleAddNewDocument(context *gin.Context) {
	var postData = &struct {
		Title    string `json:"title"`
		Category string `json:"category"`
		Language string `json:"language"`
		MdText   string `json:"md_text"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
	}
	// 请求rpc函数
	resp, err := grpcClient.DocumentServiceClient.AddDocument(sysContext.Background(), &pb.AddDocumentRequest{
		Title:    postData.Title,
		Category: postData.Category,
		Language: postData.Language,
		Body:     postData.MdText,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
