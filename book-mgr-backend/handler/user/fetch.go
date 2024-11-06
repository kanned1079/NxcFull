package user

import (
	"book-mgr-backend/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllBooks_User(context *gin.Context) {
	err, page, size := handler.GetPage2SizeFormQueryParams(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "缺少查询参数",
		})
		return
	}
	log.Println(page, size)

}
