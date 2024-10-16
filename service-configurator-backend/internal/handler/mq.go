package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/model"
)

func HandleSaveMqConf(context *gin.Context) {
	var mqConfig model.MqConfig
	if err := context.ShouldBind(&mqConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}
	tempXml, err := xml.MarshalIndent(mqConfig, "", "  ")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "转换XML格式失败",
		})
		return
	}
	//log.Println(string(redisConfigXml))
	if err := etcd.RegisterConfig2Etcd("mq", string(tempXml)); err != nil {
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

func HandleGetMqConf(context *gin.Context) {
	//log.Println("HandleGetMqConf")
	var mqConfig model.MqConfig
	var err error
	str, err := etcd.GetConfigFromEtcd("mq")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取错误",
		})
		return
	}
	if err := xml.Unmarshal([]byte(str), &mqConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "反序列化错误" + err.Error(),
		})
		return
	}
	//log.Println(mqConfig)
	context.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"mq_config": mqConfig,
		"msg":       "查询成功",
	})
}
