package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/model"
)

func HandleSaveRedisConf(context *gin.Context) {
	var redisConfig model.RedisConfig
	if err := context.ShouldBind(&redisConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}
	//log.Println(redisConfig)
	redisConfigXml, err := xml.MarshalIndent(redisConfig, "", "  ")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "转换XML格式失败",
		})
		return
	}
	//log.Println(string(redisConfigXml))
	if err := etcd.RegisterConfig2Etcd("redis", string(redisConfigXml)); err != nil {
		//log.Println(err)
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
	//log.Println("HandleGetRedisConf")
	var redisConfig model.RedisConfig
	var err error
	str, err := etcd.GetConfigFromEtcd("redis")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取错误",
		})
		return
	}
	if err := xml.Unmarshal([]byte(str), &redisConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "反序列化错误" + err.Error(),
		})
		return
	}
	//log.Println(redisConfig)
	context.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"redis_config": redisConfig,
		"msg":          "查询成功",
	})
}
