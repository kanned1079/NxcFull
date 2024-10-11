package services

import (
	"context"
	pb "groupServices/api/proto"
	"groupServices/internal/dao"
	"groupServices/internal/model"
	"log"
	"net/http"
)

type GroupServices struct {
	pb.UnimplementedGroupServiceServer
}

func NewGroupServices() *GroupServices {
	return &GroupServices{}
}

// v1
//func (s *GroupServices) GetAllGroups(ctx context.Context, request *pb.GetAllGroupsRequest) (*pb.GetAllGroupsResponse, error) {
//	// pass
//	// 这里有request.page和request.page还没有实现
//	// pass
//	log.Println(request.Page, request.Size)
//	var privilegeGroups []model.Group
//	type ResponseData struct {
//		Id        int64  `json:"id"`
//		Name      string `json:"name"`
//		UserCount int64  `json:"user_count"`
//		PlanCount int64  `json:"plan_count"`
//	}
//	var responseDataList []pb.ResponseGroup
//
//	// Query all groups from the database
//	if err := dao.Db.Model(&model.Group{}).Find(&privilegeGroups).Error; err != nil {
//		log.Println(err)
//		return &pb.GetAllGroupsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "查询失败",
//		}, nil
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
//		if err := dao.Db.Model(&model.Plan{}).Where("group_id = ?", group.Id).Count(&planCount).Error; err != nil {
//			log.Println("Failed to get plan count:", err)
//			return &pb.GetAllGroupsResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "查询计划数量失败",
//			}, nil
//		}
//		temp.PlanCount = planCount
//
//		// Query UserCount from ActiveOrders database
//		var userCount int64
//
//		if err := dao.Db.Model(&model.ActiveOrders{}).Where("group_id = ? && is_active = ?", group.Id, true).Count(&userCount).Error; err != nil {
//			log.Println("Failed to get user count:", err)
//			return &pb.GetAllGroupsResponse{
//				Code: http.StatusInternalServerError,
//				Msg:  "查询用户数量失败",
//			}, nil
//		}
//		temp.UserCount = userCount
//
//		// Append temp to responseDataList
//		responseDataList = append(responseDataList, pb.ResponseGroup{
//			Id:        temp.Id,
//			Name:      temp.Name,
//			PlanCount: temp.PlanCount,
//			UserCount: temp.UserCount,
//		})
//	}
//
//	log.Println(responseDataList)
//	return &pb.GetAllGroupsResponse{
//		Code:      http.StatusOK,
//		GroupList: ConvertToPointerResponseGroups(responseDataList),
//		Msg:       "查询成功",
//		PageCount: 10, // 实现这里的页面数 根据前面提供的size来计算有多少页面
//	}, nil
//}

