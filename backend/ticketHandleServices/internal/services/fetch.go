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
	if result := dao.Db.Model(&model.Ticket{}).Find(&tickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找失败" + result.Error.Error(),
		}, nil
	}
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
	return &pb.GetChatContentResponse{}, nil
}
