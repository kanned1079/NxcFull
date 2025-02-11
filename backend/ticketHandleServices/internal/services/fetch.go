package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
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
			Where("`status` != ?", 204). // 204 表示关闭的状态
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
			Where("`user_id` = ? AND `status` != ?", request.UserId, 204).
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

//func (s *TicketHandleServices) GetUserTicketsByUserId(context context.Context, request *pb.GetUserTicketsByUserIdRequest) (*pb.GetUserTicketsByUserIdResponse, error) {
//	// request.Page
//	// request.Size
//	var tickets []model.Ticket
//	if result := dao.Db.Model(&model.Ticket{}).Where("user_id = ?", request.UserId).Find(&tickets); result.Error != nil {
//		log.Println(result.Error)
//		return &pb.GetUserTicketsByUserIdResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查找失败" + result.Error.Error(),
//		}, nil
//	}
//	if ticketsJson, err := json.Marshal(tickets); err != nil {
//		return &pb.GetUserTicketsByUserIdResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "转化格式失败",
//		}, nil
//	} else {
//		return &pb.GetUserTicketsByUserIdResponse{
//			Code:      http.StatusOK,
//			Msg:       "查询成功",
//			Tickets:   ticketsJson,
//			PageCount: 10, // 根据page和size进行计算
//		}, nil
//	}
//}

func (s *TicketHandleServices) GetUserTicketsByUserId(context context.Context, request *pb.GetUserTicketsByUserIdRequest) (*pb.GetUserTicketsByUserIdResponse, error) {
	// 参数校验
	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Size <= 0 {
		request.Size = 10
	}

	var tickets []model.Ticket
	var total int64

	// 计算总记录数
	if result := dao.Db.Model(&model.Ticket{}).Where("user_id = ?", request.UserId).Count(&total); result.Error != nil {
		//log.Println(result.Error)
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查找失败" + result.Error.Error(),
		}, nil
	}

	// 分页查询
	if result := dao.Db.Model(&model.Ticket{}).Where("user_id = ?", request.UserId).
		Offset(int((request.Page - 1) * request.Size)).
		Limit(int(request.Size)).
		Order("`created_at` DESC").
		Find(&tickets); result.Error != nil {
		//log.Println(result.Error)
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败" + result.Error.Error(),
		}, nil
	}

	// 计算总页数
	pageCount := math.Ceil(float64(total) / float64(request.Size))

	if ticketsJson, err := json.Marshal(tickets); err != nil {
		return &pb.GetUserTicketsByUserIdResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转化格式失败",
		}, nil
	} else {
		return &pb.GetUserTicketsByUserIdResponse{
			Code:      http.StatusOK,
			Msg:       "查询成功",
			Tickets:   ticketsJson,
			PageCount: int64(pageCount),
		}, nil
	}
}

//func (s *TicketHandleServices) GetChatContent(context context.Context, request *pb.GetChatContentRequest) (*pb.GetChatContentResponse, error) {
//	log.Printf("查询聊天记录 用户%v 工单%v", request.UserId, request.TicketId)
//	var chatHistory []model.Chat
//	if result := dao.Db.Model(&model.Chat{}).Where("ticket_id = ?", request.TicketId).Find(&chatHistory); result.Error != nil {
//		log.Println(result.Error)
//		return &pb.GetChatContentResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "获取失败",
//		}, nil
//	}
//	if len(chatHistory) == 0 {
//		log.Println("no chat history")
//		return &pb.GetChatContentResponse{
//			Code: http.StatusNotFound,
//			Msg:  "还没有聊天内容",
//		}, nil
//	}
//	//log.Println("history:", chatHistory)
//	chatHistoryJson, err := json.Marshal(chatHistory)
//	if err != nil {
//		return &pb.GetChatContentResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "转换失败",
//		}, nil
//	}
//	return &pb.GetChatContentResponse{
//		Code:    http.StatusOK,
//		Msg:     "获取失败",
//		Content: chatHistoryJson,
//	}, nil
//}

