package puiblisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"orderServices/internal/dao"
)

const (
	routerKey  = "key1"
	queueName  = "order_queue"
	switchName = "order_sw"
)

// PublishOrderMessage 将订单信息发布到RabbitMQ
func PublishOrderMessage(postOrderData []byte) error {
	// 声明交换机
	var err error
	//err := dao.MqChannel.ExchangeDeclare(
	//	switchName, // 交换机名
	//	"direct",   // 交换机类型
	//	true,       // 持久化
	//	false,      // 自动删除
	//	false,      // 内部使用
	//	false,      // 是否等待
	//	nil,        // 额外参数
	//)
	//if err != nil {
	//	log.Printf("Failed to declare exchange: %v", err)
	//	return err
	//}

	log.Println(string(postOrderData))

	// 发布消息
	err = dao.MqChannel.Publish(
		switchName, // 交换机名
		routerKey,  // 路由键
		false,      // 是否强制推送
		false,      // 是否等待推送完成
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        postOrderData,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Order %s has been published to MQ")
	return nil
}
