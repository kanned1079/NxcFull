package handler

import (
	sysContext "context"
	"gateway/internal/grpc"
	pb "gateway/internal/grpc/api/user/proto"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"strconv"

	//"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

var grpcClient = grpc.NewClients()

func HandleUserLogin(context *gin.Context) {
	var req = struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		TwoFaCode string `json:"two_fa_code"`
		Role      string `json:"role"`
	}{}
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Println(req)
	log.Println("用户输入凭证 ", req.Email, req.Password, req.TwoFaCode, req.Role)
	// 在这里发送rpc请求
	// 构造grpc请求
	rpcReq := &pb.UserLoginRequest{
		Email:     req.Email,
		Password:  req.Password,
		TwoFaCode: req.TwoFaCode,
		Role:      req.Role,
	}
	// 调用grpc的login方法
	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), time.Second)
	defer cancel()
	//resp, err := GrpcClients.UserServiceClient.Login()
	resp, err := grpcClient.UserServiceClient.Login(ctx, rpcReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"isAuthed":  resp.IsAuthed,
		"msg":       resp.Msg,
		"token":     resp.Token,
		"user_data": resp.UserData,
	})
	//context.JSON(http.StatusOK, gin.H{
	//	"code":      http.StatusOK,
	//	"isAuthed":  true,
	//	"msg":       resp.Msg,
	//	"token":     resp.Token,
	//	"user_data": resp.UserData,
	//})

}

func HandleUserRegister(context *gin.Context) {
	postForm := &struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		InviteUserId string `json:"invite_user_id"`
	}{}
	if err := context.ShouldBind(&postForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "数据绑定错误",
			"error": err.Error(),
		})
	}
	log.Println("最终注册", postForm)

	// 调用rpc方法
	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	defer cancel()
	resp, err := grpcClient.UserServiceClient.Register(ctx, &pb.UserRegisterRequest{
		Email:        postForm.Email,
		Password:     postForm.Password,
		InviteUserId: postForm.InviteUserId,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":          resp.Code,
		"is_registered": resp.IsRegistered,
		"msg":           resp.Msg,
	})

}

func HandleCheckPreviousPassword(context *gin.Context) {
	// 绑定 JSON 请求体中的数据
	postData := &struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"` // 已经哈希过的密码
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
		return
	}

	log.Println(postData)

	resp, err := grpcClient.UserServiceClient.CheckPreviousPassword(sysContext.Background(), &pb.CheckPreviousPasswordRequest{
		Email:       postData.Email,
		OldPassword: postData.OldPassword,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	log.Println(resp)

	// 如果密码验证通过
	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"verified": resp.Verified,
		"msg":      resp.Msg,
	})
}

func HandleApplyNewPassword(context *gin.Context) {
	// 绑定 JSON 请求体中的数据
	postData := &struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"` // 新密码已经哈希过
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
		return
	}

	log.Println(postData)

	resp, err := grpcClient.UserServiceClient.ApplyNewPassword(sysContext.Background(), &pb.ApplyNewPasswordRequest{
		Email:       postData.Email,
		NewPassword: postData.NewPassword,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	log.Println(resp)

	//log.Println("修改密码成功")
	// 密码更新成功
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"updated": resp.Updated,
		"msg":     resp.Msg,
	})
}

func HandleSetup2FA(context *gin.Context) {
	postData := &struct {
		Id    int64  `json:"id"`
		Email string `json:"email"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
		return
	}
	log.Println(postData)
	resp, err := grpcClient.UserServiceClient.Setup2FA(sysContext.Background(), &pb.Setup2FARequest{
		Id:    postData.Id,
		Email: postData.Email,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务调用失败无返回值",
		})
		return
	}
	log.Println(resp)
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		"url":  resp.Url,
	})
}

func HandleTest2FA(context *gin.Context) {
	postData := &struct {
		Id        int64  `json:"id"`
		Email     string `json:"email"`
		TwoFaCode string `json:"two_fa_code"`
	}{}
	if err := context.ShouldBind(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println(postData)
	resp, err := grpcClient.UserServiceClient.Test2FA(sysContext.Background(), &pb.Test2FARequest{
		Id:        postData.Id,
		Email:     postData.Email,
		TwoFaCode: postData.TwoFaCode,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务器调用失败无返回值",
		})
		return
	}
	log.Println(resp)
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

// HandleGet2FAStatus GET请求查询是否有使用2FA
func HandleGet2FAStatus(context *gin.Context) {
	//postData := &struct {
	//	Id    int64  `json:"id"`
	//	Email string `json:"email"`
	//}{}
	//if err := context.ShouldBind(&postData); err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	tempIdStr := context.Query("id")
	id, err := strconv.Atoi(tempIdStr)
	log.Println(id)

	resp, err := grpcClient.UserServiceClient.Get2FAStatus(sysContext.Background(), &pb.Get2FAStatusRequest{
		Id: int64(id),
		//Email: postData.Email,
	})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc调用失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"enabled": resp.Enabled,
	})
}

func HandleCancelSetup2FA(context *gin.Context) {
	//idStr := context.Query("id")
	email := context.Query("email")
	resp, err := grpcClient.UserServiceClient.CancelSetup2FA(sysContext.Background(), &pb.CancelSetup2FARequest{
		Email: email,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "rpc服务器调用失败无返回值",
		})
		return
	}
	log.Println(resp)
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func HandleDisable2FA(context *gin.Context) {
	email := context.Query("email")
	resp, err := grpcClient.UserServiceClient.Disable2FA(sysContext.Background(), &pb.Disable2FARequest{
		Email: email,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"disabled": false,
			"msg":      err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusInternalServerError,
			"disabled": false,
			"msg":      "rpc服务器调用失败无返回值",
		})
		return
	}
	log.Println(resp)
	context.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"disabled": resp.Disabled,
		"msg":      resp.Msg,
	})
}

func HandleUpdateUserInfo(context *gin.Context) {
	paramsUserId := context.Query("user_id")
	var userId int64
	var err error
	if paramsUserId == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请提供用户Id",
		})
		return
	}
	userId, err = strconv.ParseInt(paramsUserId, 10, 64)
	log.Println("查询的用户id：", userId)
	resp, err := grpcClient.UserServiceClient.UpdateUserInfo(sysContext.Background(), &pb.UpdateUserInfoRequest{
		UserId: userId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//if err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "rpc调用错误" + err.Error(),
	//	})
	//	return
	//}
	//if resp == nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "rpc调用错误无返回值",
	//	})
	//	return
	//}
	var userMap map[string]interface{}
	err = json.Unmarshal(resp.UserInfo, &userMap)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"msg":       resp.Msg,
		"user_info": userMap,
	})
}
