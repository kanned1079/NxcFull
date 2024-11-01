package services

import (
	pb "documentServices/api/proto"
)

//type DocType {
//
//}

type DocumentService struct {
	pb.UnimplementedDocumentServiceServer
}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}
