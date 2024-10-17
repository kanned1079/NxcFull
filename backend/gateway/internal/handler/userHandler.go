package handler

import (
	sysContext "context"
	"gateway/internal/grpc"
	pb "gateway/internal/grpc/api/user/proto"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

var grpcClient = grpc.NewClients()

func HandleUserLogin(context *gin.Context) {
	var req = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}{}
	if err := context.ShouldBind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//log.Println(req)
	log.Println("用户输入凭证 ", req.Email, req.Password)
	// 在这里发送rpc请求
	// 构造grpc请求
	rpcReq := &pb.UserLoginRequest{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
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
