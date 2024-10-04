package handler

import (
	pb "NxcFull/backend/userServices/api/proto"
	sysContext "context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

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
	resp, err := grpcClient.Login(ctx, rpcReq)
	if err != nil {
		log.Println("grpc请求错误", err)
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "grpc调用失败",
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"isAuthed":  true,
		"msg":       "验证通过",
		"token":     resp.Token,
		"user_data": resp.UserData,
	})
}
