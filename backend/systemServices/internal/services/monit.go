package services

/**
rpc定义的返回值内容
message GetServerLiveStatusResponse {
  int32 code = 1;
  int64 status200 = 2;
  int64 status404 = 3;
  int64 status500 = 4;
  int64 login_req = 5;
  int64 register_req = 6;
  bytes api_log_list = 7;
  int64 page_size = 8;
}
*/

//func (s *SettingServices) GetServerLiveStatus(ctx context.Context, request *pb.GetServerLiveStatusRequest) (*pb.GetServerLiveStatusResponse, error)  {
//	dao.Db.Model(&model)
//
//
//	return &pb.GetServerLiveStatusResponse{}, nil
//}
