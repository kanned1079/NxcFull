package routers

import (
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllNotices(context *gin.Context) {
	var notices []coupon.PublicNotices
	result := dao.Db.Model(&coupon.PublicNotices{}).Where("deleted_at is null").Find(&notices)
	if result.Error != nil {
		log.Print(result.Error)
	}
	//log.Println("通知条数", len(notices))
	//log.Println(notices)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"notices": notices,
		"msg":     "",
	})
}

type noticeFormData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	ImgUrl  string `json:"img_url"`
}

func HandleAddNotice(context *gin.Context) {
	var fromData noticeFormData
	if err := context.ShouldBind(&fromData); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusInternalServerError,
			"msg:": err.Error(),
		})
	}
	log.Println(fromData)
	var notice coupon.PublicNotices
	notice.Title = fromData.Title
	notice.Content = fromData.Content
	notice.Tags = fromData.Tags
	notice.ImgUrl = fromData.ImgUrl
	if result := dao.Db.Model(&coupon.PublicNotices{}).Create(&notice); result.Error != nil {
		log.Println("新建通知错误", result.Error)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func HandleDeleteNotice(context *gin.Context) {
	deleteNotice := &struct {
		NoticeId int `json:"notice_id"`
	}{}
	var err error
	if deleteNotice.NoticeId, err = strconv.Atoi(context.Query("notice_id")); err != nil {
		log.Println(err)
	}

	log.Println("消息id: ", deleteNotice.NoticeId)

	if err = dao.Db.Model(&coupon.PublicNotices{}).Where("id = ?", deleteNotice.NoticeId).Delete(&coupon.PublicNotices{}).Error; err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
			"msg":   "删除失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "删除成功",
		})
	}
}
