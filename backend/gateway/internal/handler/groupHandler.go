package handler

import (
	sysContext "context"
	pb "gateway/internal/grpc/api/group/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllGroups(context *gin.Context) {
	pageData := &struct {
		Page int64 `form:"page" json:"page"`
		Size int64 `form:"size" json:"size"`
	}{}
	if err := context.ShouldBindQuery(pageData); err != nil {
		log.Println(err)
	}
	log.Println(pageData)
	resp, err := grpcClient.GroupServiceClient.GetAllGroups(sysContext.Background(), &pb.GetAllGroupsRequest{
		Page: pageData.Page,
		Size: pageData.Size,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  resp.Msg,
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"group_list": resp.GroupList,
		"msg":        resp.Msg,
		"page_count": resp.PageCount,
	})

}

func HandleGetAllGroupsKv(context *gin.Context) {
	resp, err := grpcClient.GroupServiceClient.GetAllGroupsKv(sysContext.Background(), &pb.GetAllGroupsKvRequest{})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc方法错误" + err.Error(),
		})
		return
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc方法错误无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"group_list": resp.GroupList,
		"msg":        resp.Msg,
	})
}

func HandleAddNewGroup(context *gin.Context) {
	postData := &struct {
		GroupName string `json:"group_name"`
	}{}
	if err := context.ShouldBind(&postData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "绑定数据失败",
			"error": err.Error(),
		})
	}
	log.Println(postData)
	resp, err := grpcClient.GroupServiceClient.AddNewGroup(sysContext.Background(), &pb.AddNewGroupRequest{
		GroupName: postData.GroupName,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  resp.Msg,
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})

	//var newPrivilegeGroup Group
	//newPrivilegeGroup.Name = postData.GroupName
	//if err := dao.Db.Model(&Group{}).Create(&newPrivilegeGroup).Error; err != nil {
	//	log.Println(err)
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"code":  http.StatusInternalServerError,
	//		"msg":   "插入数据失败",
	//		"error": err.Error(),
	//	})
	//}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
}

func HandleUpdateGroup(context *gin.Context) {
	// put方法 请求体数据保存到
	putData := &struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}
	if err := context.ShouldBind(&putData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "解析失败",
			"error": err.Error(),
		})
	}
	resp, err := grpcClient.GroupServiceClient.UpdateGroup(sysContext.Background(), &pb.UpdateGroupRequest{
		Id:   putData.Id,
		Name: putData.Name,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  resp.Msg,
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func HandleDeleteGroup(context *gin.Context) {
	delData := &struct {
		Id int64 `json:"id" binding:"required"`
	}{}
	if err := context.ShouldBindJSON(&delData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}
	resp, err := grpcClient.GroupServiceClient.DeleteGroup(sysContext.Background(), &pb.DeleteGroupRequest{
		Id: delData.Id,
	})
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  resp.Msg,
		})
	}
	if resp == nil {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "调用rpc服务器失败无返回值",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
