package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/document/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllDocuments(context *gin.Context) {
	lang := context.Query("lang")
	find := context.Query("find")

	resp, err := grpcClient.DocumentServiceClient.GetDocuments(sysContext.Background(), &pb.GetDocumentsRequest{
		Language: lang,
		Find:     find,
	})
	if err := failOnRpcError(err, lang); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	//var docMap []map[string]interface{}
	//
	//if err := json.Unmarshal(resp.Documents, &docMap); err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "Unmarshal failure" + err.Error(),
	//	})
	//}

	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"doc_list": json.RawMessage(resp.Documents),
	})
}

func HandleAddNewDocument(context *gin.Context) {
	/**
	title: string,
	    category: string,
	    language: string,
	    md_text: string,
	    sort: number,
	*/
	var postData = &struct {
		Title    string `json:"title"`
		Category string `json:"category"`
		Language string `json:"language"`
		MdText   string `json:"md_text"`
		Sort     int64  `json:"sort"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败" + err.Error(),
		})
		return
	}
	// 请求rpc函数
	resp, err := grpcClient.DocumentServiceClient.AddDocument(sysContext.Background(), &pb.AddDocumentRequest{
		Title:    postData.Title,
		Category: postData.Category,
		Language: postData.Language,
		Body:     postData.MdText,
		Sort:     postData.Sort,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
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

// HandleUpdateDocumentShow 注意這個是PATCH請求方法
func HandleUpdateDocumentShow(context *gin.Context) {
	postData := &struct {
		Id int64 `json:"id"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(postData.Id)
	resp, err := grpcClient.DocumentServiceClient.UpdateDocumentShow(sysContext.Background(), &pb.UpdateDocumentShowRequest{
		Id: postData.Id,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"success": resp.Success,
	})
}

func HandleUpdateDocument(context *gin.Context) {
	postData := &struct {
		Id       int64  `json:"id"`
		Category string `json:"category"`
		Language string `json:"language"`
		Title    string `json:"title"`
		MdText   string `json:"md_text"`
		Sort     int64  `json:"sort"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(postData)
	resp, err := grpcClient.DocumentServiceClient.UpdateDocument(sysContext.Background(), &pb.UpdateDocumentRequest{
		Id:       postData.Id,
		Category: postData.Category,
		Language: postData.Language,
		Title:    postData.Title,
		Body:     postData.MdText,
		Sort:     postData.Sort,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"success": resp.Success,
	})
}

func HandleGetAllDocumentsAdmin(context *gin.Context) {
	page, err := strconv.ParseInt(context.Query("page"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Please provide query params: page.",
		})
		return
	}
	size, err := strconv.ParseInt(context.Query("size"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Please provide query params: size.",
		})
		return
	}
	log.Println(page, size)
	resp, err := grpcClient.DocumentServiceClient.GetDocumentsAdmin(sysContext.Background(), &pb.GetDocumentsAdminRequest{
		Page: page,
		Size: size,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var documentsMap []map[string]any
	//if err := json.Unmarshal(resp.Documents, &documentsMap); err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "Unmarshal failure " + err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		//"documents":  documentsMap,
		"documents":  json.RawMessage(resp.Documents),
		"page_count": resp.PageCount,
	})

}

func HandleDeleteDocumentById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Please provide query params: id.",
		})
		return
	}
	log.Println("id:", id)
	resp, err := grpcClient.DocumentServiceClient.DeleteDocument(sysContext.Background(), &pb.DeleteDocumentRequest{
		Id: id,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"deleted": resp.Deleted,
		"msg":     resp.Msg,
	})
}

//func HandleUpdateDocument()  {
//
//}
