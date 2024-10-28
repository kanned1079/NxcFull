package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})

}
