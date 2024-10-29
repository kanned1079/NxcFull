package handler

import (
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
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
