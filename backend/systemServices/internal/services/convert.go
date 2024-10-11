package services

import (
	"encoding/json"
	"fmt"
	"log"
	pb "systemServices/api/proto"
)

// 将 gRPC 的 *Value 转换为 json.RawMessage
func valueToJSONRawMessage(val *pb.Value) (json.RawMessage, error) {
	var result json.RawMessage

	// 根据 Value 的类型进行序列化
	switch v := val.Kind.(type) {
	case *pb.Value_StringValue:
		jsonBytes, err := json.Marshal(v.StringValue)
		if err != nil {
			return nil, err
		}
		result = json.RawMessage(jsonBytes)
	case *pb.Value_IntValue:
		jsonBytes, err := json.Marshal(v.IntValue)
		if err != nil {
			return nil, err
		}
		result = json.RawMessage(jsonBytes)
	case *pb.Value_BoolValue:
		jsonBytes, err := json.Marshal(v.BoolValue)
		if err != nil {
			return nil, err
		}
		result = json.RawMessage(jsonBytes)
	case *pb.Value_DoubleValue:
		jsonBytes, err := json.Marshal(v.DoubleValue)
		if err != nil {
			return nil, err
		}
		result = json.RawMessage(jsonBytes)
	default:
		log.Println("Unsupported Value type")
		return nil, fmt.Errorf("unsupported Value type")
	}

	return result, nil
}

// 辅助函数：将 Go 的 `any` 类型转换为 gRPC 的 `*pb.Value`
func anyToProtoValue(val any) (*pb.Value, error) {
	switch v := val.(type) {
	case string:
		return &pb.Value{Kind: &pb.Value_StringValue{StringValue: v}}, nil
	case int:
		return &pb.Value{Kind: &pb.Value_IntValue{IntValue: int64(v)}}, nil
	case bool:
		return &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: v}}, nil
	case float64:
		return &pb.Value{Kind: &pb.Value_DoubleValue{DoubleValue: v}}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}
