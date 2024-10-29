package services

import (
	"context"
	"log"
	"net/http"
	pb "ticketHandleServices/api/proto"
	"ticketHandleServices/internal/dao"
	"ticketHandleServices/internal/model"
)

func (s *TicketHandleServices) CloseTicketByTicketId(context context.Context, request *pb.CloseTicketByTicketIdRequest) (*pb.CloseTicketByTicketIdResponse, error) {
	//var ticket model.Ticket
	if result := dao.Db.Model(&model.Ticket{}).Where("id = ?", request.TicketId).Update("status", http.StatusNoContent); result.Error != nil {
		log.Println(result.Error)
		return &pb.CloseTicketByTicketIdResponse{
			Code:   http.StatusInternalServerError,
			Msg:    "查詢失敗",
			Closed: false,
		}, nil
	}
	return &pb.CloseTicketByTicketIdResponse{
		Code:   http.StatusOK,
		Closed: true,
		Msg:    "关闭工单成功",
	}, nil
}
