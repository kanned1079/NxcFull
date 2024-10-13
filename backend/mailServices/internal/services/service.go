package services

import (
	pb "mailServices/api/proto"
)

type MailServices struct {
	pb.UnimplementedGroupServiceServer
}

func NewMailServices() *MailServices {
	return &MailServices{}
}
