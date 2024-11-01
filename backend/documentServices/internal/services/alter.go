package services

import (
	"context"
	pb "documentServices/api/proto"
	"documentServices/internal/dao"
	"documentServices/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CategoryDocument struct {
	Category  string           `json:"category"`
	Documents []model.Document `json:"documents"`
}

// v1 无redis缓存
//func (s *DocumentService) GetDocuments(ctx context.Context, req *pb.GetDocumentsRequest) (*pb.GetDocumentsResponse, error) {
//	log.Println("lang: ", req.Language)
//	log.Println("find: ", req.Find) // 打印 Find 的值
//
//	var documentsList []model.Document
//	query := dao.Db.Model(&model.Document{}).Where("language = ?", req.Language)
//
//	// 如果 Find 字段不为空，添加模糊查询条件
//	if req.Find != "" {
//		query = query.Where("title LIKE ?", "%"+req.Find+"%")
//	}
//
//	// 执行查询
//	if result := query.Find(&documentsList); result.Error != nil {
//		log.Println("获取所有文档失败:", result.Error)
//		return &pb.GetDocumentsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "获取所有文档失败",
//		}, nil
//	}
//
//	// 创建一个 map 来存储类别和相应的文档
//	categories := make(map[string][]model.Document)
//	for _, doc := range documentsList {
//		// 根据类别进行分类
//		categories[doc.Category] = append(categories[doc.Category], doc)
//	}
//
//	// 准备分类文档的 JSON 序列化
//	var categoryDocs []CategoryDocument
//	for category, documents := range categories {
//		categoryDocs = append(categoryDocs, CategoryDocument{
//			Category:  category,
//			Documents: documents,
//		})
//	}
//
//	// 将 categoryDocs 转为 JSON 字符串
//	docJson, err := json.Marshal(categoryDocs)
//	if err != nil {
//		log.Println("JSON 序列化失败:", err)
//		return &pb.GetDocumentsResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "JSON 序列化失败",
//		}, nil
//	}
//
//	// 返回 JSON 序列化的文档内容
//	return &pb.GetDocumentsResponse{
//		Code:      http.StatusOK,
//		Msg:       "获取成功",
//		Documents: docJson,
//	}, nil
//}

// GetDocuments 获取所有文档 包含redis缓存
func (s *DocumentService) GetDocuments(ctx context.Context, req *pb.GetDocumentsRequest) (*pb.GetDocumentsResponse, error) {
	log.Println("lang: ", req.Language)
	log.Println("find: ", req.Find)

	// 生成缓存键，仅使用语言作为唯一标识
	cacheKey := fmt.Sprintf("documents:%s", req.Language)

	// 如果 Find 字段为空，检查 Redis 缓存
	if req.Find == "" {
		cachedData, err := dao.Rdb.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			log.Println("从缓存中获取数据")

			// 直接返回缓存的 JSON 数据
			return &pb.GetDocumentsResponse{
				Code:      http.StatusOK,
				Msg:       "获取成功",
				Documents: []byte(cachedData), // JSON 字符串
			}, nil
		}
	}

	// 缓存未命中或有 Find 参数，执行数据库查询
	var documentsList []model.Document
	query := dao.Db.Model(&model.Document{}).Where("language = ?", req.Language)

	// 如果 Find 字段不为空，添加模糊查询条件
	if req.Find != "" {
		query = query.Where("title LIKE ?", "%"+req.Find+"%")
	}

	// 执行查询
	if result := query.Find(&documentsList); result.Error != nil {
		log.Println("获取所有文档失败:", result.Error)
		return &pb.GetDocumentsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取所有文档失败",
		}, nil
	}

	// 创建一个 map 来存储类别和相应的文档
	categories := make(map[string][]model.Document)
	for _, doc := range documentsList {
		// 根据类别进行分类
		categories[doc.Category] = append(categories[doc.Category], doc)
	}

	// 准备分类文档的 JSON 序列化
	var categoryDocs []CategoryDocument
	for category, documents := range categories {
		categoryDocs = append(categoryDocs, CategoryDocument{
			Category:  category,
			Documents: documents,
		})
	}

	// 将 categoryDocs 转为 JSON 字符串
	docJson, err := json.Marshal(categoryDocs)
	if err != nil {
		log.Println("JSON 序列化失败:", err)
		return &pb.GetDocumentsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "JSON 序列化失败",
		}, nil
	}

	// 如果 Find 为空，将结果存入 Redis 缓存
	if req.Find == "" {
		if err := dao.Rdb.Set(ctx, cacheKey, docJson, time.Minute*10).Err(); err != nil {
			log.Println("设置缓存失败:", err)
		}
	}

	// 返回查询结果，Documents 字段为 JSON 字符串
	return &pb.GetDocumentsResponse{
		Code:      http.StatusOK,
		Msg:       "获取成功",
		Documents: docJson, // JSON 字符串
	}, nil
}

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

	// 保存更新
	if err := dao.Db.Save(&doc).Error; err != nil {
		return &pb.UpdateDocumentResponse{
			Code: http.StatusInternalServerError,
			Msg:  "保存修改失败",
		}, nil
	}

	return &pb.UpdateDocumentResponse{
		Code: http.StatusOK,
		Msg:  "更新成功",
	}, nil
}
