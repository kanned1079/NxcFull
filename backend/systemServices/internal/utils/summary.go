package utils

import (
	"log"
	"os"
	"runtime"
	"systemServices/internal/dao"
	"time"
)

//func Get() {
//
//}

func GetSystemSummaryConfAndStatus(args ...string) (
	timeNow string, // 服务器时间
	dbStatus bool, // 数据库连接状态
	redisStatus bool, // redis连接状态
	EtcdStatus bool, // Etcd连接状态
	OsType string, // 操作系统类型
	OsArch string, // 操作系统架构
	NumCPU int, // 主机CPU个数
) {
	//var dbConnectionsStatus []string = []string{
	//	"Threads_connected",    // 当前活跃的连接数，即客户端与数据库的连接数。
	//	"Threads_running",      // 当前正在运行的线程数（非空闲状态）。
	//	"Connections",          // 自数据库启动以来的总连接尝试数，包括成功和失败的连接。
	//	"Aborted_connects",     // 失败的连接尝试总数，可能指示连接池配置问题或恶意尝试。
	//	"Max_used_connections", // 自启动以来使用的最大连接数，反映了高峰期的连接需求。
	//}
	//var redisConnectionsStatus []string = []string{
	//	""
	//}
	time.Now().Format("2006-01-02 15:04:05")
	//log.Println(dbConnectionsStatus)
	log.Println(os.Hostname())
	log.Println(runtime.GOOS)
	log.Println(runtime.GOARCH)
	log.Println(runtime.NumCPU())

	dbStatus = true
	redisStatus = true
	if err := dao.Db.Exec("SELECT 1+1;").Error; err != nil {
		dbStatus = false
	}

	//return time.Now().String(), dbStatus, true, true, OsType, OsArch, NumCPU

	return
}
