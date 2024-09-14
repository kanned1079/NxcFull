package subscribePlan

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllPlans(context *gin.Context) {
	var plans []Plan
	if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取订阅列表失败",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"plans": plans,
	})
}
