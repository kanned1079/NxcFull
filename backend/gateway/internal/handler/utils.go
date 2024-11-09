package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
