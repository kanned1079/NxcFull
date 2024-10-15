package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PassOptions(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(http.StatusOK)
		return
	}
	context.Next()
}
