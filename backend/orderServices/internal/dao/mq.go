package dao

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"orderServices/internal/config/remote"
)

var (
	MqChannel *amqp.Channel
	conn      *amqp.Connection
)

const (
	routerKey  = "key1"
	queueName  = "order_queue"
	switchName = "order_sw"
)

// InitMq 初始化RabbitMQ连接和通道
func InitMq() {
	var err error
	mqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d/", remote.MyMqConfig.Username, remote.MyMqConfig.Password, remote.MyMqConfig.Endpoints, remote.MyMqConfig.Port)
	log.Println(mqDSN)
	conn, err = amqp.Dial(mqDSN)
	if err != nil {
		//return err
		log.Panicln("初始化Mq连接错误" + err.Error())
	}
	log.Println("初始化Mq连接成功")

	MqChannel, err = conn.Channel()
	if err != nil {
		//return err
		log.Panicln("初始化Mq通道" + err.Error())
	}

	log.Println("RabbitMQ initialized successfully")

	// 接下来是声明交换机
	// 声明交换机
	err = MqChannel.ExchangeDeclare(
		switchName, // 交换机名
		"direct",   // 交换机类型
		true,       // 持久化
		false,      // 自动删除
		false,      // 内部使用
		false,      // 是否等待
		nil,        // 额外参数
	)
	if err != nil {
		log.Printf("Failed to declare exchange: %v", err)
		log.Panicln("Failed to declare exchange" + err.Error())
	}

}

// CloseMq 关闭RabbitMQ连接
func CloseMq() {
	if err := MqChannel.Close(); err != nil {
		log.Println("Error closing channel:", err)
	}
	if err := conn.Close(); err != nil {
		log.Println("Error closing connection:", err)
	}
}