func (s *GroupServices) GetAllGroups(ctx context.Context, request *pb.GetAllGroupsRequest) (*pb.GetAllGroupsResponse, error) {
	log.Println(request.Page, request.Size)
	var privilegeGroups []model.Group
	type ResponseData struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		UserCount int64  `json:"user_count"`
		PlanCount int64  `json:"plan_count"`
	}
	var responseDataList []pb.ResponseGroup

	// 分页参数
	page := request.Page
	size := request.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10 // 默认每页显示10条记录
	}
	offset := (page - 1) * size

	// 获取总记录数
	var totalCount int64
	if err := dao.Db.Model(&model.Group{}).Count(&totalCount).Error; err != nil {
		log.Println(err)
		return &pb.GetAllGroupsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败",
		}, nil
	}

	// 分页查询
	if err := dao.Db.Model(&model.Group{}).Limit(int(size)).Offset(int(offset)).Find(&privilegeGroups).Error; err != nil {
		log.Println(err)
		return &pb.GetAllGroupsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败",
		}, nil
	}

	// 填充 responseDataList
	for _, group := range privilegeGroups {
		temp := ResponseData{
			Id:   group.Id,
			Name: group.Name,
		}

		// 查询 PlanCount
		var planCount int64
		if err := dao.Db.Model(&model.Plan{}).Where("group_id = ?", group.Id).Count(&planCount).Error; err != nil {
			log.Println("Failed to get plan count:", err)
			return &pb.GetAllGroupsResponse{
				Code: http.StatusInternalServerError,
				Msg:  "查询计划数量失败",
			}, nil
		}
		temp.PlanCount = planCount

		// 查询 UserCount
		var userCount int64
		if err := dao.Db.Model(&model.ActiveOrders{}).Where("group_id = ? AND is_active = ?", group.Id, true).Count(&userCount).Error; err != nil {
			log.Println("Failed to get user count:", err)
			return &pb.GetAllGroupsResponse{
				Code: http.StatusInternalServerError,
				Msg:  "查询用户数量失败",
			}, nil
		}
		temp.UserCount = userCount

		// 将数据追加到 responseDataList
		responseDataList = append(responseDataList, pb.ResponseGroup{
			Id:        temp.Id,
			Name:      temp.Name,
			PlanCount: temp.PlanCount,
			UserCount: temp.UserCount,
		})
	}

	// 计算总页数
	pageCount := int64((totalCount + int64(size) - 1) / int64(size)) // 向上取整

	log.Println(responseDataList)
	return &pb.GetAllGroupsResponse{
		Code:      http.StatusOK,
		GroupList: ConvertToPointerResponseGroups(responseDataList),
		Msg:       "查询成功",
		PageCount: pageCount,
	}, nil
}

func (s *GroupServices) GetAllGroupsKv(ctx context.Context, request *pb.GetAllGroupsKvRequest) (*pb.GetAllGroupsKvResponse, error) {
	var groups []model.Group
	var responseDataList []*pb.ResponseGroupKv

	// 查询所有 Group 的 Id 和 Name
	if err := dao.Db.Model(&model.Group{}).Select("id, name").Find(&groups).Error; err != nil {
		log.Println(err)
		return &pb.GetAllGroupsKvResponse{
			Code: http.StatusInternalServerError,
			Msg:  "查询失败" + err.Error(),
		}, nil
	}

	// 填充 responseDataList
	for _, group := range groups {
		responseDataList = append(responseDataList, &pb.ResponseGroupKv{
			Id:   group.Id,
			Name: group.Name,
		})
	}

	log.Println(responseDataList)
	return &pb.GetAllGroupsKvResponse{
		Code:      http.StatusOK,
		GroupList: responseDataList,
		Msg:       "查询成功",
	}, nil
}

func (s *GroupServices) AddNewGroup(ctx context.Context, request *pb.AddNewGroupRequest) (*pb.AddNewGroupResponse, error) {
	var newPrivilegeGroup model.Group
	newPrivilegeGroup.Name = request.GroupName
	if err := dao.Db.Model(&model.Group{}).Create(&newPrivilegeGroup).Error; err != nil {
		log.Println(err)
		return &pb.AddNewGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据失败",
		}, nil
	}
	return &pb.AddNewGroupResponse{
		Code: http.StatusOK,
		Msg:  "插入数据成功",
	}, nil
}

func (s *GroupServices) UpdateGroup(ctx context.Context, request *pb.UpdateGroupRequest) (*pb.UpdateGroupResponse, error) {
	if result := dao.Db.Model(&model.Group{}).Where("id = ?", request.Id).Updates(&model.Group{Name: request.Name}); result.Error != nil {
		log.Println(result.Error)
		return &pb.UpdateGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存到数据库失败",
		}, nil
	}
	return &pb.UpdateGroupResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil
}

func (s *GroupServices) DeleteGroup(ctx context.Context, request *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	// 从数据库中删除对应 ID 的 Group
	if err := dao.Db.Delete(&model.Group{}, request.Id).Error; err != nil {
		log.Println(err)
		return &pb.DeleteGroupResponse{
			Code: http.StatusInternalServerError,
			Msg:  "删除失败",
		}, nil
	}

	return &pb.DeleteGroupResponse{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
}
