package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

func failOnRpcError(err error, response interface{}) error {
	if err != nil {
		return errors.New("rpc调用失败" + err.Error())
	}
	if response == nil {
		return errors.New("rpc调用失败无返回值")
	}
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域，实际生产环境需要配置更严格的跨域检查
	},
}

func getPage2Size(context *gin.Context) (err error, page int64, size int64) {
	page, err = strconv.ParseInt(context.Query("page"), 10, 64)
	size, err = strconv.ParseInt(context.Query("size"), 10, 64)
	return
}

func getUserIdAndOrderIdFromContextViaGetMethod(context *gin.Context) (err error, userId int64, orderId string) {
	orderId = context.Query("order_id")
	userIdStr := context.Query("user_id")
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if orderId == "" || userIdStr == "" || err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "提供的信息不正确",
		})
	}
	log.Println(userId, orderId, err)
	return
}

func getUserIdAndOrderIdFromContextViaPostMethod(context *gin.Context) (err error, userId int64, orderId string) {
	postData := &struct {
		UserId  int64  `json:"user_id"`
		OrderId string `json:"order_id"`
	}{}
	if err = context.BindJSON(postData); err != nil {
		err = errors.New("提供的信息不正确" + err.Error())
		return
	}
	userId = postData.UserId
	orderId = postData.OrderId
	return
}

func GetPage2SizeFromQuery(context *gin.Context) (err error, page int64, size int64) {
	page, err = strconv.ParseInt(context.Query("page"), 10, 64)
	if err != nil {
		return
	}
	size, err = strconv.ParseInt(context.Query("size"), 10, 64)
	if err != nil {
		return
	}
	return nil, page, size
}
