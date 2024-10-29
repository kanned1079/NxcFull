package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/user/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var jwtSecret = []byte("Og6zf&J#OJTkw4blmpeQ_(hx~!1p%r%fCq%Stv&^fL%6@4kL0i#l$O7(4ZddI71s)_&+KuX")

type sendMsgResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Sent bool   `json:"sent"`
}

//func HandleChatWs(context *gin.Context) {
//	tokenString := context.Query("token") // 从 URL 或其他方式获取 token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if err != nil || !token.Valid {
//		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//		context.Abort()
//		return
//	}
//
//	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
//	if err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Failed to upgrade to WebSocket",
//		})
//		return
//	}
//	defer func() {
//		if err := conn.Close(); err != nil {
//			log.Println("Failed to close WebSocket connection:", err)
//		}
//	}()
//
//	// 启动一个 goroutine 定时向客户端推送消息
//	go func() {
//		ticker := time.NewTicker(1 * time.Second) // 每1秒发送一次
//		defer ticker.Stop()
//		for {
//			select {
//			case <-ticker.C:
//				//message := []byte("Server heartbeat") // 示例消息
//				jsonData, _ := json.Marshal(&struct {
//					Id   int64  `json:"id"`
//					Name string `json:"name"`
//				}{
//					Id:   234,
//					Name: "kanna",
//				})
//				if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
//					log.Println("Failed to send message:", err)
//					return // 连接关闭或出错时退出
//				}
//			}
//		}
//	}()
//
//	// 接收并处理客户端消息
//	for {
//		messageType, message, err := conn.ReadMessage()
//		if err != nil {
//			log.Println("Failed to read message:", err)
//			break
//		}
//		var selfSentResponse sendMsgResponse
//		var msg = &struct {
//			UserId   int64  `json:"user_id"`
//			TicketId int64  `json:"ticket_id"`
//			Content  string `json:"content"`
//			Role     string `json:"role"`
//		}{}
//		if err := json.Unmarshal(message, &msg); err != nil {
//			log.Println("Failed to unmarshal message:", err)
//			selfSentResponse.Code = http.StatusInternalServerError
//			selfSentResponse.Msg = "Failed to unmarshal message"
//			selfSentResponse.Sent = false
//			// 回传给客户端
//			selfSentResponseJson, _ := json.Marshal(selfSentResponse)
//			if err := conn.WriteMessage(messageType, selfSentResponseJson); err != nil {
//				log.Println("Failed to write message:", err)
//				break
//			}
//		} else {
//			resp, err := grpcClient.UserServiceClient.SendChatMessage(sysContext.Background(), &pb.SendChatMessageRequest{
//				UserId:   msg.UserId,
//				TicketId: msg.TicketId,
//				Content:  msg.Content,
//				Role:     msg.Role,
//			})
//			if err = failOnRpcError(err, resp); err != nil {
//				log.Println("rpc err:", err)
//				selfSentResponse.Code = http.StatusInternalServerError
//				selfSentResponse.Msg = "rpc调用失败" + err.Error()
//				selfSentResponse.Sent = false
//				// 回传给客户端
//				selfSentResponseJson, _ := json.Marshal(selfSentResponse)
//				if err := conn.WriteMessage(messageType, selfSentResponseJson); err != nil {
//					log.Println("Failed to write message:", err)
//					break
//				}
//			}
//			if err := conn.WriteMessage(messageType, message); err != nil {
//				selfSentResponse.Code = resp.Code
//				selfSentResponse.Msg = resp.Msg
//				selfSentResponse.Sent = resp.Sent
//				selfSentResponseJson, _ := json.Marshal(selfSentResponse)
//				if err := conn.WriteMessage(messageType, selfSentResponseJson); err != nil {
//					log.Println("Failed to write message:", err)
//					break
//				}
//			}
//
//		}
//		//log.Println("Received message:", string(message))
//
//		//// 回传给客户端
//		//if err := conn.WriteMessage(messageType, message); err != nil {
//		//	log.Println("Failed to write message:", err)
//		//	break
//		//}
//	}
//}

func HandleChatWs(context *gin.Context) {
	// JWT 验证
	tokenString := context.Query("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}

	// WebSocket 升级
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	// 结束信道
	done := make(chan struct{})
	defer close(done)

	// 定时消息发送的 goroutine
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// 发送心跳包
				jsonData, _ := json.Marshal(struct {
					Id   int64  `json:"id"`
					Name string `json:"name"`
				}{
					Id:   234,
					Name: "kanna",
				})
				if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
					log.Println("Failed to send message:", err)
					return
				}
			}
		}
	}()

	// 主循环：接收和处理客户端消息
	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// 解析客户端消息
		var msg struct {
			UserId   int64  `json:"user_id"`
			TicketId int64  `json:"ticket_id"`
			Content  string `json:"content"`
			Role     string `json:"role"`
		}
		if err := json.Unmarshal(message, &msg); err != nil {
			sendErrorResponse(conn, messageType, "Failed to unmarshal message", http.StatusInternalServerError)
			continue
		}
		log.Println("路由层反序列化信息成功 ")

		// RPC 调用
		resp, err := grpcClient.UserServiceClient.SendChatMessage(sysContext.Background(), &pb.SendChatMessageRequest{
			UserId:   msg.UserId,
			TicketId: msg.TicketId,
			Content:  msg.Content,
			Role:     msg.Role,
		})
		if err = failOnRpcError(err, resp); err != nil {
			sendErrorResponse(conn, messageType, "rpc调用失败: "+err.Error(), http.StatusInternalServerError)
			continue
		}

		// 发送成功响应
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Failed to write message:", err)
			break
		}
	}
}

// 封装错误响应发送逻辑
func sendErrorResponse(conn *websocket.Conn, messageType int, msg string, code int32) {
	response := sendMsgResponse{Code: code, Msg: msg, Sent: false}
	jsonResponse, _ := json.Marshal(response)
	if err := conn.WriteMessage(messageType, jsonResponse); err != nil {
		log.Println("Failed to write error response:", err)
	}
}
