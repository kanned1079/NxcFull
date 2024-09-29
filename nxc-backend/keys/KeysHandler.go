package keys

import (
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/orders"
	"NxcFull/nxc-backend/subscribePlan"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllUserKeys(context *gin.Context) {
	idStr := context.Query("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取用户id失败",
		})
		return
	}

	// 查询用户的所有有效订单
	var activeOrders []orders.ActiveOrders
	if result := dao.Db.Where("user_id = ?", id).Find(&activeOrders); result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "获取用户的有效订单错误",
			"error": result.Error.Error(),
		})
		return
	}

	// 创建返回的 key 列表
	var responseDataList []struct {
		OrderId        int64  `json:"order_id"`
		KeyId          int64  `json:"key_id"`
		Key            string `json:"key"`
		PlanId         int64  `json:"plan_id"`
		PlanName       string `json:"plan_name"`
		IsActive       bool   `json:"is_active"`
		IsUsed         bool   `json:"is_used"`
		ClientId       string `json:"client_id"`
		CreatedAt      string `json:"created_at"`
		ExpirationDate string `json:"expiration_date"`
	}

	for _, order := range activeOrders {
		// 根据 key_id 查询密钥信息
		var key Keys
		if result := dao.Db.Where("id = ?", order.KeyId).First(&key); result.Error != nil {
			log.Println(result.Error)
			continue
		}

		// 根据 plan_id 查询计划名称
		var plan subscribePlan.Plan
		if result := dao.Db.Where("id = ?", order.PlanId).First(&plan); result.Error != nil {
			log.Println(result.Error)
			continue
		}

		// 组装订单和密钥的信息
		responseData := struct {
			OrderId        int64  `json:"order_id"`
			KeyId          int64  `json:"key_id"`
			Key            string `json:"key"`
			PlanId         int64  `json:"plan_id"`
			PlanName       string `json:"plan_name"`
			IsActive       bool   `json:"is_active"`
			IsUsed         bool   `json:"is_used"`
			ClientId       string `json:"client_id"`
			CreatedAt      string `json:"created_at"`
			ExpirationDate string `json:"expiration_date"`
		}{
			OrderId:  order.ID,
			KeyId:    order.KeyId,
			Key:      key.Key,
			PlanId:   order.PlanId,
			PlanName: plan.Name,
			IsActive: order.IsActive,
			IsUsed:   order.IsUsed,
			ClientId: key.ClientId,
			//CreatedAt:      key.CreatedAt,
			CreatedAt:      key.CreatedAt.Format("2006-01-02 15:04:05"),
			ExpirationDate: order.ExpirationDate.Format("2006-01-02 15:04:05"),
		}

		// 将每个订单信息添加到 responseDataList
		responseDataList = append(responseDataList, responseData)
	}

	// 返回结果
	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"key_list": responseDataList,
	})
}
