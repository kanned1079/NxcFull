package services

import (
	pb "NxcFull/backend/groupServices/api/proto"
	"NxcFull/backend/groupServices/internal/dao"
	"NxcFull/backend/groupServices/internal/model"
	"context"
	"log"
	"net/http"
)

type GroupServices struct {
	pb.UnimplementedGroupServiceServer
}

func NewGroupServices() *GroupServices {
	return &GroupServices{}
}

func (s *GroupServices) GetAllGroups(ctx context.Context, request *pb.GetAllGroupsRequest) (*pb.GetAllGroupsResponse, error) {
	// pass
	// 这里有request.page和request.page还没有实现
	// pass
	var privilegeGroups []model.Group
	type ResponseData struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		UserCount int64  `json:"user_count"`
		PlanCount int64  `json:"plan_count"`
	}
	var responseDataList []pb.ResponseGroup

	// Query all groups from the database
	if err := dao.Db.Model(&model.Group{}).Find(&privilegeGroups).Error; err != nil {
		log.Println(err)
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "查询失败",
		//	"error": err.Error(),
		//})
		//return
		return &pb.GetAllGroupsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败",
		}, nil
	}

	// Populate responseDataList with group information
	for _, group := range privilegeGroups {
		temp := ResponseData{
			Id:   group.Id,
			Name: group.Name,
		}

		// Query PlanCount from Plan database
		var planCount int64
		if err := dao.Db.Model(&model.Plan{}).Where("group_id = ?", group.Id).Count(&planCount).Error; err != nil {
			log.Println("Failed to get plan count:", err)
			//context.JSON(http.StatusInternalServerError, gin.H{
			//	"code":  http.StatusInternalServerError,
			//	"msg":   "查询计划数量失败",
			//	"error": err.Error(),
			//})
			//return
			return &pb.GetAllGroupsResponse{
				Code: http.StatusInternalServerError,
				Msg:  "查询计划数量失败",
			}, nil
		}
		temp.PlanCount = planCount

		// Query UserCount from ActiveOrders database
		var userCount int64

		if err := dao.Db.Model(&model.ActiveOrders{}).Where("group_id = ? && is_active = ?", group.Id, true).Count(&userCount).Error; err != nil {
			log.Println("Failed to get user count:", err)
			//context.JSON(http.StatusInternalServerError, gin.H{
			//	"code":  http.StatusInternalServerError,
			//	"msg":   "查询用户数量失败",
			//	"error": err.Error(),
			//})
			//return
			return &pb.GetAllGroupsResponse{
				Code: http.StatusInternalServerError,
				Msg:  "查询用户数量失败",
			}, nil
		}
		temp.UserCount = userCount

		// Append temp to responseDataList
		responseDataList = append(responseDataList, pb.ResponseGroup{
			Id:        temp.Id,
			Name:      temp.Name,
			PlanCount: temp.PlanCount,
			UserCount: temp.UserCount,
		})
	}

	log.Println(responseDataList)
	//context.JSON(http.StatusOK, gin.H{
	//	"code":       http.StatusOK,
	//	"group_list": responseDataList,
	//})
	return &pb.GetAllGroupsResponse{
		Code:      http.StatusOK,
		GroupList: ConvertToPointerResponseGroups(responseDataList),
		Msg:       "查询成功",
	}, nil
}

func (s *GroupServices) AddNewGroup(ctx context.Context, request *pb.AddNewGroupRequest) (*pb.AddNewGroupResponse, error) {
	var newPrivilegeGroup model.Group
	newPrivilegeGroup.Name = request.GroupName
	if err := dao.Db.Model(&model.Group{}).Create(&newPrivilegeGroup).Error; err != nil {
		log.Println(err)
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "插入数据失败",
		//	"error": err.Error(),
		//})
		return &pb.AddNewGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据失败",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "插入数据成功",
	//})
	return &pb.AddNewGroupResponse{
		Code: http.StatusOK,
		Msg:  "插入数据成功",
	}, nil
}

func (s *GroupServices) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest) (*pb.UpdateGroupResponse, error) {
	if result := dao.Db.Model(&model.Group{}).Where("id = ?", request.Id).Updates(&model.Group{Name: request.Name}); result.Error != nil {
		log.Println(result.Error)
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"meg":   "保存到数据库失败",
		//	"error": result.Error.Error(),
		//})
		return &pb.UpdateGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存到数据库失败",
		}, nil
	}
	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "更新成功",
	//})
	return &pb.UpdateGroupResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil
}

func (s *GroupServices) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	// 从数据库中删除对应 ID 的 Group
	if err := dao.Db.Delete(&model.Group{}, request.Id).Error; err != nil {
		log.Println(err)
		//context.JSON(http.StatusInternalServerError, gin.H{
		//	"code":  http.StatusInternalServerError,
		//	"msg":   "删除失败",
		//	"error": err.Error(),
		//})
		//return
		return &pb.DeleteGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除失败",
		}, nil
	}

	//context.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "删除成功",
	//})
	return &pb.DeleteGroupResponse{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
}

