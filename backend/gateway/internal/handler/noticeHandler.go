package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/notice/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// HandleGetAllNotices 处理通知
func HandleGetAllNotices(context *gin.Context) {
	queryData := &struct {
		Page   int64 `form:"page" json:"page"`
		Size   int64 `form:"size" json:"size"`
		IsUser bool  `form:"is_user" json:"is_user"`
	}{}
	if err := context.ShouldBindQuery(queryData); err != nil {
		log.Println(err)
	}
	log.Println(queryData)
	// 调用 gRPC 客户端获取通知
	resp, err := grpcClient.NoticeServiceClient.GetAllNotices(sysContext.Background(), &pb.GetAllNoticesRequest{
		Page:   queryData.Page,
		Size:   queryData.Size,
		IsUser: queryData.IsUser,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//// 反序列化 notices
	//var notices []map[string]interface{}
	////log.Println(string(resp.Notices))
	//err = json.Unmarshal(resp.Notices, &notices)
	//if err != nil {
	//	log.Println("反序列化通知列表失败:", err)
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "通知列表反序列化失败: " + err.Error(),
	//	})
	//	return
	//}

	// 返回响应给前端
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"notices":    json.RawMessage(resp.Notices), // 返回反序列化后的通知数据
		"msg":        resp.Msg,
		"page_count": resp.PageCount, // 总页数
	})
}

//type noticeFormData struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//	Tags    string `json:"tags"`
//	ImgUrl  string `json:"img_url"`
//}

func HandleAddNotice(context *gin.Context) {
	formData := &struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Tags    string `json:"tags"`
		ImgUrl  string `json:"img_url"`
	}{}
	if err := context.ShouldBind(&formData); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusInternalServerError,
			"msg:": err.Error(),
		})
	}
	log.Println(formData)
	resp, err := grpcClient.NoticeServiceClient.AddNotice(sysContext.Background(), &pb.AddNoticeRequest{
		Title:   formData.Title,
		Content: formData.Content,
		Tags:    formData.Tags,
		ImgUrl:  formData.ImgUrl,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
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

func HandleUpdateNotice(context *gin.Context) {
	formData := &struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Tags    string `json:"tags"`
		ImgUrl  string `json:"img_url"`
	}{}
	if err := context.ShouldBind(&formData); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusInternalServerError,
			"msg:": err.Error(),
		})
	}
	log.Println(formData)
	resp, err := grpcClient.NoticeServiceClient.UpdateNotice(sysContext.Background(), &pb.UpdateNoticeRequest{
		Id:      formData.Id,
		Title:   formData.Title,
		Content: formData.Content,
		Tags:    formData.Tags,
		ImgUrl:  formData.ImgUrl,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
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

func HandleUpdateNoticeStatus(context *gin.Context) {
	postData := &struct {
		Id     int64 `form:"id" json:"id"`
		IsShow bool  `form:"is_show" json:"is_show"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	log.Println(postData)
	resp, err := grpcClient.NoticeServiceClient.UpdateNoticeStatus(sysContext.Background(), &pb.UpdateNoticeStatusRequest{
		Id:     postData.Id,
		IsShow: postData.IsShow,
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
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

//func HandleUpdateNotice(context *gin.Context) {
//
//}

func HandleDeleteNotice(context *gin.Context) {
	deleteNotice := &struct {
		NoticeId int `json:"notice_id"`
	}{}
	var err error
	if deleteNotice.NoticeId, err = strconv.Atoi(context.Query("notice_id")); err != nil {
		log.Println(err)
	}

	log.Println("消息id: ", deleteNotice.NoticeId)

	resp, err := grpcClient.NoticeServiceClient.DeleteNotice(sysContext.Background(), &pb.DeleteNoticeRequest{
		NoticeId: int64(deleteNotice.NoticeId),
	})
	if err = failOnRpcError(err, context.Request.RequestURI); err != nil {
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
