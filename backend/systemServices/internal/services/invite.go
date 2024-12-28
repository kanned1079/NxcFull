package services

import (
	"context"
	"encoding/json"
	"net/http"
	pb "systemServices/api/proto"
)

func (s *SettingServices) GetInviteUserMsg(ctx context.Context, request *pb.GetInviteUserMsgRequest) (*pb.GetInviteUserMsgResponse, error) {
	var inviteMsgJson json.RawMessage
	inviteMsgJson, err = readSettingFromDB("invite", "invite_info")
	if err != nil {
		return &pb.GetInviteUserMsgResponse{
			Code:      http.StatusNotFound,
			InviteMsg: "",
			Msg:       "no invite info found",
		}, nil
	}
	var inviteMsgStr string
	err = json.Unmarshal(inviteMsgJson, &inviteMsgStr)
	if err != nil {
		return &pb.GetInviteUserMsgResponse{
			Code:      http.StatusInternalServerError,
			InviteMsg: "",
			Msg:       "error unmarshalling invite info",
		}, nil
	}
	return &pb.GetInviteUserMsgResponse{
		Code:      http.StatusOK,
		InviteMsg: inviteMsgStr,
		Msg:       "ok",
	}, nil
}
