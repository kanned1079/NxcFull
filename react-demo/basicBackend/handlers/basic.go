package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Comments struct {
	Id      int64  `json:"id"`
	Comment string `json:"comment"`
}

func HandleTestComment17(context *gin.Context) {
	var comments = make([]Comments, 100)
	for i := 0; i < 100; i++ {
		comments[i] = Comments{
			Id:      int64(i + 1),
			Comment: fmt.Sprintf("测试评论%d", i+1),
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"msg":      "success",
		"comments": comments,
	})
}
