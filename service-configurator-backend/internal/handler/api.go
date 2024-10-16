package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/model"
)

func HandleSaveServerConf(context *gin.Context) {
	var apiServerConfig model.ApiServerConfig
	if err := context.ShouldBind(&apiServerConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}
	log.Println(apiServerConfig)
	tempXml, err := xml.MarshalIndent(apiServerConfig, "", "  ")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "转换XML格式失败",
		})
		return
	}
	//log.Println(string(redisConfigXml))
	if err := etcd.RegisterConfig2Etcd("apiServer", string(tempXml)); err != nil {
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

func HandleGetServerConf(context *gin.Context) {
	//log.Println("HandleGetMqConf")
	var apiServerConfig model.ApiServerConfig
	var err error
	str, err := etcd.GetConfigFromEtcd("apiServer")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取错误",
		})
		return
	}
	if err := xml.Unmarshal([]byte(str), &apiServerConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "反序列化错误" + err.Error(),
		})
		return
	}
	//log.Println(mqConfig)
	context.JSON(http.StatusOK, gin.H{
		"code":              http.StatusOK,
		"api_server_config": apiServerConfig,
		"msg":               "查询成功",
	})
}
