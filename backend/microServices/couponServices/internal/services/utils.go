package services

import (
	pb "couponServices/api/proto"
	"couponServices/internal/model"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Convert2rpcCoupons(original []model.Coupon) []*pb.CouponInfo {
	var result []*pb.CouponInfo
	for _, coupon := range original {
		result = append(result, &pb.CouponInfo{
			Id:         coupon.Id,
			Name:       coupon.Name,
			Code:       coupon.Code,
			Enabled:    coupon.Enabled,
			PercentOff: float32(coupon.PercentOff),
			//StartTime:    coupon.StartTime, // StartTime as int64
			StartTime:    coupon.StartTime.UnixNano() / int64(time.Millisecond),
			EndTime:      coupon.EndTime.UnixNano() / int64(time.Millisecond), // EndTime as int64
			PerUserLimit: coupon.PerUserLimit,
			Capacity:     coupon.Capacity,
			Residue:      coupon.Residue,
			PlanLimit:    coupon.PlanLimit,
			CreatedAt: &pb.Timestamp2{
				Seconds: coupon.CreatedAt.Unix(),
				Nanos:   int32(coupon.CreatedAt.Nanosecond()),
			},
			UpdatedAt: &pb.Timestamp2{
				Seconds: coupon.UpdatedAt.Unix(),
				Nanos:   int32(coupon.UpdatedAt.Nanosecond()),
			},
			DeletedAt: func() *pb.Timestamp2 {
				if coupon.DeletedAt.Valid {
					return &pb.Timestamp2{
						Seconds: coupon.DeletedAt.Time.Unix(),
						Nanos:   int32(coupon.DeletedAt.Time.Nanosecond()),
					}
				}
				return nil
			}(),
		})
	}
	return result
}

// 生成随机12位字符串
func generateRandomCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 检查券码是否已存在
func isCodeExist(db *gorm.DB, code string) bool {
	var count int64
	db.Model(&model.Coupon{}).Where("code = ?", code).Count(&count)
	return count > 0
}
