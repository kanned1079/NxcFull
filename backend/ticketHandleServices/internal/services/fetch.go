package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "ticketHandleServices/api/proto"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
)

func (s *TicketHandleServices) GetUserTicketsByUserId(context context.Context, request *pb.GetUserTicketsByUserIdRequest) (*pb.GetUserTicketsByUserIdResponse, error) {
	var tickets []model.Ticket
	if result := dao.Db.Model(&model.Ticket{}).Where("user_id = ?", request.UserId).Find(&tickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找失败" + result.Error.Error(),
		}, nil
	}
	//log.Println("tickets:", tickets)
	if len(tickets) == 0 {
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusNotFound,
			Msg:  "没有工单记录",
		}, nil
	}
	if ticketsJson, err := json.Marshal(tickets); err != nil {
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转化格式失败",
		}, nil
	} else {
		return &pb.GetUserTicketsByUserIdResponse{
			Code:    http.StatusOK,
			Msg:     "查询成功",
			Tickets: ticketsJson,
		}, nil
	}
}

func (s *TicketHandleServices) GetChatContent(context context.Context, request *pb.GetChatContentRequest) (*pb.GetChatContentResponse, error) {
	log.Printf("查询聊天记录 用户%v 工单%v", request.UserId, request.TicketId)
	var chatHistory []model.Chat
	if result := dao.Db.Model(&model.Chat{}).Where("ticket_id = ?", request.TicketId).Find(&chatHistory); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetChatContentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取失败",
		}, nil
	}
	if len(chatHistory) == 0 {
		log.Println("no chat history")
		return &pb.GetChatContentResponse{
			Code: http.StatusNotFound,
			Msg:  "还没有聊天内容",
		}, nil
	}
	//log.Println("history:", chatHistory)
	chatHistoryJson, err := json.Marshal(chatHistory)
	if err != nil {
		return &pb.GetChatContentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转换失败",
		}, nil
	}
	return &pb.GetChatContentResponse{
		Code:    http.StatusOK,
		Msg:     "获取失败",
		Content: chatHistoryJson,
	}, nil
}