//func HandleGetAllGroups(context *gin.Context) {
//	pageData := &struct {
//		Page int64 `form:"page" json:"page"`
//		Size int64 `form:"size" json:"size"`
//	}{}
//	if err := context.ShouldBindQuery(pageData); err != nil {
//		log.Println(err)
//	}
//	log.Println(pageData)
//	// pass
//	var privilegeGroups []Group
//	type ResponseData struct {
//		Id        int64  `json:"id"`
//		Name      string `json:"name"`
//		UserCount int64  `json:"user_count"`
//		PlanCount int64  `json:"plan_count"`
//	}
//	var responseDataList []ResponseData
//
//	// Query all groups from the database
//	if err := dao.Db.Model(&Group{}).Find(&privilegeGroups).Error; err != nil {
//		log.Println(err)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "查询失败",
//			"error": err.Error(),
//		})
//		return
//	}
//
//	// Populate responseDataList with group information
//	for _, group := range privilegeGroups {
//		temp := ResponseData{
//			Id:   group.Id,
//			Name: group.Name,
//		}
//
//		// Query PlanCount from Plan database
//		var planCount int64
//		if err := dao.Db.Model(&subscribePlan.Plan{}).Where("group_id = ?", group.Id).Count(&planCount).Error; err != nil {
//			log.Println("Failed to get plan count:", err)
//			context.JSON(http.StatusInternalServerError, gin.H{
//				"code":  http.StatusInternalServerError,
//				"msg":   "查询计划数量失败",
//				"error": err.Error(),
//			})
//			return
//		}
//		temp.PlanCount = planCount
//
//		// Query UserCount from ActiveOrders database
//		var userCount int64
//
//		if err := dao.Db.Model(&orders.ActiveOrders{}).Where("group_id = ? && is_active = ?", group.Id, true).Count(&userCount).Error; err != nil {
//			log.Println("Failed to get user count:", err)
//			context.JSON(http.StatusInternalServerError, gin.H{
//				"code":  http.StatusInternalServerError,
//				"msg":   "查询用户数量失败",
//				"error": err.Error(),
//			})
//			return
//		}
//		temp.UserCount = userCount
//
//		// Append temp to responseDataList
//		responseDataList = append(responseDataList, temp)
//	}
//
//	log.Println(responseDataList)
//	context.JSON(http.StatusOK, gin.H{
//		"code":       http.StatusOK,
//		"group_list": responseDataList,
//	})
//}

//func HandleAddNewGroup(context *gin.Context) {
//	postData := &struct {
//		GroupName string `json:"group_name"`
//	}{}
//	if err := context.ShouldBind(&postData); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "绑定数据失败",
//			"error": err.Error(),
//		})
//	}
//	log.Println(postData)
//	var newPrivilegeGroup Group
//	newPrivilegeGroup.Name = postData.GroupName
//	if err := dao.Db.Model(&Group{}).Create(&newPrivilegeGroup).Error; err != nil {
//		log.Println(err)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "插入数据失败",
//			"error": err.Error(),
//		})
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "插入数据成功",
//	})
//}

//func HandleUpdateGroup(context *gin.Context) {
//	// put方法 请求体数据保存到
//	putData := &struct {
//		Id   int64  `json:"id"`
//		Name string `json:"name"`
//	}{}
//	if err := context.ShouldBind(&putData); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "解析失败",
//			"error": err.Error(),
//		})
//	}
//	//groupNewInfo := Group{
//	//	//Id:   putData.Id,
//	//	Name: putData.Name,
//	//}
//	if result := dao.Db.Model(&Group{}).Where("id = ?", putData.Id).Updates(&Group{Name: putData.Name}); result.Error != nil {
//		log.Println(result.Error)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"meg":   "保存到数据库失败",
//			"error": result.Error.Error(),
//		})
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "更新成功",
//	})
//}
//
//func HandleDeleteGroup(context *gin.Context) {
//	delData := &struct {
//		Id int64 `json:"id" binding:"required"`
//	}{}
//	if err := context.ShouldBindJSON(&delData); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code": http.StatusInternalServerError,
//		})
//	}
//
//	// 从数据库中删除对应 ID 的 Group
//	if err := dao.Db.Delete(&Group{}, delData.Id).Error; err != nil {
//		log.Println(err)
//		context.JSON(http.StatusInternalServerError, gin.H{
//			"code":  http.StatusInternalServerError,
//			"msg":   "删除失败",
//			"error": err.Error(),
//		})
//		return
//	}
//
//	context.JSON(http.StatusOK, gin.H{
//		"code": http.StatusOK,
//		"msg":  "删除成功",
//	})
//}
