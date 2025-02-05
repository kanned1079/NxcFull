package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/logs/proto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleGetSystemStatus(context *gin.Context) {
	searchCode, err := strconv.ParseInt(context.Query("code"), 10, 64)
	err, page, size := GetPage2SizeFromQuery(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "page and size not found",
		})
		return
	}
	resp, err := grpcClient.LogServiceClient.GetServerLiveStatus(sysContext.Background(), &pb.GetServerLiveStatusRequest{
		Page: page,
		Size: size,
		Code: int32(searchCode),
	})
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
	context.SecureJSON(http.StatusOK, gin.H{
		"code":                 http.StatusOK,
		"msg":                  "success",
		"status200":            resp.Status200,
		"status404":            resp.Status404,
		"status500":            resp.Status500,
		"login_req":            resp.LoginReq,
		"reg_req":              resp.RegisterReq,
		"api_log_list":         apiLogList,
		"page_count":           resp.PageSize,
		"log_table_size":       resp.TableSize,
		"log_table_rows_count": resp.LogTableRowsCount,
	})
}

func HandleClearPreviousLog(context *gin.Context) {
	resp, err := grpcClient.LogServiceClient.ClearPreviousApiLog(sysContext.Background(), &pb.ClearPreviousApiLogRequest{})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":         resp.Code,
		"rows_deleted": resp.RowDeleted,
		"msg":          resp.Msg,
	})
}
