package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/document/proto"
	sysContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllDocuments(context *gin.Context) {
	lang := context.Query("lang")
	find := context.Query("find")

	resp, err := grpcClient.DocumentServiceClient.GetDocuments(sysContext.Background(), &pb.GetDocumentsRequest{
		Language: lang,
		Find:     find,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc方法没有返回值",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"doc_list": resp.Documents,
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
	////log.Println(postData)
	//newDoc := {
	//	Title:    postData.Title,
	//	Category: postData.Category,
	//	Language: postData.Language,
	//	Body:     postData.MdText,
	//}
	//if err := dao.Db.Model(&Document{}).Create(&newDoc).Error; err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "插入数据库失败",
	//	})
	//}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
