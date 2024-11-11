package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type QueueInfo struct {
	Name            string `json:"name"`           // 队列名称
	Messages        int    `json:"messages"`       // 当前消息数量
	MessagesReady   int    `json:"messages_ready"` // 可用消息数量
	MessagesDetails struct {
		Rate float64 `json:"rate"` // 入队速率
	} `json:"messages_details"`
	MessagesStats struct {
		DeliverGet struct {
			Rate float64 `json:"rate"` // 出队速率
		} `json:"deliver_get_details"`
	} `json:"message_stats"`
}

func getQueueInfo(username, password, host string, port int, vhost, queueName string) (*QueueInfo, error) {
	url := fmt.Sprintf("https://%s:%d/api/queues/%s/%s", host, port, vhost, queueName)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var queueInfo QueueInfo
	if err := json.Unmarshal(body, &queueInfo); err != nil {
		return nil, err
	}

	return &queueInfo, nil
}

func GetSystemInfrastructureInfo(context *gin.Context) {
	username := "kanna"    // RabbitMQ 用户名
	password := "password" // RabbitMQ 密码
	host := "ikanned.com"  // RabbitMQ 主机地址
	port := 25673          // RabbitMQ 管理端口
	vhost := "/"           // 默认的虚拟主机，通常是 "/"

	queueNames := []string{"order_queue", "ticket_queue"}
	for _, queueName := range queueNames {
		queueInfo, err := getQueueInfo(username, password, host, port, vhost, queueName)
		if err != nil {
			fmt.Printf("Error fetching queue info for %s: %v\n", queueName, err)
			continue
		}

		fmt.Printf("Queue Name: %s\n", queueInfo.Name)
		fmt.Printf("  Current Messages: %d\n", queueInfo.Messages)
		fmt.Printf("  Messages Ready: %d\n", queueInfo.MessagesReady)
		fmt.Printf("  Message Rate (per second): %.2f\n", queueInfo.MessagesDetails.Rate)
		fmt.Printf("  Delivery Rate (per second): %.2f\n\n", queueInfo.MessagesStats.DeliverGet.Rate)

		context.JSON(http.StatusOK, queueInfo)
	}
}
