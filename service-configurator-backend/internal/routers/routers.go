package routers

import (
	"github.com/gin-gonic/gin"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/handler"
	"service-configurator-backend/internal/middleware"
)

func StartRoute() {
	r := gin.Default()

	r.Use(middleware.PassOptions) // 预检OPTIONS请求

	r.POST("/api/v1/config/etcd/init", etcd.HandleSaveEtcdConf)
	r.GET("/api/v1/config/etcd/check", handler.HandleCheckEtcdClient)

	r.POST("/api/v1/config/mysql", handler.HandleSaveMysqlConf)
	r.GET("/api/v1/config/mysql", handler.HandleGetMysqlConf)

	r.POST("/api/v1/config/redis", handler.HandleSaveRedisConf)
	r.GET("/api/v1/config/redis", handler.HandleGetRedisConf)

	r.POST("/api/v1/config/mq", handler.HandleSaveMqConf)
	r.GET("/api/v1/config/mq", handler.HandleGetMqConf)

	r.POST("/api/v1/config/server", handler.HandleSaveServerConf)
	r.GET("/api/v1/config/server", handler.HandleGetServerConf)
	if err := r.Run(":5001"); err != nil {
		panic(err)
	}

}
