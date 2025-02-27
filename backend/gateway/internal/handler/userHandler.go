package handler

import (
	sysContext "context"
	pb "gateway/internal/grpc/api/user/proto"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"log"
	"strconv"

	"net/http"
	"time"
)

//func HandleUserLogin(context *gin.Context) {
//	var req = struct {
//		Email     string `json:"email"`
//		Password  string `json:"password"`
//		TwoFaCode string `json:"two_fa_code"`
//		Role      string `json:"role"`
//	}{}
//	if err := context.ShouldBind(&req); err != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code": http.StatusBadRequest,
//			"msg":  err.Error(),
//		})
//		return
//	}
//	rpcReq := &pb.UserLoginRequest{
//		Email:     req.Email,
//		Password:  req.Password,
//		TwoFaCode: req.TwoFaCode,
//		Role:      req.Role,
//	}
//	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), time.Second)
//	defer cancel()
//	resp, err := grpcClient.UserServiceClient.Login(ctx, rpcReq)
//	if err = failOnRpcError(err, resp); err != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":     http.StatusInternalServerError,
//			"isAuthed": false,
//			"msg":      err.Error(),
//		})
//	}
//
//	log.Println("用户登录信息", string(resp.UserData))
//
//	// resp.UserData 如果这个[]byte为空 那么就不需要反序列化用户信息
//
//	var userMap map[string]interface{}
//
//	if err := json.Unmarshal(resp.UserData, &userMap); err != nil {
//		context.JSON(http.StatusOK, gin.H{
//			"code":     http.StatusInternalServerError,
//			"isAuthed": false,
//			"msg":      "Error format json" + err.Error(),
//		})
//		return
//	} else {
//		context.JSON(http.StatusOK, gin.H{
//			"code":      resp.Code,
//			"isAuthed":  resp.IsAuthed,
//			"msg":       resp.Msg,
//			"token":     resp.Token,
//			"user_data": userMap,
//		})
//	}
//}

