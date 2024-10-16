package etcd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/client/v3"
	"log"
	"net/http"
	"service-configurator-backend/internal/model"
	"strconv"
	"time"
)

var (
	EtcdCli *clientv3.Client
	err     error
)

func InitEtcdCli() error {
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", model.MyEtcdConfig.Endpoints, strconv.Itoa(int(model.MyEtcdConfig.Port)))},
		DialTimeout: 5 * time.Second,
		Username:    model.MyEtcdConfig.Username,
		Password:    model.MyEtcdConfig.Password,
	})
	if err != nil {
		log.Println("Error creating etcd client", err)
		return err
	}
	//defer EtcdCli.Close()

	return nil
}

func HandleSaveEtcdConf(context *gin.Context) {
	//var mqConfig model.MqConfig
	if err := context.ShouldBind(&model.MyEtcdConfig); err != nil {
		//log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "绑定数据失败",
		})
		return
	}

	if err := InitEtcdCli(); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "初始化Etcd失败：" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "初始化Etcd成功",
	})
}
