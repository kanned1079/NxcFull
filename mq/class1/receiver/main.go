package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	conn     *amqp.Connection
	channel  *amqp.Channel
	newQueue amqp.Queue
	err      error
)

func init() {
	// 在这里初始化mq连接
	conn, err = amqp.Dial("amqp://services:Passcode1!@api.ikanned.com:25672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("RabbitMQ Connected Successfully")
	// 还需要新建一个通道 其中包含有实现大多数功能的api
	channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")
	// 创建一个新的队列
	newQueue, err = channel.QueueDeclare(
		"hello_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
}

func main() {
	startReceiver()
}

func startReceiver() {
	msg, err := channel.Consume(
		newQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	//var forever chan struct{}
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s, ID: %d\n", d.Body, d.MessageId)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}