//	func (s *TicketHandleServices) GetChatContent(ctx context.Context, request *pb.GetChatContentRequest) (*pb.GetChatContentResponse, error) {
//		log.Printf("查询聊天记录 用户%v 工单%v", request.UserId, request.TicketId)
//
//		// 创建带有超时的 context（控制 Redis 请求的超时）
//		redisCtx := ctx
//		//defer cancel()
//
//		// Redis 列表的键名
//		chatCacheKey := fmt.Sprintf("ticket:%d:chats", request.TicketId)
//
//		// 从 Redis 获取最新的聊天记录（从列表尾部获取最多 100 条记录）
//		cachedChatHistory, err := dao.Rdb.LRange(redisCtx, chatCacheKey, 0, 99).Result()
//		if err != nil {
//			log.Println("从 Redis 获取聊天记录失败:", err)
//			return &pb.GetChatContentResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "获取聊天记录失败",
//			}, nil
//		}
//
//		if len(cachedChatHistory) == 0 {
//			log.Println("no chat history")
//			return &pb.GetChatContentResponse{
//				Code: http.StatusNotFound,
//				Msg:  "还没有聊天内容",
//			}, nil
//		}
//
//		// 反序列化聊天记录并反转顺序
//		var chatHistory []model.Chat
//		for i := len(cachedChatHistory) - 1; i >= 0; i-- { // 从最后一条消息开始反转顺序
//			var chatMsg model.Chat
//			if err := json.Unmarshal([]byte(cachedChatHistory[i]), &chatMsg); err != nil {
//				log.Println("反序列化聊天记录失败:", err)
//				return &pb.GetChatContentResponse{
//					Code: http.StatusInternalServerError,
//					Msg:  "转换聊天记录失败",
//				}, nil
//			}
//			chatHistory = append(chatHistory, chatMsg)
//		}
//
//		// 将反序列化后的聊天记录转换为 JSON 格式
//		chatHistoryJson, err := json.Marshal(chatHistory)
//		if err != nil {
//			log.Println("转换聊天记录为 JSON 失败:", err)
//			return &pb.GetChatContentResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "转换失败",
//			}, nil
//		}
//
//		// 返回聊天记录
//		return &pb.GetChatContentResponse{
//			Code:    http.StatusOK,
//			Msg:     "获取成功",
//			Content: chatHistoryJson,
//		}, nil
//	}

//func (s *TicketHandleServices) GetChatContent(ctx context.Context, request *pb.GetChatContentRequest) (*pb.GetChatContentResponse, error) {
//	// log.Printf("查询聊天记录 用户%v 工单%v", request.UserId, request.TicketId)
//
//	// Redis 缓存的 key
//	chatCacheKey := fmt.Sprintf("ticket:%d:chats", request.TicketId)
//
//	// 1. 先尝试从 Redis 获取聊天记录
//	cachedChatHistory, err := dao.Rdb.LRange(ctx, chatCacheKey, 0, -1).Result()
//	if err != nil {
//		log.Println("从 Redis 获取聊天记录失败:", err)
//	}
//
//	var chatHistory []model.Chat
//
//	if len(cachedChatHistory) > 0 {
//		// 解析 Redis 存储的 JSON 数据（按顺序添加）
//		for _, chatStr := range cachedChatHistory {
//			var chatMsg model.Chat
//			if err := json.Unmarshal([]byte(chatStr), &chatMsg); err != nil {
//				log.Println("反序列化聊天记录失败:", err)
//				return &pb.GetChatContentResponse{
//					Code: http.StatusInternalServerError,
//					Msg:  "转换聊天记录失败",
//				}, nil
//			}
//			chatHistory = append(chatHistory, chatMsg)
//		}
//	} else {
//		// 2. 如果 Redis 没有数据，从 MySQL 查询全部聊天记录
//		log.Println("Redis 无缓存数据，从数据库查询")
//		if result := dao.Db.Model(&model.Chat{}).Where("`ticket_id` = ?", request.TicketId).Order("sent_at ASC").Find(&chatHistory); result.Error != nil {
//			log.Println("数据库查询聊天记录失败:", result.Error)
//			return &pb.GetChatContentResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "获取聊天记录失败",
//			}, nil
//		}
//
//		if len(chatHistory) == 0 {
//			log.Println("数据库无聊天记录")
//			return &pb.GetChatContentResponse{
//				Code: http.StatusNotFound,
//				Msg:  "还没有聊天内容",
//			}, nil
//		}
//
//		// 3. 将查询到的 MySQL 数据缓存到 Redis
//		go func() {
//			cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 5*time.Second)
//			defer cacheCancel()
//			pipe := dao.Rdb.Pipeline()
//			for _, msg := range chatHistory {
//				chatJson, _ := json.Marshal(msg)
//				pipe.RPush(cacheCtx, chatCacheKey, chatJson) // 这里用 RPush，保证顺序和 MySQL 一致
//			}
//			// 设置 Redis 过期时间，防止长期占用
//			pipe.Expire(cacheCtx, chatCacheKey, 24*time.Hour)
//			if _, err := pipe.Exec(cacheCtx); err != nil {
//				log.Println("cache in redis failure:", err)
//			}
//		}()
//	}
//
//	// 4. 返回数据
//	chatHistoryJson, err := json.Marshal(chatHistory)
//	if err != nil {
//		log.Println("failure on Marshal", err)
//		return &pb.GetChatContentResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "failure",
//		}, nil
//	}
//
//	return &pb.GetChatContentResponse{
//		Code:    http.StatusOK,
//		Msg:     "success",
//		Content: chatHistoryJson,
//	}, nil
//}

