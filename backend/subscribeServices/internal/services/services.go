package services

import (
	pb "subscribeServices/api/proto"
)

type SubscribeServices struct {
	pb.UnimplementedSubscriptionServiceServer
}

func NewOrderServices() *SubscribeServices {
	return &SubscribeServices{}
}
