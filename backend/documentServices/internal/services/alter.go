package services

import (
	"context"
	pb "documentServices/api/proto"
	"documentServices/internal/dao"
	"documentServices/internal/model"
	"documentServices/internal/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (s *DocumentService) AddDocument(ctx context.Context, req *pb.AddDocumentRequest) (*pb.AddDocumentResponse, error) {
	newDoc := model.Document{
		Title:    req.Title,
		Category: req.Category,
		Language: req.Language,
		Body:     req.Body,
	}
	if err := dao.Db.Model(&model.Document{}).Create(&newDoc).Error; err != nil {
		return &pb.AddDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "插入数据库失败",
		}, nil
	}
	if err := utils.FreshDocumentsRedisCache(); err != nil {
		return &pb.AddDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Clear documents cache failure, so results may not correct." + err.Error(),
		}, nil
	}
	return &pb.AddDocumentResponse{
		Code: http.StatusOK,
		Msg:  "添加成功",
	}, nil
}

func (s *DocumentService) UpdateDocument(ctx context.Context, req *pb.UpdateDocumentRequest) (*pb.UpdateDocumentResponse, error) {
	// 查找需要更新的文档
	var doc model.Document
	if err := dao.Db.Model(&model.Document{}).Where("id = ?", req.Id).First(&doc).Error; err != nil {
		return &pb.UpdateDocumentResponse{
			Code: http.StatusNotFound,
			Msg:  "文档未找到",
		}, nil
	}

	// 更新文档字段
	doc.Title = req.Title
	doc.Category = req.Category
	doc.Language = req.Language
	doc.Body = req.Body
	doc.Sort = req.Sort

	// 保存更新
	if err := dao.Db.Save(&doc).Error; err != nil {
		return &pb.UpdateDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存修改失败",
		}, nil
	}

	if err := utils.FreshDocumentsRedisCache(); err != nil {
		return &pb.UpdateDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Clear documents cache failure, so results may not correct." + err.Error(),
		}, nil
	}

	return &pb.UpdateDocumentResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil
}

func (s *DocumentService) UpdateDocumentShow(context context.Context, request *pb.UpdateDocumentShowRequest) (*pb.UpdateDocumentShowResponse, error) {
	//if result := dao.Db.Model(&model.Document{}).Where("id = ?", request.Id).Update("show = ?", request.Show); result.Error != nil {
	// 直接取反處理
	if result := dao.Db.Model(&model.Document{}).Where("id = ?", request.Id).Update("`show`", gorm.Expr("NOT `show`")); result.Error != nil {
		return &pb.UpdateDocumentShowResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Operate failure" + result.Error.Error(),
		}, nil
	}
	if err := utils.FreshDocumentsRedisCache(); err != nil {
		return &pb.UpdateDocumentShowResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Clear documents cache failure, so results may not correct." + err.Error(),
		}, nil
	}
	return &pb.UpdateDocumentShowResponse{
		Code: http.StatusOK,
		Msg:  "Operate ok",
	}, nil
}

func (s *DocumentService) DeleteDocument(context context.Context, request *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	if result := dao.Db.Model(&model.Document{}).Where("id = ?", request.Id).Delete(&model.Document{}); result.Error != nil {
		return &pb.DeleteDocumentResponse{
			Code:    http.StatusInternalServerError,
			Deleted: false,
			Msg:     "Operate failure" + result.Error.Error(),
		}, nil
	}
	log.Println("刷新doc緩存")
	//if err := utils.FreshDocumentsRedisCache(); err != nil {
	//	return &pb.DeleteDocumentResponse{
	//		Code: http.StatusInternalServerError,
	//		Msg:  "Clear documents cache failure, so results may not correct." + err.Error(),
	//	}, nil
	//}
	if err := utils.FreshDocumentsRedisCache(); err != nil {
		return &pb.DeleteDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Clear documents cache failure, so results may not correct." + err.Error(),
		}, nil
	}
	return &pb.DeleteDocumentResponse{
		Code:    http.StatusOK,
		Deleted: true,
		Msg:     "Operate ok",
	}, nil
}
