package dao

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"orderHandleServices/internal/config/remote"
)

const (
	// 订单队列
	OrderRouterKey  = "key1"
	OrderQueueName  = "order_queue"
	OrderSwitchName = "order_sw"

	// 充值队列
	TopUpOrderRouterKey  = "key2"
	TopUpOrderQueueName  = "top_up_order_queue"
	TopUpOrderSwitchName = "top_up_order_sw"
)

var (
	MqConn          *amqp.Connection
	MqChannel       *amqp.Channel
	OrderQueue      amqp.Queue
	TopUpOrderQueue amqp.Queue
)

func InitOrderMq() {
	log.Println("init mq")
	mqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		remote.MyMqConfig.Username,
		remote.MyMqConfig.Password,
		remote.MyMqConfig.Endpoints,
		remote.MyMqConfig.Port,
	)

	var err error
	MqConn, err = amqp.Dial(mqDSN)
	if err != nil {
		log.Panicln("初始化Mq服务出错" + err.Error())
	}
	// 创建通道
	MqChannel, err = MqConn.Channel()
	if err != nil {
		log.Panicln("创建通道失败" + err.Error())
	}

	// 声明交换机 (direct 类型)
	err = MqChannel.ExchangeDeclare(
		OrderSwitchName, // 交换机名称
		"direct",        // 交换机类型
		true,            // 持久化
		false,           // 自动删除
		false,           // 排他性
		false,           // 是否阻塞
		nil,             // 额外参数
	)
	if err != nil {
		//return fmt.Errorf("声明交换机失败: %v", err)
		log.Panicln("声明交换机失败" + err.Error())
	}

	// 声明队列
	OrderQueue, err = MqChannel.QueueDeclare(
		OrderQueueName, // 队列名称
		true,           // 持久化
		false,          // 自动删除
		false,          // 排他性
		false,          // 是否阻塞
		nil,            // 额外参数
	)
	if err != nil {
		//return fmt.Errorf("声明队列失败: %v", err)
		log.Panicln("声明队列失败" + err.Error())
	}
	log.Println("声明队列成功")

	// 绑定队列到交换机，使用路由键
	//routingKey := "key"
	err = MqChannel.QueueBind(
		OrderQueue.Name, // 队列名称
		OrderRouterKey,  // 路由键
		OrderSwitchName, // 交换机名称
		false,
		nil,
	)
	if err != nil {
		//return fmt.Errorf("绑定队列失败: %v", err)
		log.Panicln("绑定队列失败" + err.Error())

	}
}

func InitTopUpOrderMq() {
	log.Println("init mq")
	mqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		remote.MyMqConfig.Username,
		remote.MyMqConfig.Password,
		remote.MyMqConfig.Endpoints,
		remote.MyMqConfig.Port,
	)

	var err error
	MqConn, err = amqp.Dial(mqDSN)
	if err != nil {
		log.Panicln("初始化Mq服务出错" + err.Error())
	}
	// 创建通道
	MqChannel, err = MqConn.Channel()
	if err != nil {
		log.Panicln("创建通道失败" + err.Error())
	}

	// 声明交换机 (direct 类型)
	err = MqChannel.ExchangeDeclare(
		TopUpOrderSwitchName, // 交换机名称
		"direct",             // 交换机类型
		true,                 // 持久化
		false,                // 自动删除
		false,                // 排他性
		false,                // 是否阻塞
		nil,                  // 额外参数
	)
	if err != nil {
		//return fmt.Errorf("声明交换机失败: %v", err)
		log.Panicln("声明交换机失败" + err.Error())
	}

	// 声明队列
	TopUpOrderQueue, err = MqChannel.QueueDeclare(
		TopUpOrderQueueName, // 队列名称
		true,                // 持久化
		false,               // 自动删除
		false,               // 排他性
		false,               // 是否阻塞
		nil,                 // 额外参数
	)
	if err != nil {
		//return fmt.Errorf("声明队列失败: %v", err)
		log.Panicln("声明队列失败" + err.Error())
	}
	log.Println("声明队列成功")

	// 绑定队列到交换机，使用路由键
	//routingKey := "key"
	err = MqChannel.QueueBind(
		TopUpOrderQueue.Name, // 队列名称
		TopUpOrderRouterKey,  // 路由键
		TopUpOrderSwitchName, // 交换机名称
		false,
		nil,
	)
	if err != nil {
		//return fmt.Errorf("绑定队列失败: %v", err)
		log.Panicln("绑定队列失败" + err.Error())

	}
}
