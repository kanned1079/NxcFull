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
