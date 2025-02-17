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

const (
	OrangeBold = "\033[1;38;5;214m"
	CyanBold   = "\033[1;36m" // 青色加粗
	ResetStyle = "\033[0m"    // 重置样式
)

func init() {
	//if err := config.LocalCfg.Get("./config/db_credential.yaml"); err != nil {
	//	log.Println(err)
	//	return
	//}
	if err := config.LocalCfg.GetFromArgs(); err != nil {
		log.Println(err)
		return
	}
	utils.ShowOK()
	fmt.Println("Successfully filled or loaded ./config/db_credential.yaml")
}

func main() {

	width, err := utils.GetTerminalWidthUnix()
	if err != nil {
		fmt.Printf("无法获取终端宽度: %v\n", err)
		return
	}
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
	fmt.Println("3. Show you the management panel's default secure path.")

	fmt.Print("### ENTER STEP1 ")
	for i := 0; i < width-16; i++ {
		fmt.Printf("#")
	}
	fmt.Println()
	fmt.Print("Please enter the database name: ")

	// 获取用户输入的数据库名
	reader := bufio.NewReader(os.Stdin)
	dbName, _ := reader.ReadString('\n')
	config.LocalCfg.Db.Database = strings.TrimSpace(dbName)

	if err := dao.InitMysqlServer(); err != nil {
		utils.ShowFailure()
		log.Println("Failed to connect database, please check the database configuration.")
		return
	} else {
		utils.ShowOK()
		fmt.Println("Successfully connected to server, preparing to create tables.")
	}

	if err := setup.StartAutoMigrate(); err != nil {
		utils.ShowFailure()
		log.Println("Failed to start auto migration, please check the database configuration.")
		return
	}
	utils.ShowOK()
	fmt.Println("Successfully migrated all tables.")

	// 清除数据
	utils.ShowAttention()
	fmt.Print("Settings table will be truncated (type \"yes\" to continue): ")
	// 获取确认输入
	confirmationClear, _ := reader.ReadString('\n')
	confirmationClear = strings.TrimSpace(confirmationClear)

	if confirmationClear != "yes" {
		utils.ShowFailure()
		fmt.Println("User cancelled process.")
		return
	}
	if err := utils.ClearTable(model.SiteSetting{}.TableName()); err != nil {
		utils.ShowFailure()
		fmt.Println("Failed to clear table, please check the database configuration.")
		return
	}
	// 读取csv数据
	fmt.Println("Start reading default configuration file.")
	settings, err := utils.ReadCSVData("./config/x_site_settings_202501270004.csv")
	if err != nil {
		utils.ShowFailure()
		log.Println("Failed to read settings, please check the database configuration.", err)
		return
	}
	utils.ShowOK()
	fmt.Println("Successfully read settings file, ready to insert into db.")
	//path := ""
	//_ = json.Unmarshal(settings[19].Value, &path)

	if err := utils.InsertDataToDB(settings); err != nil {
		utils.ShowFailure()
		log.Println("Failed to insert settings, please check the database configuration.")
		return
	} else {
		utils.ShowOK()
		fmt.Println("Successfully inserted settings to database.")
	}
	utils.ShowOK()
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
	fmt.Printf("The admin email you entered is: %s%s%s\n", CyanBold, adminEmail, ResetStyle)
	fmt.Printf("The admin password you entered is: %s%s%s\n", CyanBold, adminPassword, ResetStyle)
	utils.ShowAttention()
	fmt.Print("Please confirm if you want to proceed (type \"yes\" to continue): ")

	// 获取确认输入
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(confirmation)

	if confirmation != "yes" {
		utils.ShowFailure()
		fmt.Println("User cancelled process.")
		return
	}
	if err := setup.SetupNewAdminAccount(adminEmail, adminPassword); err != nil {
		utils.ShowFailure()
		log.Println("Failed to setup new admin account, please check the database configuration.")
		return
	}
	utils.ShowOK()
	fmt.Println("Successfully setup new admin account.")
	utils.ShowOK()
	fmt.Println("Your database initialization has been completed. Please follow the documentation to continue.")
	//fmt.Printf("secure path: %s%s%s\n", OrangeBold, string(settings[18].Value), ResetStyle)

	utils.ShowAttention()
	fmt.Printf("The string of your management backend security path is %s%s%s. To ensure system security, you will need to recharge this security path in the system configuration later.\n", OrangeBold, string(settings[18].Value), ResetStyle)
}
