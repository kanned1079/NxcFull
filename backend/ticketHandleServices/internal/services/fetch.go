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
	//if len(tickets) == 0 {
	//	return &pb.GetUserTicketsByUserIdResponse{
	//		Code: http.StatusNotFound,
	//		Msg:  "没有工单记录",
	//	}, nil
	//}
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

func (s *TicketHandleServices) GetAllTicket(context context.Context, request *pb.GetAllTicketRequest) (*pb.GetAllTicketResponse, error) {
	// 初始化变量
	var err error
	var tickets []model.Ticket
	var activeTickets []model.Ticket
	var ticketsJson, activeTicketsJson json.RawMessage

	// 计算分页的起始位置
	offset := (request.Page - 1) * request.Size
	limit := request.Size

	// 查询状态为 204 的工单（所有工单）并分页
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status = ?", 204).
		Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&tickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败：" + result.Error.Error(),
		}, nil
	}

	// 查询所有活跃的工单（status != 204，不分页）
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status != ?", 204).
		Order("created_at DESC").
		Find(&activeTickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败：" + result.Error.Error(),
		}, nil
	}

	// 计算所有工单总数以进行页数计算
	var totalTickets int64
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status = ?", 204).
		Count(&totalTickets); result.Error != nil {
		log.Println("无法计算总页数：", result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "无法计算总页数：" + result.Error.Error(),
		}, nil
	}
	pageCount := int((totalTickets + int64(request.Size) - 1) / int64(request.Size))

	// 序列化 tickets 和 activeTickets
	if ticketsJson, err = json.Marshal(tickets); err != nil {
		log.Println("tickets 序列化失败")
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化失败：" + err.Error(),
		}, nil
	}
	if activeTicketsJson, err = json.Marshal(activeTickets); err != nil {
		log.Println("activeTickets 序列化失败")
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "序列化失败：" + err.Error(),
		}, nil
	}

	// 返回结果
	return &pb.GetAllTicketResponse{
		Code:           http.StatusOK,
		Msg:            "获取成功",
		Tickets:        ticketsJson,
		PendingTickets: activeTicketsJson,
		PageCount:      int64(pageCount), // 设置总页数
	}, nil
}
