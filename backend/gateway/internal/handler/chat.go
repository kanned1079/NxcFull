package handler

//
//import (
//	sysContext "context"
//	"encoding/json"
//	ticketPb "gateway/internal/grpc/api/ticketHandle/proto"
//	pb "gateway/internal/grpc/api/user/proto"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//var jwtSecret = []byte("Og6zf&J#OJTkw4blmpeQ_(hx~!1p%r%fCq%Stv&^fL%6@4kL0i#l$O7(4ZddI71s)_&+KuX")
//
//type response struct {
//	Code    int32  `json:"code"`
//	Type    string `json:"type"`
//	Msg     string `json:"msg"`
//	Sent    bool   `json:"sent"`
//	History []byte `json:"history"`
//}
//
//func HandleChatWs(context *gin.Context) {
//	// JWT 验证
//	tokenString := context.Query("token")
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//	if err != nil || !token.Valid {
//		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//		context.Abort()
//		return
//	}
//	userId, _ := strconv.ParseInt(context.Query("user_id"), 10, 64)
//	ticketId, _ := strconv.ParseInt(context.Query("ticket_id"), 10, 64)
//
//	// WebSocket 升级
//	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
//	if err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
//		return
//	}
//	defer conn.Close()
//
//	// 结束信道
//	done := make(chan struct{})
//	defer close(done)
//
//	// 定时消息发送的 goroutine
//	go func() {
//		ticker := time.NewTicker(1 * time.Second)
//		defer ticker.Stop()
//		for {
//			select {
//			case <-done:
//				return
//			case <-ticker.C:
//				// 发送心跳包
//				resp, err := grpcClient.TicketHandleClient.GetChatContent(sysContext.Background(), &ticketPb.GetChatContentRequest{
//					UserId:   userId,
//					TicketId: ticketId,
//				})
//				if err = failOnRpcError(err, resp); err != nil {
//					log.Println(err.Error())
//					sendErrorResponse(conn, 200, "rpc调用失败: "+err.Error(), http.StatusInternalServerError)
//				}
//				log.Println(string(resp.Content))
//				if err := conn.WriteMessage(websocket.TextMessage, resp.Content); err != nil {
//					log.Println("Failed to send message:", err)
//					return
//				}
//			}
//		}
//	}()
//
//	// 主循环：接收和处理客户端消息
//	for {
//		// 读取消息
//		messageType, message, err := conn.ReadMessage()
//		if err != nil {
//			log.Println("Failed to read message:", err)
//			break
//		}
//
//		// 解析客户端消息
//		var msg struct {
//			UserId   int64  `json:"user_id"`
//			TicketId int64  `json:"ticket_id"`
//			Content  string `json:"content"`
//			Role     string `json:"role"`
//		}
//		if err := json.Unmarshal(message, &msg); err != nil {
//			log.Println("消息结构体反序列化失败")
//			sendErrorResponse(conn, messageType, "Failed to unmarshal message", http.StatusInternalServerError)
//			continue
//		}
//		log.Println("路由层反序列化信息成功 ")
//
//		// RPC 调用
//		resp, err := grpcClient.UserServiceClient.SendChatMessage(sysContext.Background(), &pb.SendChatMessageRequest{
//			UserId:   msg.UserId,
//			TicketId: msg.TicketId,
//			Content:  msg.Content,
//			Role:     msg.Role,
//		})
//		if err = failOnRpcError(err, resp); err != nil {
//			sendErrorResponse(conn, messageType, "rpc调用失败: "+err.Error(), http.StatusInternalServerError)
//			continue
//		}
//
//		// 发送成功响应
//		if err := conn.WriteMessage(messageType, message); err != nil {
//			log.Println("Failed to write message:", err)
//			break
//		}
//	}
//}
//
//// 封装错误响应发送逻辑
//func sendErrorResponse(conn *websocket.Conn, messageType int, msg string, code int32) {
//	response := response{Code: code, Msg: msg, Sent: false}
//	jsonResponse, _ := json.Marshal(response)
//	if err := conn.WriteMessage(messageType, jsonResponse); err != nil {
//		log.Println("Failed to write error response:", err)
//	}
//}

import (
	sysContext "context"
	"encoding/json"
	ticketPb "gateway/internal/grpc/api/ticketHandle/proto"
	pb "gateway/internal/grpc/api/user/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var jwtSecret = []byte("Og6zf&J#OJTkw4blmpeQ_(hx~!1p%r%fCq%Stv&^fL%6@4kL0i#l$O7(4ZddI71s)_&+KuX")

type response struct {
	Code    int32           `json:"code"`
	Type    string          `json:"type"`
	Msg     string          `json:"msg"`
	Sent    bool            `json:"sent"`
	History json.RawMessage `json:"history"`
}

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
	userId, _ := strconv.ParseInt(context.Query("user_id"), 10, 64)
	ticketId, _ := strconv.ParseInt(context.Query("ticket_id"), 10, 64)

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
				resp, err := grpcClient.TicketHandleClient.GetChatContent(sysContext.Background(), &ticketPb.GetChatContentRequest{
					UserId:   userId,
					TicketId: ticketId,
				})
				if err = failOnRpcError(err, resp); err != nil {
					log.Println(err.Error())
					sendErrorResponse(conn, websocket.TextMessage, "rpc调用失败: "+err.Error(), http.StatusInternalServerError)
					continue
				}
				historyResponse := response{
					Code:    200,
					Type:    "history",
					Msg:     "History retrieved successfully",
					Sent:    true,
					History: resp.Content,
				}
				//log.Println(historyResponse)
				historyJson, _ := json.Marshal(historyResponse)
				if err := conn.WriteMessage(websocket.TextMessage, historyJson); err != nil {
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
			log.Println("消息结构体反序列化失败")
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
		checkResponse := response{
			Code: 200,
			Type: "check",
			Msg:  "Message sent successfully",
			Sent: true,
		}
		checkJson, _ := json.Marshal(checkResponse)
		if err := conn.WriteMessage(messageType, checkJson); err != nil {
			log.Println("Failed to write message:", err)
			break
		}
	}
}

// 封装错误响应发送逻辑
func sendErrorResponse(conn *websocket.Conn, messageType int, msg string, code int32) {
	errorResponse := response{Code: code, Msg: msg, Sent: false, Type: "check"}
	jsonResponse, _ := json.Marshal(errorResponse)
	if err := conn.WriteMessage(messageType, jsonResponse); err != nil {
		log.Println("Failed to write error response:", err)
	}
}
