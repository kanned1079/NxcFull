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

// CheckIsUserHaveOpeningTickets 检查用户是否有还没有关闭的工单 管理员和用户都可请求
// 返回值中 Exist:是否有没关闭的工单 TicketCount:有多少没关闭的工单
func (s *TicketHandleServices) CheckIsUserHaveOpeningTickets(ctx context.Context, request *pb.CheckIsUserHaveOpeningTicketsRequest) (*pb.CheckIsUserHaveOpeningTicketsResponse, error) {
	// 初始化变量
	var isAdmin bool
	var ticketCount int64

	// 1. 检查用户是否为管理员
	if err := dao.Db.Model(&model.User{}).
		Where("id = ?", request.UserId).
		Select("is_admin").
		Scan(&isAdmin).Error; err != nil {
		return &pb.CheckIsUserHaveOpeningTicketsResponse{
			Code:        http.StatusInternalServerError,
			Msg:         "查询用户信息失败",
			Exist:       false,
			TicketCount: 0,
		}, err
	}

	// 2. 根据用户身份查询未关闭的工单数量
	if isAdmin {
		// 管理员：查询所有未关闭的工单
		if err := dao.Db.Model(&model.Ticket{}).
			Where("status != ?", 204). // 204 表示关闭的状态
			Count(&ticketCount).Error; err != nil {
			return &pb.CheckIsUserHaveOpeningTicketsResponse{
				Code:        http.StatusInternalServerError,
				Msg:         "查询管理员工单失败",
				Exist:       false,
				TicketCount: 0,
			}, err
		}
	} else {
		// 普通用户：查询该用户的未关闭工单
		if err := dao.Db.Model(&model.Ticket{}).
			Where("user_id = ? AND status != ?", request.UserId, 204).
			Count(&ticketCount).Error; err != nil {
			return &pb.CheckIsUserHaveOpeningTicketsResponse{
				Code:        http.StatusInternalServerError,
				Msg:         "查询用户工单失败",
				Exist:       false,
				TicketCount: 0,
			}, err
		}
	}

	// 3. 根据结果构造返回值
	return &pb.CheckIsUserHaveOpeningTicketsResponse{
		Code:        http.StatusOK,
		Msg:         "success",
		Exist:       ticketCount > 0, // 存在未关闭工单的标志
		TicketCount: ticketCount,
	}, nil
}

func (s *TicketHandleServices) GetUserTicketsByUserId(context context.Context, request *pb.GetUserTicketsByUserIdRequest) (*pb.GetUserTicketsByUserIdResponse, error) {
	var tickets []model.Ticket
	if result := dao.Db.Model(&model.Ticket{}).Where("user_id = ?", request.UserId).Find(&tickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找失败" + result.Error.Error(),
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

func (s *TicketHandleServices) GetAllTicket(context context.Context, request *pb.GetAllTicketRequest) (*pb.GetAllTicketResponse, error) {
	/*
		修改功能
		该查询分为两个部分 1.查询活跃的工单 2.查询已经完成的工单的数量
		因此需要修改分页参数 一共有四个 每一组两个
		request.PendingPage	活跃工单的页面
		request.PendingSize	活跃工单的每一个页面的条目
		request.FinishedPage 已经结束的工单的页面
		request.FinishedSize 已经结束工单的每一个页面的条目
		还需要根据这些分页参数来计算总的页面数量
	*/

	var err error
	var pendingTickets, finishedTickets []model.Ticket
	var pendingTicketsJson, finishedTicketsJson json.RawMessage

	// 计算活跃工单分页的起始位置
	pendingOffset := (request.PendingPage - 1) * request.PendingSize
	pendingLimit := request.PendingSize

	// 计算已结束工单分页的起始位置
	finishedOffset := (request.FinishedPage - 1) * request.FinishedSize
	finishedLimit := request.FinishedSize

	// 查询活跃的工单（status != 204 并分页）
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status != ?", 204).
		Order("created_at DESC").
		Offset(int(pendingOffset)).
		Limit(int(pendingLimit)).
		Find(&pendingTickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询活跃工单失败：" + result.Error.Error(),
		}, nil
	}

	// 查询已结束的工单（status = 204 并分页）
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status = ?", 204).
		Order("created_at DESC").
		Offset(int(finishedOffset)).
		Limit(int(finishedLimit)).
		Find(&finishedTickets); result.Error != nil {
		log.Println(result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询已结束工单失败：" + result.Error.Error(),
		}, nil
	}

	// 计算活跃工单总数
	var totalPendingTickets int64
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status != ?", 204).
		Count(&totalPendingTickets); result.Error != nil {
		log.Println("无法计算活跃工单总数：", result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "无法计算活跃工单总数：" + result.Error.Error(),
		}, nil
	}

	// 计算已结束工单总数
	var totalFinishedTickets int64
	if result := dao.Db.Model(&model.Ticket{}).
		Where("status = ?", 204).
		Count(&totalFinishedTickets); result.Error != nil {
		log.Println("无法计算已结束工单总数：", result.Error)
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "无法计算已结束工单总数：" + result.Error.Error(),
		}, nil
	}

	// 计算活跃工单的总页数
	pendingPageCount := int((totalPendingTickets + int64(request.PendingSize) - 1) / int64(request.PendingSize))

	// 计算已结束工单的总页数
	finishedPageCount := int((totalFinishedTickets + int64(request.FinishedSize) - 1) / int64(request.FinishedSize))

	// 序列化工单数据
	if pendingTicketsJson, err = json.Marshal(pendingTickets); err != nil {
		log.Println("活跃工单序列化失败")
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "活跃工单序列化失败：" + err.Error(),
		}, nil
	}
	if finishedTicketsJson, err = json.Marshal(finishedTickets); err != nil {
		log.Println("已结束工单序列化失败")
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "已结束工单序列化失败：" + err.Error(),
		}, nil
	}

	// 返回结果
	return &pb.GetAllTicketResponse{
		Code:              http.StatusOK,
		Msg:               "获取成功",
		PendingTickets:    pendingTicketsJson,
		Tickets:           finishedTicketsJson,
		PendingPageCount:  int64(pendingPageCount),
		FinishedPageCount: int64(finishedPageCount),
	}, nil
}
