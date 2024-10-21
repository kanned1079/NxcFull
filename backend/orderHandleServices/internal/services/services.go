package services

import (
	pb "orderHandleServices/api/proto"
)

type OrderHandleServices struct {
	pb.UnimplementedOrderHandleServiceServer
}

func NewOrderHandleServices() *OrderHandleServices {
	return &OrderHandleServices{}
}
