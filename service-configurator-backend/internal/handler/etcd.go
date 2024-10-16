package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service-configurator-backend/internal/etcd"
)

func HandleCheckEtcdClient(context *gin.Context) {
	if etcd.EtcdCli == nil {
		log.Println("Etcd client not initialized")
		context.JSON(http.StatusOK, gin.H{
			"code":       http.StatusBadRequest,
			"initialled": false,
			"message":    "Etcd客户端未初始化",
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":       http.StatusOK,
			"initialled": true,
			"message":    "Etcd客户端已初始化",
		})
		return
	}
}
