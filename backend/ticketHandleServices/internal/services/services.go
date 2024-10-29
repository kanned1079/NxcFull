package services

import (
	pb "ticketHandleServices/api/proto"
)

type TicketHandleServices struct {
	pb.UnimplementedTicketHandleServiceServer
}

func NewTicketHandleServices() *TicketHandleServices {
	return &TicketHandleServices{}
}
