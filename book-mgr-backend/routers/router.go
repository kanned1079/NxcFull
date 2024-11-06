package routers

import (
	"book-mgr-backend/handler/admin"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
}

func (a *App) RunServer() {
	r := gin.Default()

	r.GET("books")

	adminGroup := r.Group("/api/admin")
	{
		adminGroup.POST("login", admin.HandleAdminLogin)
		adminGroup.GET("book", admin.HandleGetAllBooks_Admin)
		adminGroup.POST("book", admin.HandleAddBook_Admin)
		adminGroup.PUT("book", admin.HandleUpdateBook_Admin)
		adminGroup.DELETE("book", admin.HandleDeleteBook_Admin)
	}

	userGroup := r.Group("/api/user")
	{
		userGroup.POST("login")
	}

	if err := r.Run("localhost:8080"); err != nil {
		log.Panicln("端口可能已被占用 服务器启动失败" + err.Error())
	}
}
