package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConnectionTestForAdmin 用于测试响应时间
func ConnectionTestForAdmin(context *gin.Context) {
	//time.Sleep(time.Second * 2)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
