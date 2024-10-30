package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "userServices/api/proto"
	"userServices/internal/dao"
	"userServices/internal/model"
	"userServices/internal/puiblisher"
)

func (s *UserService) CreateNewTicket(context context.Context, request *pb.CreateNewTicketRequest) (*pb.CreateNewTicketResponse, error) {
	// 启动事务
	tx := dao.Db.Begin()
	if tx.Error != nil {
		return &pb.CreateNewTicketResponse{
			Code:    http.StatusInternalServerError,
			Msg:     "启动事务失败: " + tx.Error.Error(),
			Created: false,
		}, nil
	}

	// 创建工单
	newTicket := &model.Ticket{
		UserId:  request.UserId,
		Subject: request.Subject,
		Urgency: int8(request.Urgency),
		Status:  http.StatusCreated, // 201表示创建成功
	}
	if result := tx.Model(&model.Ticket{}).Create(newTicket); result.Error != nil {
		tx.Rollback() // 回滚事务
		return &pb.CreateNewTicketResponse{
			Code:    http.StatusInternalServerError,
			Msg:     "插入数据失败: " + result.Error.Error(),
			Created: false,
		}, nil
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // 如果提交失败，回滚事务
		return &pb.CreateNewTicketResponse{
			Code:    http.StatusInternalServerError,
			Msg:     "提交事务失败: " + err.Error(),
			Created: false,
		}, nil
	}

	return &pb.CreateNewTicketResponse{
		Code:    http.StatusOK,
		Msg:     "新建工单成功",
		Created: true,
	}, nil
}

func (s *UserService) SendChatMessage(context context.Context, request *pb.SendChatMessageRequest) (*pb.SendChatMessageResponse, error) {
	var msg = &struct {
		UserId   int64  `json:"user_id"`
		TicketId int64  `json:"ticket_id"`
		Content  string `json:"content"`
		Role     string `json:"role"`
	}{
		UserId:   request.UserId,
		TicketId: request.TicketId,
		Content:  request.Content,
		Role:     request.Role,
	}
	if jsonMsg, err := json.Marshal(msg); err != nil {
		return &pb.SendChatMessageResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化失败",
			Sent: false,
		}, nil
	} else {
		log.Println("user准备推入消息队列")
		if err := puiblisher.PublishChatMessage(jsonMsg); err != nil {
			return &pb.SendChatMessageResponse{
				Code: http.StatusInternalServerError,
				Msg:  "推入消息队列失败",
				Sent: false,
			}, nil
		}
		log.Println("user推入消息队列成功")
	}
	return &pb.SendChatMessageResponse{
		Code: http.StatusOK,
		Msg:  "发送成功",
		Sent: true,
	}, nil
}
