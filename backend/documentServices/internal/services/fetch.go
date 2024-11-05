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
	//query := dao.Db.Model(&model.Document{}).Order("sort DESC").Where("language = ? AND show = ?", req.Language, true)
	//query := dao.Db.Model(&model.Document{}).Order("sort DESC").Where("language = ? AND show = ?", req.Language, true)
	query := dao.Db.Model(&model.Document{}).Order("sort DESC").Where("language = ? AND `show` = ?", req.Language, true)
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

//func (s *DocumentService) GetDocumentsAdmin(context context.Context, request *pb.GetDocumentsAdminRequest) (*pb.GetDocumentsAdminResponse, error) {
//	//request.Page
//	//request.Size
//	// 完成上述的功能
//	var documents []model.Document
//	if result := dao.Db.Model(&model.Document{}).Find(&documents); result.Error != nil {
//		return &pb.GetDocumentsAdminResponse{
//			Code: http.StatusInternalServerError,
//			Msg:  "Query failure" + result.Error.Error(),
//		}, nil
//	}
//	return &pb.GetDocumentsAdminResponse{
//		Code:      http.StatusOK,
//		PageCount: 0, // 由给出的page计算而来
//		Msg:       "Query ok",
//	}, nil
//}

func (s *DocumentService) GetDocumentsAdmin(context context.Context, request *pb.GetDocumentsAdminRequest) (*pb.GetDocumentsAdminResponse, error) {
	var documents []model.Document

	// 获取分页参数
	page := request.Page
	size := request.Size

	// 设置默认分页参数，以防请求中未提供
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	// 计算偏移量和每页限制
	offset := (page - 1) * size

	// 查询总记录数
	var total int64
	if result := dao.Db.Model(&model.Document{}).Count(&total); result.Error != nil {
		return &pb.GetDocumentsAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Count failure: " + result.Error.Error(),
		}, nil
	}

	// 查询带分页的数据
	if result := dao.Db.Model(&model.Document{}).Offset(int(offset)).Limit(int(size)).Order("sort DESC").Find(&documents); result.Error != nil {
		return &pb.GetDocumentsAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Query failure: " + result.Error.Error(),
		}, nil
	}

	// 计算总页数
	pageCount := (total + int64(size) - 1) / int64(size)

	if documentsJson, err := json.Marshal(documents); err != nil {
		return &pb.GetDocumentsAdminResponse{
			Code: http.StatusInternalServerError,
			Msg:  "Marshal failure: " + err.Error(),
		}, nil
	} else {
		return &pb.GetDocumentsAdminResponse{
			Code:      http.StatusOK,
			PageCount: pageCount,
			Msg:       "Query ok",
			Documents: documentsJson,
		}, nil
	}

}
