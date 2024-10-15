package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/model"
)

func HandleSaveRedisConf(context *gin.Context) {
	var redisConfig model.RedisConfig
	if err := context.ShouldBind(&redisConfig); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println(redisConfig)
	redisConfigJson, err := json.Marshal(redisConfig)
	if err != nil {
		log.Println(err)
		return
	}
	if err := etcd.RegisterConfig2Etcd("mysql", string(redisConfigJson)); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "注册设置失败" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "注册设置成功",
	})
}

func HandleGetRedisConf(context *gin.Context) {

}