func HandleUserLogin(context *gin.Context) {
	// 定义请求参数结构体，并解析客户端传入的 JSON 数据
	var req = struct {
		Email     string `json:"email"`       // 用户邮箱
		Password  string `json:"password"`    // 用户密码
		TwoFaCode string `json:"two_fa_code"` // 用户的 2FA 验证码（如果有）
		Role      string `json:"role"`        // 用户角色（如管理员或普通用户）
	}{}

	// 解析客户端传来的请求体数据
	if err := context.ShouldBind(&req); err != nil {
		// 如果解析失败，返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest, // HTTP 状态码 400（客户端请求有误）
			"msg":  err.Error(),           // 返回错误信息
		})
		return
	}

	// 创建 gRPC 请求对象，将客户端数据封装为 protobuf 消息
	rpcReq := &pb.UserLoginRequest{
		Email:     req.Email,     // 用户邮箱
		Password:  req.Password,  // 用户密码
		TwoFaCode: req.TwoFaCode, // 2FA 验证码
		Role:      req.Role,      // 用户角色
	}

	// 设置 gRPC 请求的上下文和超时时间
	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), time.Second)
	defer cancel() // 确保超时后释放资源

	// 调用远程 gRPC 服务的 Login 方法进行用户登录逻辑
	resp, err := grpcClient.UserServiceClient.Login(ctx, rpcReq)
	if err = failOnRpcError(err, resp); err != nil {
		// 如果 gRPC 调用失败，返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusInternalServerError, // HTTP 状态码 500（服务器内部错误）
			"isAuthed": false,                          // 登录失败
			"msg":      err.Error(),                    // 错误信息
		})
		return
	}

	// 日志记录：打印用户登录返回的原始用户信息（UserData）
	log.Println("用户登录信息", string(resp.UserData))

	// 检查 gRPC 返回的用户数据是否为空
	if len(resp.UserData) == 0 {
		// 如果用户数据为空，则直接返回登录响应
		context.JSON(http.StatusOK, gin.H{
			"code":      resp.Code,     // gRPC 返回的状态码
			"isAuthed":  resp.IsAuthed, // 是否认证成功
			"msg":       resp.Msg,      // 消息内容
			"user_data": nil,           // 用户数据为空
		})
		return
	}

	// 定义用于解析用户信息的 Map
	var userMap map[string]interface{}

	// 反序列化用户数据（JSON 格式）到 map
	if err := json.Unmarshal(resp.UserData, &userMap); err != nil {
		// 如果反序列化失败，返回错误信息
		context.JSON(http.StatusOK, gin.H{
			"code":     http.StatusInternalServerError,      // HTTP 状态码 500（服务器内部错误）
			"isAuthed": false,                               // 登录失败
			"msg":      "Error format json: " + err.Error(), // 错误信息
		})
		return
	}

	// 返回登录成功的响应
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,     // gRPC 返回的状态码
		"isAuthed":  resp.IsAuthed, // 是否认证成功
		"msg":       resp.Msg,      // 消息内容
		"token":     resp.Token,    // 登录 token
		"user_data": userMap,       // 反序列化后的用户信息
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
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	//log.Println("最终注册", postForm)

	// 调用rpc方法
	ctx, cancel := sysContext.WithTimeout(sysContext.Background(), 5*time.Second)
	defer cancel()
	resp, err := grpcClient.UserServiceClient.Register(ctx, &pb.UserRegisterRequest{
		Email:        postForm.Email,
		Password:     postForm.Password,
		InviteUserId: postForm.InviteUserId,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
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

	//log.Println(postData)

	resp, err := grpcClient.UserServiceClient.CheckPreviousPassword(sysContext.Background(), &pb.CheckPreviousPasswordRequest{
		Email:       postData.Email,
		OldPassword: postData.OldPassword,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)

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

	//log.Println(postData)

	resp, err := grpcClient.UserServiceClient.ApplyNewPassword(sysContext.Background(), &pb.ApplyNewPasswordRequest{
		Email:       postData.Email,
		NewPassword: postData.NewPassword,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)

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
	//log.Println(postData)
	resp, err := grpcClient.UserServiceClient.Setup2FA(sysContext.Background(), &pb.Setup2FARequest{
		Id:    postData.Id,
		Email: postData.Email,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)
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
	//log.Println(postData)
	resp, err := grpcClient.UserServiceClient.Test2FA(sysContext.Background(), &pb.Test2FARequest{
		Id:        postData.Id,
		Email:     postData.Email,
		TwoFaCode: postData.TwoFaCode,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)
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
	//log.Println(id)

	resp, err := grpcClient.UserServiceClient.Get2FAStatus(sysContext.Background(), &pb.Get2FAStatusRequest{
		Id: int64(id),
		//Email: postData.Email,
	})
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
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
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)
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
	if err = failOnRpcError(err, resp); err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//log.Println(resp)
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
			"msg":  "no user_id provided",
		})
		return
	}
	userId, err = strconv.ParseInt(paramsUserId, 10, 64)
	//log.Println("查询的用户id：", userId)
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
	//var userMap map[string]interface{}
	//err = json.Unmarshal(resp.UserInfo, &userMap)
	//if err != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"msg":       resp.Msg,
		"user_info": json.RawMessage(resp.UserInfo),
	})
}

func HandleDeleteUserAccount(context *gin.Context) {
	var err error
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	confirmed, err := strconv.ParseBool(context.Query("confirmed"))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"deleted": false,
			"msg":     "Bad request",
		})
		return
	}
	//log.Println(userId, confirmed)
	//
	if !confirmed {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusConflict,
			"deleted": false,
			"msg":     "Not confirmed",
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.DeleteMyAccount(sysContext.Background(), &pb.DeleteMyAccountRequest{
		UserId: userId,
	})
	if err := failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"deleted": false,
			"msg":     err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"deleted": resp.Deleted,
		"msg":     resp.Msg,
	})

}

func HandleGetAllUsers(context *gin.Context) {
	var page, size int64
	var err error
	if err, page, size = getPage2Size(context); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	var searchEmail = context.Query("email")
	//log.Println(page, size)
	resp, err := grpcClient.UserServiceClient.GetAllUsers(sysContext.Background(), &pb.GetAllUsersRequest{
		Page:        page,
		Size:        size,
		SearchEmail: searchEmail,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var userMap []map[string]interface{}
	//err = json.Unmarshal(resp.Users, &userMap)
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
		"users":      json.RawMessage(resp.Users),
		"page_count": resp.PageCount,
	})

}

