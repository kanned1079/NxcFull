package puiblisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"orderServices/internal/dao"
)

const (
	topUpRouterKey  = "key2"
	topUpQueueName  = "top_up_order_queue"
	topUpSwitchName = "top_up_order_sw"
)

// PublishTopUpOrderMessage 将充值订单信息发布到RabbitMQ
func PublishTopUpOrderMessage(postOrderData []byte) error {
	// 声明交换机
	var err error

	log.Println(string(postOrderData))

	// 发布消息
	err = dao.MqChannel.Publish(
		topUpSwitchName, // 交换机名
		topUpRouterKey,  // 路由键
		false,           // 是否强制推送
		false,           // 是否等待推送完成
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        postOrderData,
		},
	)
	if err != nil {
		log.Printf("Failed to publish top-up message: %v", err)
		return err
	}

	log.Printf("TopUp Order %s has been published to MQ")
	return nil
}
