package model

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// PublishOrderMessage 将订单信息发布到RabbitMQ
func PublishOrderMessage(postOrderData []byte) error {
	// 声明交换机
	err := MqChannel.ExchangeDeclare(
		"test_sw", // 交换机名
		"direct",  // 交换机类型
		true,      // 持久化
		false,     // 自动删除
		false,     // 内部使用
		false,     // 是否等待
		nil,       // 额外参数
	)
	if err != nil {
		log.Printf("Failed to declare exchange: %v", err)
		return err
	}

	// 发布消息
	err = MqChannel.Publish(
		"test_sw", // 交换机名
		"key",     // 路由键
		false,     // 是否强制推送
		false,     // 是否等待推送完成
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
