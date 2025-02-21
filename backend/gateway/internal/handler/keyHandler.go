package handler

import (
	sysContext "context"
	"encoding/json"
	pb "gateway/internal/grpc/api/key/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	//log.Println("Unmarshal")
	//var keyMap []map[string]any
	//if json.Unmarshal(resp.Keys, &keyMap) != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	//log.Println(keyMap)
	//for k, v := range keyMap {
	//	log.Println(k, v)
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"my_keys":    json.RawMessage(resp.Keys),
		"page_count": resp.PageCount,
	})

}

func HandleGetAllMyActivationLogs(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	err, page, size := GetPage2SizeFromQuery(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.GetActivateLogByUserId(sysContext.Background(), &pb.GetActivateLogByUserIdRequest{
		UserId: userId,
		Page:   page,
		Size:   size,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var recordsMap []map[string]any
	if err := json.Unmarshal(resp.Log, &recordsMap); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"log":        recordsMap,
		"page_count": resp.PageCount,
	})
}

func HandleGetKeyDetailsById(context *gin.Context) {
	keyId, err := strconv.ParseInt(context.Query("key_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.GetKeyInfoById(sysContext.Background(), &pb.GetKeyInfoByIdRequest{
		KeyId: keyId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var detailsMap map[string]any
	err = json.Unmarshal(resp.Details, &detailsMap)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"details": detailsMap,
	})
}

func HandleDisableBindKey(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	activationId, err := strconv.ParseInt(context.Query("activation_id"), 10, 64)
	log.Println(activationId, userId, err)
	resp, err := grpcClient.KeyServicesClient.UnbindKey(sysContext.Background(), &pb.UnbindKeyRequest{
		UserId:   userId,
		RecordId: activationId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusInternalServerError,
			"msg":      err.Error(),
			"finished": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"finished": true,
	})
}

func HandleBindKeyToThirdExternalApp(context *gin.Context) {
	postData := &struct {
		Email         string `json:"email"`
		Password      string `json:"password"`
		Key           string `json:"key"`
		ClientVersion string `json:"client_version"`
		ClientId      string `json:"client_id"`
		OsType        string `json:"os_type"`
		Remark        string `json:"remark"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}

	if strings.Trim(postData.Email, " ") == "" || strings.Trim(postData.Password, " ") == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}
	log.Println(postData)

	resp, err := grpcClient.KeyServicesClient.BindOrActiveMyKey2App(sysContext.Background(), &pb.BindOrActiveMyKey2AppRequest{
		Email:         postData.Email,
		Password:      postData.Password,
		Key:           postData.Key,
		ClientVersion: postData.ClientVersion,
		ClientId:      postData.ClientId,
		OsType:        postData.OsType,
		Remark:        postData.Remark,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":               resp.Code,
		"msg":                resp.Msg,
		"significant_number": resp.SignificantNumber,
	})
}

func HandleAlterRemarkByUser(context *gin.Context) {
	postData := &struct {
		UserId   int64  `json:"user_id"`
		RecordId int64  `json:"record_id"`
		Remark   string `json:"remark"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.AlterKeyBindRemark(sysContext.Background(), &pb.AlterKeyBindRemarkRequest{
		UserId:   postData.UserId,
		RecordId: postData.RecordId,
		Remark:   postData.Remark,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func GetAllActivateLogByAdmin(context *gin.Context) {
	err, page, size := GetPage2SizeFromQuery(context)

	email := context.Query("email")
	valid, err := strconv.ParseBool(context.Query("valid"))
	sort := context.Query("sort")
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.GetActivateLogByAdmin(sysContext.Background(), &pb.GetActivateLogByAdminRequest{
		Page:  page,
		Size:  size,
		Email: email,
		Valid: valid,
		Sort:  sort,
	})

	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var logMap []map[string]any
	//err = json.Unmarshal(resp.Log, &logMap)
	//if err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"log":        json.RawMessage(resp.Log),
		"page_count": resp.PageCount,
	})
}

func HandleGetAllKeysByAdmin(context *gin.Context) {
	err, page, size := GetPage2SizeFromQuery(context)
	keyContent := context.Query("key_content")
	keyId, err := strconv.ParseInt(context.Query("key_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.GetAllKeysByAdminDesc(sysContext.Background(), &pb.GetAllKeysByAdminDescRequest{
		Page:             page,
		Size:             size,
		SearchKeyId:      keyId,
		SearchKeyContent: keyContent,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var keysMap []map[string]any
	//err = json.Unmarshal(resp.Keys, &keysMap)
	//if err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "gateway " + err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"keys":       json.RawMessage(resp.Keys),
		"page_count": resp.PageCount,
	})
}

func HandleAlterKeyIsValidByAdmin(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	keyId, err := strconv.ParseInt(context.Query("key_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request: " + err.Error(),
		})
		return
	}
	resp, err := grpcClient.KeyServicesClient.AlterKeyIsValidByAdmin(sysContext.Background(), &pb.AlterKeyIsValidByAdminRequest{
		UserId: userId,
		KeyId:  keyId,
	})

	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
