package model

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	MqChannel *amqp.Channel
	conn      *amqp.Connection
)

// InitMq 初始化RabbitMQ连接和通道
func InitMq() {
	var err error
	//mqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d/", remote.MyMqConfig.Username, remote.MyMqConfig.Password, remote.MyMqConfig.Endpoints, remote.MyMqConfig.Port)
	//log.Println(mqDSN)
	conn, err = amqp.Dial("amqp://services:Passcode1!@api.ikanned.com:25672/")
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
