package main

import (
	"NxcFull/nxc-backend/coupon"
	"NxcFull/nxc-backend/dao"
	"NxcFull/nxc-backend/routers"
	"NxcFull/nxc-backend/settings"
	"NxcFull/nxc-backend/user"
)

type adminConfig struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	IsAdmin  bool   `yaml:"is_admin"`
	Password string `yaml:"password"`
}

func main() {

	//ginLog, _ := os.Create("gin.log")
	//defer func() {
	//	if err := ginLog.Close(); err != nil {
	//		log.Println("关闭gin日志失败")
	//	}
	//}()
	//
	//gin.DefaultWriter = ginLog

	// 自动迁移
	if err := dao.Db.AutoMigrate(&settings.SiteSetting{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&user.User{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&user.Auth{}); err != nil {
		panic("failed to migrate database")
	}

	if err := dao.Db.AutoMigrate(&coupon.Coupon{}); err != nil {
		panic("failed to migrate database")
	}

	dao.Db.Model(&coupon.Coupon{}).Create(&coupon.Coupon{
		Title: "2024夏季优惠券发放 20%OFF",
		Content: `优惠券码：HappySummer2024Pre
			此优惠码使用截至日期2024-08-31，主站点与常州站点同步皆可使用，请阁下在确认订单时在页面上方填入此优惠码。祝大家2024夏天快乐！`,
		ImgUrl: "",
		Tags:   "coupons",
		Show:   true,
	})

	// 添加管理员
	//func() {
	//	f, err := os.OpenFile("./config.yaml", os.O_RDONLY, 0444)
	//	if err != nil {
	//		log.Fatal("配置文件打开错误")
	//	}
	//	defer func() {
	//		if err := f.Close(); err != nil {
	//			panic(err)
	//		}
	//	}()
	//	// 插入读取yaml文件进行序列化
	//	var configData map[string]any
	//	decoder := yaml.NewDecoder(f)
	//	if err := decoder.Decode(&configData); err != nil {
	//		log.Print("解析yaml失败")
	//	}
	//	if result := dao.Db.Model(&user.User{}).Where("is_admin IS NULL"); result.Error != nil {
	//		log.Print("管理员不存在")
	//		var encryptor encrypt.Encryptor
	//		var admin user.User = user.User{
	//			Name:    "administrator",
	//			Email:   "admin@ikanned.com",
	//			IsAdmin: true,
	//		}
	//		var adminAuth = user.Auth{
	//			Email:    "admin@ikanned.com",
	//			Password: encryptor.Base64Encrypt("password"),
	//		}
	//		dao.Db.Model(&user.User{}).Create(&admin)
	//		dao.Db.Model(&user.Auth{}).Create(&adminAuth)
	//		adminConfig := adminConfig{
	//			Name:     admin.Name,
	//			Email:    admin.Email,
	//			IsAdmin:  admin.IsAdmin,
	//			Password: adminAuth.Password,
	//		}
	//		configData["admin"] = adminConfig
	//		fw, err := os.OpenFile("./config.yaml", os.O_RDWR|os.O_CREATE, 0644)
	//		if err != nil {
	//			log.Fatal("无法打开文件进行写入：", err)
	//		}
	//		defer fw.Close()
	//		encoder := yaml.NewEncoder(fw)
	//		if err := encoder.Encode(configData); err != nil {
	//			log.Fatal("写入yaml错误")
	//		}
	//	} else {
	//		log.Print("管理员存在 跳过创建")
	//	}
	//}()

	//dao.Db.Create(&user.User{
	//	Name:              "kanna",
	//	Email:             "admin@ikanned.com",
	//	IsAdmin:           true,
	//	Balance:           29.8,
	//	LastLogin:         time.Date(2019, time.October, 16, 7, 45, 0, 0, time.UTC),
	//	LicenseExpiration: time.Date(2029, time.October, 16, 7, 45, 0, 0, time.UTC),
	//})
	//dao.Db.Create(&user.Auth{
	//	Email:    "admin@ikanned.com",
	//	Password: "cGFzc3dvcmQ=",
	//})

	routers.StartAdminReq()
}
