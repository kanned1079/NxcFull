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
