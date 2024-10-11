package handler

import (
	pb "NxcFull/backend/gateway/internal/grpc/api/notice/proto"
	sysContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllNotices(context *gin.Context) {
	//var notices []publicNotice.PublicNotices
	//result := dao.Db.Model(&publicNotice.PublicNotices{}).Where("deleted_at is null").Find(&notices)
	//if result.Error != nil {
	//	log.Print(result.Error)
	//}
	resp, err := grpcClient.NoticeServiceClient.GetAllNotices(sysContext.Background(), &pb.GetAllNoticesRequest{})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	//log.Println("通知条数", len(Document))
	//log.Println(Document)
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"notices": resp.Notices,
		"msg":     resp.Msg,
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
	//var notice publicNotice.PublicNotices
	//notice.Title = formData.Title
	//notice.Content = formData.Content
	//notice.Tags = formData.Tags
	//notice.ImgUrl = formData.ImgUrl
	//if result := dao.Db.Model(&publicNotice.PublicNotices{}).Create(&notice); result.Error != nil {
	//	log.Println("新建通知错误", result.Error)
	//} else {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusOK,
	//	})
	//}
	resp, err := grpcClient.NoticeServiceClient.AddNotice(sysContext.Background(), &pb.AddNoticeRequest{
		Title:   formData.Title,
		Content: formData.Content,
		Tags:    formData.Tags,
		ImgUrl:  formData.ImgUrl,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

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
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})

	//if err = dao.Db.Model(&publicNotice.PublicNotices{}).Where("id = ?", deleteNotice.NoticeId).Delete(&publicNotice.PublicNotices{}).Error; err != nil {
	//	log.Println(err)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"error": err.Error(),
	//		"msg":   "删除失败",
	//	})
	//} else {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusOK,
	//		"msg":  "删除成功",
	//	})
	//}
}
