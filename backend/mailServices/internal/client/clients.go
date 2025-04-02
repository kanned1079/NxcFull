package client

import (
	"mailServices/internal/client/api/settings"
	settingsPb "mailServices/internal/client/api/settings/proto"
	"mailServices/internal/client/api/user"
	userPb "mailServices/internal/client/api/user/proto"
)

type Clients struct {
	UserServiceClient    userPb.UserServiceClient
	SystemServicesClient settingsPb.SettingsServiceClient
}

func NewClients() *Clients {
	// 初始化gRPC客户端
	return &Clients{
		UserServiceClient:    user.NewUserClient(),
		SystemServicesClient: settings.NewSettingClient(),
	}

}
