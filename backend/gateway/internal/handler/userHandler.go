package handler

import (
	grpc2 "NxcFull/backend/gateway/internal/grpc"
	pb "NxcFull/backend/gateway/internal/grpc/api/proto"
	sysContext "context"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

var grpcClient = grpc2.NewClients()

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
	}
	//if err != nil || resp.Code != http.StatusOK {
	//	log.Println("grpc请求错误", err)
	//	context.JSON(http.StatusOK, gin.H{
	//		"code":     resp.Code,
	//		"isAuthed": resp.IsAuthed,
	//		"msg":      "grpc调用失败",
	//		"error":    err.Error(),
	//	})
	//	return
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code":      http.StatusOK,
	//	"isAuthed":  true,
	//	"msg":       "验证通过",
	//	"token":     resp.Token,
	//	"user_data": resp.UserData,
	//})
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

	context.JSON(http.StatusOK, gin.H{
		"code":          resp.Code,
		"is_registered": resp.IsRegistered,
		"msg":           resp.Msg,
	})

	//var newUserInfo user.User = user.User{
	//	Email:        postForm.Email,
	//	InviteUserID: postForm.InviteUserId,
	//	//LicenseExpiration: nil,
	//}
	//var newUserAuth user.Auth = user.Auth{
	//	Email:    postForm.Email,
	//	Password: postForm.Password,
	//}
	//result1 := dao.Db.Model(&user.User{}).Create(&newUserInfo)
	//if result1.Error != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "插入数据错误",
	//		"error": result1.Error.Error(),
	//	})
	//}
	//result2 := dao.Db.Model(&user.Auth{}).Create(&newUserAuth)
	//if result2.Error != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "插入验证数据错误",
	//		"error": result2.Error.Error(),
	//	})
	//}
	//if result1.Error == nil && result2.Error == nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusOK,
	//		"msg":  "新建用户成功",
	//	})
	//}
}
