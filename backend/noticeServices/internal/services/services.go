package services

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	pb "noticeServices/api/proto"
	"noticeServices/internal/dao"
	"noticeServices/internal/model"
)

type NoticeServices struct {
	pb.UnimplementedNoticeServiceServer
}

func NewNoticeServices() *NoticeServices {
	return &NoticeServices{}
}

func (s *NoticeServices) GetAllNotices(ctx context.Context, request *pb.GetAllNoticesRequest) (*pb.GetAllNoticesResponse, error) {
	//log.Println("GetAllNotices被调用了")
	var notices []model.PublicNotices
	var result *gorm.DB

	// 如果 request.IsUser 为 true，说明是用户请求，返回特定字段且无需分页
	if request.IsUser {
		// 只返回 `show` 为 true 的通知，并按创建时间倒序排序
		result = dao.Db.
			Model(&model.PublicNotices{}).
			Where("`show` = ?", true). // 用反引号包裹 show 以避免 SQL 语法冲突
			Order("created_at DESC").
			Select("id, title, content, img_url, tags").
			Find(&notices)

		// 检查查询是否有错误
		if result.Error != nil {
			log.Print(result.Error)
			return &pb.GetAllNoticesResponse{
				Code: http.StatusInternalServerError,
				Msg:  "获取通知失败: " + result.Error.Error(),
			}, nil
		}

		// 直接将 notices 数据序列化为 bytes
		serializedNotices, err := json.Marshal(notices)
		if err != nil {
			return &pb.GetAllNoticesResponse{
				Code: http.StatusInternalServerError,
				Msg:  "序列化失败: " + err.Error(),
			}, nil
		}

		// 返回序列化后的 Notices 数据
		return &pb.GetAllNoticesResponse{
			Code:    http.StatusOK,
			Notices: serializedNotices, // Notices 为 bytes 类型
		}, nil
	}

	// 管理员请求，启用分页
	offset := (request.Page - 1) * request.Size
	result = dao.Db.Model(&model.PublicNotices{}).
		Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(request.Size)).
		Find(&notices)

	// 检查查询是否有错误
	if result.Error != nil {
		log.Print(result.Error)
		return &pb.GetAllNoticesResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取所有通知失败: " + result.Error.Error(),
		}, nil
	}

	// 计算总页数
	var totalRecords int64
	dao.Db.Model(&model.PublicNotices{}).Count(&totalRecords)
	totalPages := int64(math.Ceil(float64(totalRecords) / float64(request.Size)))

	// 直接将 notices 数据序列化为 bytes
	serializedNotices, err := json.Marshal(notices)
	if err != nil {
		return &pb.GetAllNoticesResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化失败: " + err.Error(),
		}, nil
	}

	// 返回序列化后的 Notices 数据和分页信息
	//log.Println(string(serializedNotices))
	return &pb.GetAllNoticesResponse{
		Code:      http.StatusOK,
		Notices:   serializedNotices, // Notices 为 bytes 类型
		PageCount: totalPages,        // 返回总页数
		Msg:       "获取成功",
	}, nil
}

// passed

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

func (s *NoticeServices) UpdateNotice(context context.Context, request *pb.UpdateNoticeRequest) (*pb.UpdateNoticeResponse, error) {
	notice := model.PublicNotices{
		Id:      request.Id,
		Title:   request.Title,
		Content: request.Content,
		Tags:    request.Tags,
		ImgUrl:  request.ImgUrl,
	}
	// 更新操作
	if result := dao.Db.Model(&model.PublicNotices{}).Where("id = ?", request.Id).Updates(&notice); result.Error != nil {
		log.Println("更新通知错误", result.Error)
		return &pb.UpdateNoticeResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新通知错误: " + result.Error.Error(),
		}, nil
	}

	return &pb.UpdateNoticeResponse{
		Code: http.StatusOK,
		Msg:  "更新通知成功",
	}, nil
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

func (s *NoticeServices) UpdateNoticeStatus(context context.Context, request *pb.UpdateNoticeStatusRequest) (*pb.UpdateNoticeStatusResponse, error) {
	log.Println("UpdateNoticeStatusRequest: ", request.Id, request.IsShow)
	if result := dao.Db.Model(&model.PublicNotices{}).Where("id = ?", request.Id).Update("`show`", request.IsShow); result.Error != nil {
		log.Println("更新通知失败" + result.Error.Error())
		return &pb.UpdateNoticeStatusResponse{
			Code: http.StatusInternalServerError,
			Msg:  "更新通知失败" + result.Error.Error(),
		}, nil
	}
	return &pb.UpdateNoticeStatusResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil

}