func HandleAddUserManuallyFromAdmin(context *gin.Context) {
	postData := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.AddUserManually(sysContext.Background(), &pb.AddUserManuallyRequest{
		Email:    postData.Email,
		Password: postData.Password,
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
	//log.Println(postData)
}

func HandleUpdateUserInfoByIdFromAdmin(context *gin.Context) {
	postData := &struct {
		Id         int64   `json:"id"`
		Email      string  `json:"email"`
		Balance    float32 `json:"balance"`
		InviteCode string  `json:"invite_code"`
		IsAdmin    bool    `json:"is_admin"`
		IsStaff    bool    `json:"is_staff"`
		Password   string  `json:"password"`
	}{}
	if err := context.ShouldBindJSON(postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	//log.Println(postData)
	//log.Println(postData)
	resp, err := grpcClient.UserServiceClient.UpdateUserInfoAdmin(sysContext.Background(), &pb.UpdateUserInfoAdminRequest{
		Id:         postData.Id,
		Email:      postData.Email,
		Password:   postData.Password,
		InviteCode: postData.InviteCode,
		IsAdmin:    postData.IsAdmin,
		IsStaff:    postData.IsStaff,
		Balance:    postData.Balance,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"success": false,
			"msg":     err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"success": resp.Success,
	})
}

func HandleBlock2UnblockUserByIdFromAdmin(context *gin.Context) {
	postData := &struct {
		UserId int64 `json:"user_id"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.Block2UnblockUserById(sysContext.Background(), &pb.Block2UnblockUserByIdRequest{
		UserId: postData.UserId,
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

func HandleGetUserInviteCodeByUserId(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.GetUserInviteCodeByUserId(sysContext.Background(), &pb.GetUserInviteCodeByUserIdRequest{
		UserId: userId,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":           resp.Code,
		"my_invite_code": resp.MyInviteCode,
		"msg":            resp.Msg,
	})
}

func HandleCreateUserInviteCodeByUserId(context *gin.Context) {
	postData := &struct {
		UserId int64 `json:"user_id"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.CreateUserInviteCodeByUserId(sysContext.Background(), &pb.CreateUserInviteCodeByUserIdRequest{
		UserId: postData.UserId,
	})
	if err = failOnRpcError(err, resp); err != nil {
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

func HandleGetUserInvitedUserListByUserId(context *gin.Context) {
	err, page, size := getPage2Size(context)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.GetUserInvitedUserListByUserId(sysContext.Background(), &pb.GetUserInvitedUserListByUserIdRequest{
		UserId: userId,
		Page:   page,
		Size:   size,
	})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	//var usersMap []map[string]interface{}
	//if err := json.Unmarshal(resp.UserList, &usersMap); err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  err.Error(),
	//	})
	//	return
	//}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
		//"user_list":  usersMap,
		"user_list":  json.RawMessage(resp.UserList),
		"page_count": resp.PageCount,
	})
}

func HandleUserUploadAvatar(context *gin.Context) {
	userIDStr := context.PostForm("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "invalid user_id, must be int64",
		})
		return
	}

	fileName := context.PostForm("file_name")

	// 获取上传的文件
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": "file is required",
		})
		return
	}

	// 打开文件并读取数据
	src, err := file.Open()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Failed to open file",
		})
		return
	}
	defer src.Close()

	// 读取文件内容到字节数组
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Failed to read file",
		})
		return
	}

	// 调用 gRPC 服务上传头像
	resp, err := grpcClient.UserServiceClient.UploadUserAvatar(context.Request.Context(), &pb.UploadUserAvatarRequest{
		UserId:   userID,
		FileName: fileName,
		FileData: fileBytes,
	})
	if err = failOnRpcError(err, resp); err != nil {
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

func HandleGetUserAvatar(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Query("user_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.GetUserAvatar(sysContext.Background(), &pb.GetUserAvatarRequest{
		UserId: userId,
	})
	if failOnRpcError(err, resp) != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":      resp.Code,
		"msg":       resp.Msg,
		"file_name": resp.FileName,
		"file_data": resp.FileData,
	})
}

func HandleAlterUsername(context *gin.Context) {
	postData := &struct {
		UserId int64  `json:"user_id"`
		Name   string `json:"name"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	resp, err := grpcClient.UserServiceClient.AlterUserSecondaryName(sysContext.Background(), &pb.AlterUserSecondaryNameRequest{
		UserId:  postData.UserId,
		NewName: postData.Name,
	})
	if err = failOnRpcError(err, resp); err != nil {
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

func HandleGetUserLayout(context *gin.Context) {
	resp, err := grpcClient.UserServiceClient.GetUsersLayout(sysContext.Background(), &pb.GetUsersLayoutRequest{})
	if err = failOnRpcError(err, resp); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":          resp.Code,
		"msg":           resp.Msg,
		"user_all":      resp.UserAll,
		"user_active":   resp.UserActive,
		"user_inactive": resp.UserInactive,
		"user_blocked":  resp.UserBlocked,
	})
}
