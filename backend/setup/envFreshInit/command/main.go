package main

import (
	"bufio"
	"envFreshInit/internal/config"
	"envFreshInit/internal/dao"
	"envFreshInit/internal/model"
	"envFreshInit/internal/setup"
	"envFreshInit/internal/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

// StartInitDb 假设这些函数已经定义好，具体实现会根据你的数据库和 ORM 来实现
func StartInitDb(dbName string) error {
	// 这里是数据库初始化的逻辑
	fmt.Printf("数据库 '%s' 初始化完成。\n", dbName)
	return nil
}

func CreateNewAdminUser(email, password string) error {
	// 这里是创建管理员的逻辑
	fmt.Printf("管理员账户创建成功！邮箱：%s，密码：%s\n", email, password)
	return nil
}

// 根据语言选择显示对应的提示
//func getLanguagePrompt(lang string) (func(), string) {
//	if lang == "1" {
//		// 中文提示
//		return promptInChinese, "表初始化完成！请输入管理员信息："
//	} else {
//		// 英文提示
//		return promptInEnglish, "Table initialization completed! Please enter admin details:"
//	}
//}

func putPrompt(lang string, chPrompt, enPrompt string) {
	if lang == "1" {
		fmt.Print(chPrompt)
	} else {
		fmt.Print(enPrompt)
	}
}

func init() {
	if err := config.LocalCfg.Get("./config/db_credential.yaml"); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	// 0.选择中文还是英文
	//lang := chooseLanguage()

	// 根据语言显示提示
	//promptFunc, successMessage := getLanguagePrompt(lang)
	//fmt.Println("Table initialization completed! Please enter admin details:")

	// 显示初始化提示
	//promptFunc()

	fmt.Println("This is an initialization program. It will perform the following actions:")
	fmt.Println("1. Create all database tables.")
	fmt.Println("2. Create a default admin user.")

	fmt.Println("ENTER STEP1---------->")
	fmt.Print("Please enter the database name: ")

	// 获取用户输入的数据库名
	reader := bufio.NewReader(os.Stdin)
	dbName, _ := reader.ReadString('\n')
	config.LocalCfg.Db.Database = strings.TrimSpace(dbName)

	if err := dao.InitMysqlServer(); err != nil {
		log.Println("Failed to connect database, please check the database configuration.")
		return
	} else {
		log.Println("Successfully connected to server, preparing to create tables.")
	}

	if err := setup.StartAutoMigrate(); err != nil {
		log.Println("Failed to start auto migration, please check the database configuration.")
		return
	}
	fmt.Println("Successfully migrated all tables.")

	// 清除数据
	fmt.Print("Settings table will be truncated (type yes to continue):")
	// 获取确认输入
	confirmationClear, _ := reader.ReadString('\n')
	confirmationClear = strings.TrimSpace(confirmationClear)

	if confirmationClear != "yes" {
		fmt.Println("User cancelled process.")
		return
	}
	if err := utils.ClearTable(model.SiteSetting{}.TableName()); err != nil {
		fmt.Println("Failed to clear table, please check the database configuration.")
		return
	}
	// 读取csv数据
	fmt.Println("Start reading default configuration file.")
	settings, err := utils.ReadCSVData("./config/x_site_settings_202501270004.csv")
	if err != nil {
		log.Println("Failed to read settings, please check the database configuration.")
		return
	}
	fmt.Println("Successfully read settings file, ready to insert into db.")

	if err := utils.InsertDataToDB(settings); err != nil {
		log.Println("Failed to insert settings, please check the database configuration.")
		return
	} else {
		log.Println("Successfully inserted settings to database.")
	}

	fmt.Println("Congratulations on completing the first step of creating a database. Next, you need to create a default administrator.")
	fmt.Println("ENTER STEP2---------->")

	// 2.让用户创建一个管理员用户
	//putPrompt(lang, "请输入管理员邮箱: ", "")
	fmt.Print("Please enter administrator's email address: ")
	adminEmail, _ := reader.ReadString('\n')
	adminEmail = strings.TrimSpace(adminEmail)
	fmt.Print("Please enter administrator's password: ")
	adminPassword, _ := reader.ReadString('\n')
	adminPassword = strings.TrimSpace(adminPassword)

	// 输出用户输入的邮箱和密码供检查
	fmt.Printf("The admin email you entered is: %s\n", adminEmail)
	fmt.Printf("The admin password you entered is: %s\n", adminPassword)
	fmt.Print("Please confirm if you want to proceed (type yes to continue):")

	// 获取确认输入
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(confirmation)

	if confirmation != "yes" {
		fmt.Println("User cancelled process.")
		return
	}
	//if err := utils.ClearTable(model.Auth{}.TableName()); err != nil {
	//	fmt.Println("Failed to clear table, please check the database configuration.")
	//	return
	//}
	//if err := utils.ClearTable(model.User{}.TableName()); err != nil {
	//	fmt.Println("Failed to clear table, please check the database configuration.")
	//	return
	//}
	if err := setup.SetupNewAdminAccount(adminEmail, adminPassword); err != nil {
		log.Println("Failed to setup new admin account, please check the database configuration.")
		return
	}
	fmt.Println("Successfully setup new admin account.")
	fmt.Println("Your database initialization has been completed. Please follow the documentation to continue.")

}
