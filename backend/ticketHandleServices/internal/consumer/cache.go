package consumer

//func SaveMSgInRedis(ticketId int64, msg string) error {
//	// 将消息转换为 JSON 字符串
//	messageData, err := json.Marshal(chatMsg)
//	if err != nil {
//		tx.Rollback()
//		return &pb.CreateNewTicketResponse{
//			Code:    http.StatusInternalServerError,
//			Msg:     "消息序列化失败: " + err.Error(),
//			Created: false,
//		}, nil
//	}
//
//	// Redis List 键，使用工单 ID 作为 List 名称
//	cacheKey := "ticket:messages:" + strconv.FormatInt(newTicket.Id, 10)
//
//	// 初始化 Redis List 并插入第一条消息
//	redisContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	err = dao.Rdb.RPush(redisContext, cacheKey, messageData).Err()
//	if err != nil {
//		tx.Rollback()
//		return &pb.CreateNewTicketResponse{
//			Code:    http.StatusInternalServerError,
//			Msg:     "Redis 插入数据失败: " + err.Error(),
//			Created: false,
//		}, nil
//	}
//}
