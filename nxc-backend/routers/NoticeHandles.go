package routers

import (
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/settings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllNotices(context *gin.Context) {
	var notices []coupon.Notices
	result := dao.Db.Model(&settings.Notice{}).Find(&notices)
	if result.Error != nil {
		log.Print(result.Error)
	}
	log.Println("通知条数", len(notices))
	log.Println(notices)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": notices,
	})
}
