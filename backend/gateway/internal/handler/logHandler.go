package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/logs/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetSystemStatus(context *gin.Context) {
	resp, err := grpcClient.LogServiceClient.GetServerLiveStatus(sysContext.Background(), &pb.GetServerLiveStatusRequest{})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var apiLogList []map[string]interface{}
	if err = json.Unmarshal(resp.ApiLogList, &apiLogList); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"msg":          "success",
		"status200":    resp.Status200,
		"status404":    resp.Status404,
		"status500":    resp.Status500,
		"login_req":    resp.LoginReq,
		"reg_req":      resp.RegisterReq,
		"api_log_list": apiLogList,
		"page_count":   0,
	})
}
