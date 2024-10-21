package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"time"
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
	// 初始化

	// 发送消息Hello World
	//for i := 0; i < 100000; i++ {
	//	sendMsg()
	//	time.Sleep(time.Nanosecond * 200)
	//
	//}
	for {
		sendMsg()
		//time.Sleep(time.Nanosecond * 10)
	}

	// 在推出时关闭连接和通道
	func() {
		err := conn.Close()
		failOnError(err, "Failed to close connection")
	}()
	func() {
		err := channel.Close()
		failOnError(err, "Failed to close channel")
	}()
}

func sendMsg() {

	// 首先需要先定义一个队列才可以发送消息

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	body := "Hello World!" + strconv.FormatInt(time.Now().UnixNano(), 10) // 要发送的数据
	err = channel.PublishWithContext(ctx,
		"",
		newQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}
