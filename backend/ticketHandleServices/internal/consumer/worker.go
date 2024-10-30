package consumer

import (
	"fmt"
	"log"
	"ticketHandleServices/internal/dao"
)

// StartConsumeOrders 开始消费订单消息
func StartConsumeOrders() {
	log.Println("start consume tickets")
	msgs, err := dao.MqChannel.Consume(
		dao.TicketQueue.Name, // 队列名称
		"",                   // 消费者名称
		true,                 // 自动应答
		false,                // 排他性
		false,                // 队列阻塞
		false,                // 队列立即消费
		nil,                  // 额外参数
	)
	if err != nil {
		panic(fmt.Sprintf("消费队列失败: %v", err))
	}

	go func() {
		for d := range msgs {
			// 处理订单
			log.Println("ticket 有新工单消息")
			err = ProcessTicket(d.Body)
			if err != nil {
				fmt.Printf("ticket 工单信息处理失败: %v\n", err)
			} else {
				fmt.Printf("ticket 工单信息处理成功\n")
			}
		}
	}()
}
