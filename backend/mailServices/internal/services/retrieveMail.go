package services

import (
	"context"
	pb "mailServices/api/proto"
)

func (s *MailServices) SendRetrievePwdVerifyCode2Email(context context.Context, request *pb.SendRetrievePwdVerifyCode2EmailRequest) (*pb.SendRetrievePwdVerifyCode2EmailResponse, error) {
	return &pb.SendRetrievePwdVerifyCode2EmailResponse{}, nil
}
