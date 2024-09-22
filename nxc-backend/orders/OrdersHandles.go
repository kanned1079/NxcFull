package orders

import (
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/subscribePlan"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
//func GetActivePlanListByUserId(context *gin.Context) {
//	// 使用get方法 请求示例 http://localhost:8080/user/v1/plan/info/fetch?user_id=3
//	// 使用user_id查询查询Active表中是否有包含该用户id的列表 并且该条目的IsActive是否有效
//	// 取出计划Id，计划到期日期
//	// 使用计划Id查询Plan表 得到订阅计划名称
//	// 都查询成功后将返回一个数组
//	// [{plan_name, expiration_date}]
//	//
//	// 表
//	//type ActiveOrders struct {
//	//	ID             int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 订阅记录 ID
//	//	UserID         int64          `json:"user_id"`                              // 用户 ID
//	//	Email          string         `json:"email"`                                // 邮箱地址
//	//	PlanId         int64          `json:"plan_id"`                              // 购买的订阅计划 ID
//	//	KeyId          int64          `json:"key_id"`                               // 对应表x_keys的密钥Id
//	//	IsActive       bool           `json:"is_active"`                            // 订阅是否有效
//	//	IsUsed         bool           `json:"is_used"`                              // 订阅密钥是否被使用过
//	//	CreatedAt      time.Time      `json:"created_at"`                           // 创建日期
//	//	ExpirationDate *time.Time     `json:"expiration_date"`                      // 到期日期
//	//	UpdatedAt      time.Time      `json:"updated_at"`                           // 更新时间
//	//	DeletedAt      gorm.DeletedAt `json:"deleted_at"`                           // 删除时间
//	//}
//
//	// Plan表
//	//type Plan struct {
//	//	Id            int64   `json:"id" gorm:"primary_key"`     // 订阅计划ID
//	//	GroupId       int64   `json:"group_id"`                  // 群组Id
//	//	Name          string  `json:"name"`                      // 订阅名称
//	//	IsSale        bool    `json:"is_sale"`                   // 启用销售
//	//	IsRenew       bool    `json:"is_renew"`                  // 是否允许续费
//	//	CapacityLimit int64   `json:"capacity_limit"`            // 最大用户数量限制
//	//	Residue       int64   `json:"residue"`                   // 剩余订阅数量
//	//	Describe      string  `json:"describe" gorm:"type:TEXT"` // 订阅计划的备注
//	//	MonthPrice    float64 `json:"month_price"`               // 月付价格
//	//	QuarterPrice  float64 `json:"quarter_price"`             // 季付价格
//	//	HalfYearPrice float64 `json:"half_year_price"`           // 半年付价格
//	//	YearPrice     float64 `json:"year_price"`                // 年付价格
//	//	Sort          int64   `json:"sort"`                      // 排序顺序
//	//	CreatedAt time.Time      `json:"created_at"`
//	//	UpdatedAt time.Time      `json:"updated_at"`
//	//	DeletedAt gorm.DeletedAt `json:"deleted_at"`
//	//}
//
//}

func GetActivePlanListByUserId(context *gin.Context) {
	// Parse the user_id from the query parameters
	userID := context.Query("user_id")
	if userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	var activeOrders []ActiveOrders
	// Query the ActiveOrders table for the user_id and where IsActive is true
	if err := dao.Db.Where("user_id = ? AND is_active = ?", userID, true).Find(&activeOrders).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active orders"})
		return
	}

	if len(activeOrders) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "No active plans found for the user"})
		return
	}

	// Create a result array to hold plan_name and expiration_date
	var result []map[string]interface{}

	// Iterate through the active orders and fetch plan details
	for _, order := range activeOrders {
		var plan subscribePlan.Plan
		// Fetch the plan details using the PlanId
		if err := dao.Db.Where("id = ?", order.PlanId).First(&plan).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plan details"})
			return
		}

		// Add the plan_name and expiration_date to the result
		result = append(result, map[string]interface{}{
			"plan_name":       plan.Name,
			"expiration_date": order.ExpirationDate,
		})
	}

	// Return the result as a JSON response
	context.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"my_plans": result,
		"msg":      "获取数据成功",
	})
}
