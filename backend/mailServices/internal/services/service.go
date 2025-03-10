package services

import (
	pb "mailServices/api/proto"
)

type MailServices struct {
	pb.UnimplementedMailServiceServer
	FilePrefix string
}

func NewMailServices() *MailServices {
	return &MailServices{
		FilePrefix: "/app",
	}
}
