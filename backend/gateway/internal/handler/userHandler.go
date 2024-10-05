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
