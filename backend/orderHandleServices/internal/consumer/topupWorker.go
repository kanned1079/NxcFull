package consumer

import (
	"fmt"
	"log"
	"orderHandleServices/internal/dao"
)

// StartConsumeTopUpOrders 开始消费充值订单消息
func StartConsumeTopUpOrders() {
	log.Println("start consume top-up orders")
	msgs, err := dao.MqChannel.Consume(
		dao.TopUpOrderQueue.Name, // 队列名称
		"",                       // 消费者名称
		true,                     // 自动应答
		false,                    // 排他性
		false,                    // 队列阻塞
		false,                    // 队列立即消费
		nil,                      // 额外参数
	)
	if err != nil {
		panic(fmt.Sprintf("消费充值队列失败: %v", err))
	}

	go func() {
		for d := range msgs {
			// 处理订单
			log.Println("有新充值订单")
			err = ProcessTopUpOrder(d.Body)
			if err != nil {
				fmt.Printf("充值订单处理失败: %v\n", err)
			} else {
				fmt.Printf("充值订单处理成功\n")
			}
		}
	}()
}
