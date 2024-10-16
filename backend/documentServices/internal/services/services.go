package services

import (
	"context"
	pb "documentServices/api/proto"
	"documentServices/internal/dao"
	"documentServices/internal/model"
	"log"
	"net/http"
)

type DocumentService struct {
	pb.UnimplementedDocumentServiceServer
}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

func (s *DocumentService) GetDocuments(ctx context.Context, req *pb.GetDocumentsRequest) (*pb.GetDocumentsResponse, error) {
	log.Println("lang: ", req.Language)
	log.Println("find: ", req.Find) // 打印 Find 的值

	var documentsList []model.Document
	query := dao.Db.Model(&model.Document{}).Where("language = ?", req.Language)

	// 如果 Find 字段不为空，添加模糊查询条件
	if req.Find != "" {
		query = query.Where("title LIKE ?", "%"+req.Find+"%")
	}

	// 执行查询
	if result := query.Find(&documentsList); result.Error != nil {
		log.Println("获取所有文档失败")
		return &pb.GetDocumentsResponse{
			Code: http.StatusInternalServerError,
			Msg:  "获取所有文档失败",
		}, nil
	}

	// 创建一个 map 来存储类别和相应的文档
	categories := make(map[string][]*pb.Document)
	for _, doc := range documentsList {
		// 根据类别进行分类，确保只有匹配的文档被返回
		categories[doc.Category] = append(categories[doc.Category], &pb.Document{
			Id:        doc.Id,
			Language:  doc.Language,
			Category:  doc.Category,
			Title:     doc.Title,
			Body:      doc.Body,
			Sort:      int64(doc.Sort),
			Show:      doc.Show,
			CreatedAt: doc.CreatedAt.Format("2006-01-02"), // 格式化日期
			UpdatedAt: doc.UpdatedAt.Format("2006-01-02"),
			DeletedAt: doc.DeletedAt.Time.Format("2006-01-02"), // 处理可空时间
		})
	}

	// 转换为 gRPC 响应需要的格式，只返回包含数据的类别
	var grpcDocuments []*pb.CategoryDocuments
	for category, docs := range categories {
		grpcDocuments = append(grpcDocuments, &pb.CategoryDocuments{
			Category:  category,
			Documents: docs,
		})
	}

	return &pb.GetDocumentsResponse{
		Code:      http.StatusOK,
		Msg:       "获取成功",
		Documents: grpcDocuments,
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
