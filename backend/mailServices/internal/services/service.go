package services

import (
	pb "mailServices/api/proto"
)

type MailServices struct {
	pb.UnimplementedMailServiceServer
}

func NewMailServices() *MailServices {
	return &MailServices{}
}
