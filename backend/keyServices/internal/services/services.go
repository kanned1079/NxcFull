package services

import (
	pb "keyServices/api/proto"
)

type KeyServices struct {
	pb.UnimplementedKeyServiceServer
}

func NewKeyServices() *KeyServices {
	return &KeyServices{}
}
