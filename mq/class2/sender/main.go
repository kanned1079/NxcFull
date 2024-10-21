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
	// 创建一个新的交换机
	err = channel.ExchangeDeclare(
		"logs2",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")
	//// 创建一个新的队列
	//newQueue, err = channel.QueueDeclare(
	//	"hello_queue",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//failOnError(err, "Failed to declare a queue")
	// 注意要声明队列并绑定 否则发送的消息将会被丢弃
	newQueue, err := channel.QueueDeclare(
		"test1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.QueueBind(
		newQueue.Name,
		"",
		"logs2",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
}

func main() {
	// 初始化
	for i := 0; i < 100000000; i++ {
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

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	body := "Hello World!" + strconv.FormatInt(time.Now().UnixNano(), 10) // 要发送的数据
	err = channel.PublishWithContext(ctx,
		"logs2", // 这里是定义使用的交换机 比如说这里就叫logs
		newQueue.Name,
		// ,     // routing key
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
