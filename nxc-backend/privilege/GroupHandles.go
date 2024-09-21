package privilege

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllGroups(context *gin.Context) {
	var privilegeGroups []Group
	if err := dao.Db.Model(&Group{}).Find(&privilegeGroups).Error; err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "查询失败",
			"error": err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"group_list": privilegeGroups,
	})
}

func HandleAddNewGroup(context *gin.Context) {
	postData := &struct {
		GroupName string `json:"group_name"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
	}
	log.Println(postData)
	var newPrivilegeGroup Group
	newPrivilegeGroup.Name = postData.GroupName
	if err := dao.Db.Model(&Group{}).Create(&newPrivilegeGroup).Error; err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "插入数据失败",
			"error": err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "插入数据成功",
	})
}
