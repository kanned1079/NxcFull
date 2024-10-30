package puiblisher

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"userServices/internal/dao"
)

const (
	routerKey  = "key2"
	queueName  = "ticket_queue"
	switchName = "ticket_sw"
)

// PublishChatMessage 将工单聊天内容发布到RabbitMQ
func PublishChatMessage(postChatData []byte) error {
	var err error

	log.Println(string(postChatData))

	// 发布消息
	err = dao.MqChannel.Publish(
		switchName, // 交换机名
		routerKey,  // 路由键
		false,      // 是否强制推送
		false,      // 是否等待推送完成
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        postChatData,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Chat Msg %s has been published to MQ", string(postChatData))
	return nil
}
