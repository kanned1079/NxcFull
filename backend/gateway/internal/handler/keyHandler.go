package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/key/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleGetAllMyKeys(context *gin.Context) {
	var userId int64
	var page int64
	var size int64
	var err error
	if userId, err = strconv.ParseInt(context.Query("user_id"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
	}
	if page, err = strconv.ParseInt(context.Query("page"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
		return
	}
	if size, err = strconv.ParseInt(context.Query("size"), 10, 64); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "解析数据失败",
		})
		return
	}
	log.Println("gateway: ", userId, page, size)
	resp, err := grpcClient.KeyServicesClient.GetAllMyKeysByUserIdAsc(sysContext.Background(), &pb.GetAllMyKeysByUserIdAscRequest{
		UserId: userId,
		Page:   page,
		Size:   size,
	})
	log.Println("rpc调用结束")
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println("Unmarshal")
	var keyMap []map[string]any
	if json.Unmarshal(resp.Keys, &keyMap) != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(keyMap)
	for k, v := range keyMap {
		log.Println(k, v)
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"my_keys": keyMap,
	})

}
