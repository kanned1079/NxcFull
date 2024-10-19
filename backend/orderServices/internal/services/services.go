package services

import (
	pb "orderServices/api/proto"
)

type OrderServices struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderServices() *OrderServices {
	return &OrderServices{}
}
