package privilege

import (
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/orders"
	"NxcFull/nxc-backend/subscribePlan"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
type Group struct {
	Id        int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
*/

func HandleGetAllGroups(context *gin.Context) {
	pageData := &struct {
		Page int64 `form:"page" json:"page"`
		Size int64 `form:"size" json:"size"`
	}{}
	if err := context.ShouldBindQuery(pageData); err != nil {
		log.Println(err)
	}
	log.Println(pageData)
	// pass
	var privilegeGroups []Group
	type ResponseData struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		UserCount int64  `json:"user_count"`
		PlanCount int64  `json:"plan_count"`
	}
	var responseDataList []ResponseData

	// Query all groups from the database
	if err := dao.Db.Model(&Group{}).Find(&privilegeGroups).Error; err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "查询失败",
			"error": err.Error(),
		})
		return
	}

	// Populate responseDataList with group information
	for _, group := range privilegeGroups {
		temp := ResponseData{
			Id:   group.Id,
			Name: group.Name,
		}

		// Query PlanCount from Plan database
		var planCount int64
		if err := dao.Db.Model(&subscribePlan.Plan{}).Where("group_id = ?", group.Id).Count(&planCount).Error; err != nil {
			log.Println("Failed to get plan count:", err)
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":  http.StatusInternalServerError,
				"msg":   "查询计划数量失败",
				"error": err.Error(),
			})
			return
		}
		temp.PlanCount = planCount

		// Query UserCount from ActiveOrders database
		var userCount int64

		if err := dao.Db.Model(&orders.ActiveOrders{}).Where("group_id = ? && is_active = ?", group.Id, true).Count(&userCount).Error; err != nil {
			log.Println("Failed to get user count:", err)
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":  http.StatusInternalServerError,
				"msg":   "查询用户数量失败",
				"error": err.Error(),
			})
			return
		}
		temp.UserCount = userCount

		// Append temp to responseDataList
		responseDataList = append(responseDataList, temp)
	}

	log.Println(responseDataList)
	context.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"group_list": responseDataList,
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

func HandleUpdateGroup(context *gin.Context) {
	// put方法 请求体数据保存到
	putData := &struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	if err := context.ShouldBind(&putData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "解析失败",
			"error": err.Error(),
		})
	}
	//groupNewInfo := Group{
	//	//Id:   putData.Id,
	//	Name: putData.Name,
	//}
	if result := dao.Db.Model(&Group{}).Where("id = ?", putData.Id).Updates(&Group{Name: putData.Name}); result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"meg":   "保存到数据库失败",
			"error": result.Error.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "更新成功",
	})
}

func HandleDeleteGroup(context *gin.Context) {
	delData := &struct {
		Id int64 `json:"id" binding:"required"`
	}{}
	if err := context.ShouldBindJSON(&delData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	// 从数据库中删除对应 ID 的 Group
	if err := dao.Db.Delete(&Group{}, delData.Id).Error; err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "删除失败",
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除成功",
	})
}
