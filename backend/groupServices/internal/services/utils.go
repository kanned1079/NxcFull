package services

import (
	pb "NxcFull/backend/groupServices/api/proto"
)

// 转换函数: 将 []model.PublicNotices 转换为 []*pb.PublicNotice
//func convertToProtoNotices(notices []model.PublicNotices) []*pb.PublicNotice {
//	protoNotices := make([]*pb.PublicNotice, len(notices))
//	for i, notice := range notices {
//		protoNotices[i] = &pb.PublicNotice{
//			Id:        notice.Id,
//			Title:     notice.Title,
//			Content:   notice.Content,
//			Show:      notice.Show,
//			ImgUrl:    notice.ImgUrl,
//			Tags:      notice.Tags,
//			CreatedAt: notice.CreatedAt.Format(time.RFC3339), // 时间转为字符串
//			UpdatedAt: notice.UpdatedAt.Format(time.RFC3339),
//			DeletedAt: func() string {
//				if !notice.DeletedAt.Valid {
//					return ""
//				}
//				return notice.DeletedAt.Time.Format(time.RFC3339)
//			}(),
//		}
//	}
//	return protoNotices
//}

// ConvertToPointerResponseGroups 将 []pb.ResponseGroup 转换为 []*pb.ResponseGroup 的函数
func ConvertToPointerResponseGroups(responseDataList []pb.ResponseGroup) []*pb.ResponseGroup {
	var groupList []*pb.ResponseGroup
	for _, respData := range responseDataList {
		group := &pb.ResponseGroup{
			Id:        respData.Id,
			Name:      respData.Name,
			UserCount: respData.UserCount,
			PlanCount: respData.PlanCount,
		}
		groupList = append(groupList, group)
	}
	return groupList
}
