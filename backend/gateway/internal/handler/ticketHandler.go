package handler

import (
	sysContext "context"
	"encoding/json"
	ticketPb "gateway/internal/grpc/api/ticketHandle/proto"
	pb "gateway/internal/grpc/api/user/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetAllMyTickets(context *gin.Context) {
	err, page, size := GetPage2SizeFromQuery(context)
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数不合法",
		})
		return
	}
	resp, err := grpcClient.TicketHandleClient.GetUserTicketsByUserId(sysContext.Background(), &ticketPb.GetUserTicketsByUserIdRequest{
		UserId: userId,
		Page:   page,
		Size:   size,
	})
	if err := failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var ticketsMap []map[string]any
	if err = json.Unmarshal(resp.Tickets, &ticketsMap); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"tickets":    ticketsMap,
		"page_count": resp.PageCount,
		"msg":        resp.Msg,
	})
}

func HandleCreateNewTicket(context *gin.Context) {
	var postData = &struct {
		UserId  int64  `json:"user_id"`
		Subject string `json:"subject"`
		Urgency int    `json:"urgency"`
		Body    string `json:"body"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println(postData)
	resp, err := grpcClient.UserServiceClient.CreateNewTicket(sysContext.Background(), &pb.CreateNewTicketRequest{
		UserId:  postData.UserId,
		Subject: postData.Subject,
		Urgency: int32(postData.Urgency),
		Content: postData.Body,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"created": resp.Created,
	})
}

func HandleSendChatContent(context *gin.Context) {
	var postData = &struct {
		UserId   int64  `json:"user_id"`
		TicketId int64  `json:"ticket_id"`
		Role     string `json:"role"`
		Content  string `json:"content"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println(postData)
	resp, err := grpcClient.UserServiceClient.SendChatMessage(sysContext.Background(), &pb.SendChatMessageRequest{
		UserId:   postData.UserId,
		TicketId: postData.TicketId,
		Role:     postData.Role,
		Content:  postData.Content,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		"sent": resp.Sent,
	})
}

func HandleCloseTicket(context *gin.Context) {
	var userId, ticketId int64
	var err error
	userId, err = strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	ticketId, err = strconv.ParseInt(context.Query("ticket_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	resp, err := grpcClient.TicketHandleClient.CloseTicketByTicketId(sysContext.Background(), &ticketPb.CloseTicketByTicketIdRequest{
		UserId:   userId,
		TicketId: ticketId,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":   resp.Code,
		"msg":    resp.Msg,
		"closed": resp.Closed,
	})
}

//func HandleGetChatContent(context *gin.Context) {
//
//}

func HandleGetAllTickets(context *gin.Context) {
	pendingPage, err := strconv.ParseInt(context.Query("pending_page"), 10, 64)
	pendingSize, err := strconv.ParseInt(context.Query("pending_size"), 10, 64)
	finishedPage, err := strconv.ParseInt(context.Query("finished_page"), 10, 64)
	finishedSize, err := strconv.ParseInt(context.Query("finished_size"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.TicketHandleClient.GetAllTicket(sysContext.Background(), &ticketPb.GetAllTicketRequest{
		PendingPage:  pendingPage,
		PendingSize:  pendingSize,
		FinishedPage: finishedPage,
		FinishedSize: finishedSize,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var ticketsMap []map[string]any
	var activeTicketMap []map[string]any
	if err = json.Unmarshal(resp.Tickets, &ticketsMap); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if err = json.Unmarshal(resp.PendingTickets, &activeTicketMap); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":                resp.Code,
		"msg":                 resp.Msg,
		"tickets":             ticketsMap,
		"pending_tickets":     activeTicketMap,
		"pending_page_count":  resp.PendingPageCount,
		"finished_page_count": resp.FinishedPageCount,
	})
}

func HandleCheckIsUserHaveOpeningTickets(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.TicketHandleClient.CheckIsUserHaveOpeningTickets(sysContext.Background(), &ticketPb.CheckIsUserHaveOpeningTicketsRequest{
		UserId: userId,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":         resp.Code,
		"msg":          resp.Msg,
		"exist":        resp.Exist,
		"ticket_count": resp.TicketCount,
	})
}