func (s *TicketHandleServices) GetChatContent(ctx context.Context, request *pb.GetChatContentRequest) (*pb.GetChatContentResponse, error) {
	//log.Printf("查询聊天记录 用户%v 工单%v", request.UserId, request.TicketId)

	// Redis 缓存的 key
	chatCacheKey := fmt.Sprintf("ticket:%d:chats", request.TicketId)

	// 1. 先尝试从 Redis 获取聊天记录
	cachedChatHistory, err := dao.Rdb.LRange(ctx, chatCacheKey, 0, -1).Result()
	if err != nil {
		log.Println("从 Redis 获取聊天记录失败:", err)
	}

	var chatHistory []model.Chat

	if len(cachedChatHistory) > 0 {
		// 解析 Redis 存储的 JSON 数据（按顺序添加）
		for _, chatStr := range cachedChatHistory {
			var chatMsg model.Chat
			if err := json.Unmarshal([]byte(chatStr), &chatMsg); err != nil {
				//log.Println("反序列化聊天记录失败:", err)
				return &pb.GetChatContentResponse{
					Code: http.StatusInternalServerError,
					Msg:  "转换聊天记录失败",
				}, nil
			}
			chatHistory = append(chatHistory, chatMsg)
		}
	} else {
		// 2. 如果 Redis 没有数据，从 MySQL 查询全部聊天记录
		//log.Println("Redis 无缓存数据，从数据库查询")
		if result := dao.Db.Model(&model.Chat{}).Where("ticket_id = ?", request.TicketId).Order("sent_at ASC").Find(&chatHistory); result.Error != nil {
			//log.Println("数据库查询聊天记录失败:", result.Error)
			return &pb.GetChatContentResponse{
				Code: http.StatusInternalServerError,
				Msg:  "获取聊天记录失败",
			}, nil
		}

		if len(chatHistory) == 0 {
			//log.Println("数据库无聊天记录")
			return &pb.GetChatContentResponse{
				Code: http.StatusNotFound,
				Msg:  "还没有聊天内容",
			}, nil
		}

		// 3. 将查询到的 MySQL 数据缓存到 Redis
		err := s.cacheChatHistoryToRedis(ctx, chatCacheKey, chatHistory)
		if err != nil {
			log.Println("缓存聊天记录到 Redis 失败:", err)
		}
	}

	// 4. 返回数据
	chatHistoryJson, err := json.Marshal(chatHistory)
	if err != nil {
		//log.Println("转换聊天记录为 JSON 失败:", err)
		return &pb.GetChatContentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "转换失败",
		}, nil
	}

	return &pb.GetChatContentResponse{
		Code:    http.StatusOK,
		Msg:     "获取成功",
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
		//log.Println(result.Error)
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
		//log.Println(result.Error)
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
		//log.Println("无法计算活跃工单总数：", result.Error)
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
		//log.Println("无法计算已结束工单总数：", result.Error)
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
		//log.Println("活跃工单序列化失败")
		return &pb.GetAllTicketResponse{
			Code: http.StatusInternalServerError,
			Msg:  "活跃工单序列化失败：" + err.Error(),
		}, nil
	}
	if finishedTicketsJson, err = json.Marshal(finishedTickets); err != nil {
		//log.Println("已结束工单序列化失败")
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
