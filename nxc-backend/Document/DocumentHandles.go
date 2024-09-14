package Document

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllDocument(context *gin.Context) {
	var documentsList []Document
	if result := dao.Db.Model(&Document{}).Find(&documentsList); result.Error != nil {
		log.Println("获取所有文档失败")
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取所有文档失败",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"documents": documentsList,
	})
}
