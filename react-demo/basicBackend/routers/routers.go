package routers

import (
	"basicBackend/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	err error
	App *Gateway
)

func (a *Gateway) RunApiGateway() {
	log.Println("App Id: ", a.id)
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	router.GET("/api/v1/comments17", handlers.HandleTestComment17)

	if err = router.Run("localhost:8080"); err != nil {
		log.Println("端口可能已经被占用 err: ", err)
	}
}
