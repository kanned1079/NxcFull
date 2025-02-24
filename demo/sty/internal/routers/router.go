package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ProtocolAllowance() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	}
}

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var userList []User = []User{
	{Id: 1, Email: "user1@example.com", Name: "User One", Password: "password1", Balance: 100.5, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: 2, Email: "user2@example.com", Name: "User Two", Password: "password2", Balance: 200.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func paginate(users []User, page, size int) ([]User, int) {
	start := (page - 1) * size
	end := start + size
	if start >= len(users) {
		return []User{}, 0
	}
	if end > len(users) {
		end = len(users)
	}
	pageCount := (len(users) + size - 1) / size
	return users[start:end], pageCount
}

func RunServer(id int64) {
	routerInst := NewApiInstance(id)
	routerInst.Router.Use(gin.Logger())
	routerInst.Router.Use(ProtocolAllowance())

	// 获取所有用户并分页
	routerInst.Router.GET("/user", func(context *gin.Context) {
		pageStr := context.DefaultQuery("page", "1")
		sizeStr := context.DefaultQuery("size", "10")
		log.Println(pageStr, sizeStr)
		page, _ := strconv.Atoi(pageStr)
		size, _ := strconv.Atoi(sizeStr)

		users, pageCount := paginate(userList, page, size)
		context.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"users":      users,
			"page":       page,
			"size":       size,
			"page_count": pageCount,
		})
	})

	// 创建用户
	routerInst.Router.POST("/user", func(context *gin.Context) {
		var newUser User
		if err := context.ShouldBindJSON(&newUser); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid input",
			})
			return
		}
		newUser.Id = len(userList) + 1
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()
		userList = append(userList, newUser)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": newUser,
		})
	})

	// 更新用户信息
	routerInst.Router.PUT("/user", func(context *gin.Context) {
		var updatedUser User
		if err := context.ShouldBind(&updatedUser); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Invalid input",
			})
			return
		}

		for i, user := range userList {
			if user.Id == updatedUser.Id {
				userList[i].Email = updatedUser.Email
				userList[i].Name = updatedUser.Name
				userList[i].Password = updatedUser.Password
				userList[i].Balance = updatedUser.Balance
				userList[i].UpdatedAt = time.Now()
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"data": userList[i],
				})
				return
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusNotFound,
			"error": "User not found",
		})
	})

	// 删除用户
	routerInst.Router.DELETE("/user", func(context *gin.Context) {
		id := context.DefaultQuery("id", "")
		if id == "" {
			context.JSON(http.StatusOK, gin.H{
				"code":  http.StatusBadRequest,
				"error": "Missing user id",
			})
			return
		}
		for i, user := range userList {
			if fmt.Sprintf("%d", user.Id) == id {
				userList = append(userList[:i], userList[i+1:]...)
				context.JSON(http.StatusOK, gin.H{
					"code":    http.StatusOK,
					"message": "User deleted",
				})
				return
			}
		}
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusNotFound,
			"error": "User not found",
		})
	})

	// 启动服务器
	if err := routerInst.Router.Run("localhost:5001"); err != nil {
		log.Print("服务器启动失败")
	}
}
