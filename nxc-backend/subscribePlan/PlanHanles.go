package subscribePlan

import (
	"NxcFull/nxc-backend/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllPlans(context *gin.Context) {
	var plans []Plan
	//if result := dao.Db.Model(&Plan{}).Find(&plans); result.Error != nil {
	if result := dao.Db.Model(&Plan{}).Order("sort ASC").Find(&plans); result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取订阅列表失败",
		})
	}
	log.Println(plans)
	context.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"plans": plans,
	})
}

func HandleAddNewPlan(context *gin.Context) {
	postData := &struct {
		Name          string  `json:"name"`
		Describe      string  `json:"describe"`
		IsSale        bool    `json:"is_sale"`
		IsRenew       bool    `json:"is_renew"`
		GroupId       int64   `json:"group_id"`
		CapacityLimit int64   `json:"capacity_limit"`
		MonthPrice    float64 `json:"month_price"`
		QuarterPrice  float64 `json:"quarter_price"`
		HalfYearPrice float64 `json:"half_year_price"`
		YearPrice     float64 `json:"year_price"`
		Sort          int64   `json:"sort"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
	}
	log.Println(postData)
	var newPlan = Plan{
		Name:          postData.Name,
		Describe:      postData.Describe,
		IsSale:        postData.IsSale,
		IsRenew:       postData.IsRenew,
		GroupId:       postData.GroupId,
		CapacityLimit: postData.CapacityLimit,
		Residue:       postData.CapacityLimit,
		MonthPrice:    postData.MonthPrice,
		QuarterPrice:  postData.QuarterPrice,
		HalfYearPrice: postData.HalfYearPrice,
		YearPrice:     postData.YearPrice,
		// 以下Sort用于在前端进行排序 优先级从低到高，添加新订阅时，它的优先级为上一个的优先级+1，如果表内为空不存在数据，则第一个添加的sort为1000
		// Sort:
		Sort: postData.Sort,
	}
	if result := dao.Db.Create(&newPlan); result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "插入数据库失败",
			"error": result.Error.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "插入数据成功",
	})
}

//func HandleAddNewPlan(context *gin.Context) {
//	postData := &struct {
//		Name          string  `json:"name"`
//		Describe      string  `json:"describe"`
//		IsSale        bool    `json:"is_sale"`
//		IsRenew       bool    `json:"is_renew"`
//		GroupId       int64   `json:"group_id"`
//		CapacityLimit int64   `json:"capacity_limit"`
//		MonthPrice    float64 `json:"month_price"`
//		QuarterPrice  float64 `json:"quarter_price"`
//		HalfYearPrice float64 `json:"half_year_price"`
//		YearPrice     float64 `json:"year_price"`
//	}{}
//	if err := context.ShouldBind(&postData); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{
//			"code":  http.StatusBadRequest,
//			"msg":   "绑定数据失败",
//			"error": err.Error(),
//		})
//		return
//	}
//
//	log.Println(postData)
//
//	// 查询当前最大 Sort 值
//	var maxSort int64
//	err := dao.Db.Model(&Plan{}).Select("COALESCE(MAX(sort), 999)").Scan(&maxSort).Error
//	if err != nil {
//		log.Println("查询最大 Sort 值失败:", err)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "查询排序优先级失败",
//			"error": err.Error(),
//		})
//		return
//	}
//
//	var newPlan = Plan{
//		Name:          postData.Name,
//		Describe:      postData.Describe,
//		IsSale:        postData.IsSale,
//		IsRenew:       postData.IsRenew,
//		GroupId:       postData.GroupId,
//		CapacityLimit: postData.CapacityLimit,
//		Residue:       postData.CapacityLimit,
//		MonthPrice:    postData.MonthPrice,
//		QuarterPrice:  postData.QuarterPrice,
//		HalfYearPrice: postData.HalfYearPrice,
//		YearPrice:     postData.YearPrice,
//		Sort:          maxSort + 1,  // 设置新的 Sort 值
//	}
//
//	// 将新 Plan 插入数据库
//	if result := dao.Db.Create(&newPlan); result.Error != nil {
//		log.Println("插入数据库失败:", result.Error)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "插入数据库失败",
//			"error": result.Error.Error(),
//		})
//		return
//	}
//
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "插入数据成功",
//	})
//}
//
