package services

import (
	pb "NxcFull/backend/noticeServices/api/proto"
	"NxcFull/backend/noticeServices/internal/dao"
	"NxcFull/backend/noticeServices/internal/model"
	"context"
	"log"
	"net/http"
)

type NoticeServices struct {
	pb.UnimplementedNoticeServiceServer
}

func NewNoticeServices() *NoticeServices {
	return &NoticeServices{}
}

func (s *NoticeServices) GetAllNotices(context context.Context, request *pb.GetAllNoticesRequest) (*pb.GetAllNoticesResponse, error) {
	log.Println("GetAllNotices被调用了")
	var notices []model.PublicNotices
	result := dao.Db.Model(&model.PublicNotices{}).Where("deleted_at is null").Find(&notices)
	log.Println("获取通知列表结束")
	if result.Error != nil {
		log.Print(result.Error)
		return &pb.GetAllNoticesResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取所有通知失败" + result.Error.Error(),
		}, nil
	}
	// 转换为 []*pb.PublicNotice
	log.Println("转换前")
	protoNotices := convertToProtoNotices(notices)
	return &pb.GetAllNoticesResponse{
		Code:    http.StatusOK,
		Notices: protoNotices,
	}, nil
}

func (s *NoticeServices) AddNotice(context context.Context, request *pb.AddNoticeRequest) (*pb.AddNoticeResponse, error) {
	var notice model.PublicNotices
	notice.Title = request.Title
	notice.Content = request.Content
	notice.Tags = request.Tags
	notice.ImgUrl = request.ImgUrl
	if result := dao.Db.Model(&model.PublicNotices{}).Create(&notice); result.Error != nil {
		log.Println("新建通知错误", result.Error)
		return &pb.AddNoticeResponse{
			Code: http.StatusInternalServerError,
			Msg:  "新建通知错误" + result.Error.Error(),
		}, nil
	} else {
		return &pb.AddNoticeResponse{
			Code: http.StatusOK,
			Msg:  "新建通知成功",
		}, nil
	}
}

func (s *NoticeServices) DeleteNotice(context context.Context, request *pb.DeleteNoticeRequest) (*pb.DeleteNoticeResponse, error) {

	if err := dao.Db.Model(&model.PublicNotices{}).Where("id = ?", request.NoticeId).Delete(&model.PublicNotices{}).Error; err != nil {
		log.Println(err)
		return &pb.DeleteNoticeResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除通知失败" + err.Error(),
		}, nil
	} else {
		return &pb.DeleteNoticeResponse{
			Code: http.StatusOK,
			Msg:  "删除成功",
		}, nil
	}
}

//func HandleGetAllNotices(context *gin.Context) {
//	var notices []publicNotice.PublicNotices
//	result := dao.Db.Model(&publicNotice.PublicNotices{}).Where("deleted_at is null").Find(&notices)
//	if result.Error != nil {
//		log.Print(result.Error)
//	}
//	//log.Println("通知条数", len(Document))
//	//log.Println(Document)
//	context.JSON(http.StatusOK, gin.H{
//		"code":    http.StatusOK,
//		"notices": notices,
//		"msg":     "",
//	})
//}
//
////type noticeFormData struct {
////	Title   string `json:"title"`
////	Content string `json:"content"`
////	Tags    string `json:"tags"`
////	ImgUrl  string `json:"img_url"`
////}

//func HandleAddNotice(context *gin.Context) {
//	formData := &struct {
//		Title   string `json:"title"`
//		Content string `json:"content"`
//		Tags    string `json:"tags"`
//		ImgUrl  string `json:"img_url"`
//	}{}
//	if err := context.ShouldBind(&formData); err != nil {
//		log.Println(err)
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code": http.StatusInternalServerError,
//			"msg:": err.Error(),
//		})
//	}
//	log.Println(formData)
//	var notice publicNotice.PublicNotices
//	notice.Title = formData.Title
//	notice.Content = formData.Content
//	notice.Tags = formData.Tags
//	notice.ImgUrl = formData.ImgUrl
//	if result := dao.Db.Model(&publicNotice.PublicNotices{}).Create(&notice); result.Error != nil {
//		log.Println("新建通知错误", result.Error)
//	} else {
//		context.JSON(http.StatusOK, gin.H{
//			"code": http.StatusOK,
//		})
//	}
//}
//
//func HandleDeleteNotice(context *gin.Context) {
//	deleteNotice := &struct {
//		NoticeId int `json:"notice_id"`
//	}{}
//	var err error
//	if deleteNotice.NoticeId, err = strconv.Atoi(context.Query("notice_id")); err != nil {
//		log.Println(err)
//	}
//
//	log.Println("消息id: ", deleteNotice.NoticeId)
//
//	if err = dao.Db.Model(&publicNotice.PublicNotices{}).Where("id = ?", deleteNotice.NoticeId).Delete(&publicNotice.PublicNotices{}).Error; err != nil {
//		log.Println(err)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"error": err.Error(),
//			"msg":   "删除失败",
//		})
//	} else {
//		context.JSON(http.StatusOK, gin.H{
//			"code": http.StatusOK,
//			"msg":  "删除成功",
//		})
//	}
//}
