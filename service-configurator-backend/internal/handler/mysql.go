package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-configurator-backend/internal/etcd"
	"service-configurator-backend/internal/model"
)

func HandleSaveMysqlConf(context *gin.Context) {
	var mysqlConfig model.MysqlConfig
	if err := context.ShouldBind(&mysqlConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}
	//log.Println(mysqlConfig)
	tempXml, err := xml.MarshalIndent(mysqlConfig, "", "  ")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "转换XML格式失败",
		})
		return
	}
	//log.Println(string(redisConfigXml))
	if err := etcd.RegisterConfig2Etcd("mysql", string(tempXml)); err != nil {
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

func HandleGetMysqlConf(context *gin.Context) {
	//log.Println("HandleGetRedisConf")
	var mysqlConfig model.MysqlConfig
	var err error
	str, err := etcd.GetConfigFromEtcd("mysql")
	if err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取错误",
		})
		return
	}
	if err := xml.Unmarshal([]byte(str), &mysqlConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "反序列化错误" + err.Error(),
		})
		return
	}
	//log.Println(mysqlConfig)
	context.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"mysql_config": mysqlConfig,
		"msg":          "查询成功",
	})
}
