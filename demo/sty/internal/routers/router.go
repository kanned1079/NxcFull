package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ProtocolAllowance() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	}
}

func RunServer(id int64) {
	routerInst := NewApiInstance(id)
	routerInst.Router.Use(ProtocolAllowance())

	routerInst.Router.GET("/user", func(context *gin.Context) {
		name := context.Query("name")

		log.Println("name: ", name)

		//time.Sleep(time.Second * 2)

		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "ok",
			"name": name,
		})
	})

	if err := routerInst.Router.Run("localhost:5001"); err != nil {
		log.Print("服务器启动失败")
	}
}
