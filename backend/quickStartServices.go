package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var services = []string{
	"couponServices",
	"documentServices",
	"groupServices",
	"mailServices",
	"noticeServices",
	"subscribeServices",
	"systemServices",
	"userServices",
	"orderServices",
	"orderHandleServices",
	"keyServices",
}

var gatewayService = "gateway"
var pidFile = "services_pids.txt"

// 启动单个服务
func startService(service string, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	// 切换到服务目录并启动服务
	cmd := exec.Command("go", "run", "command/server/main.go")
	cmd.Dir = service

	// 设置日志输出到控制台
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 启动服务
	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to start service %s: %v\n", service, err)
		return
	}

	// 保存 PID
	mu.Lock()
	appendToPidFile(service, cmd.Process.Pid)
	mu.Unlock()

	log.Printf("Service %s started with PID %d\n", service, cmd.Process.Pid)

	// 等待进程结束
	err = cmd.Wait()
	if err != nil {
		log.Printf("Service %s terminated with error: %v\n", service, err)
	}
}

// 追加 PID 到文件
func appendToPidFile(service string, pid int) {
	f, err := os.OpenFile(pidFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open PID file: %v\n", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s: %d\n", service, pid))
	if err != nil {
		log.Fatalf("Failed to write to PID file: %v\n", err)
	}
}

// 启动所有服务
func startServices() {
	// 清空 PID 文件
	err := os.WriteFile(pidFile, []byte{}, 0644)
	if err != nil {
		log.Fatalf("Failed to clear PID file: %v\n", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	// 并发启动服务
	for _, service := range services {
		wg.Add(1)
		fmt.Printf("---%s-------------------------------------------------------\n", service)
		go startService(service, &wg, &mu)
	}

	// 启动 gateway 服务，等待 10 秒后启动
	log.Println("Waiting 10 seconds to start gateway service...")
	time.Sleep(10 * time.Second)

	wg.Add(1)
	go startService(gatewayService, &wg, &mu)

	wg.Wait()
	log.Println("All services started.")
}

// 停止服务
func stopServices() {
	// 读取 PID 文件
	content, err := os.ReadFile(pidFile)
	if err != nil {
		log.Fatalf("Failed to read PID file: %v\n", err)
	}

	// 将内容按行分割
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line) // 去掉两边空格
		if line == "" {
			continue // 跳过空行
		}

		// 手动分割服务名称和 PID
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Printf("Failed to parse line: input does not match format | Line: %s\n", line)
			continue
		}

		service := strings.TrimSpace(parts[0])
		pidStr := strings.TrimSpace(parts[1])
		pid, err := strconv.Atoi(pidStr) // 将 PID 转换为整数
		if err != nil {
			log.Printf("Failed to convert PID to integer for service %s: %v\n", service, err)
			continue
		}

		// 发送信号终止进程
		process, err := os.FindProcess(pid)
		if err != nil {
			log.Printf("Failed to find process for service %s: %v\n", service, err)
			continue
		}

		err = process.Kill() // 你也可以使用 process.Signal 来发送其他信号
		if err != nil {
			log.Printf("Failed to kill process for service %s: %v\n", service, err)
		} else {
			log.Printf("Service %s with PID %d stopped.\n", service, pid)
		}
	}

	// 删除 PID 文件
	err = os.Remove(pidFile)
	if err != nil {
		log.Fatalf("Failed to delete PID file: %v\n", err)
	}
}

// 捕获 SIGINT 信号并优雅地停止服务
func handleSignal() {
	c := make(chan os.Signal, 1)
	// 捕获 SIGINT（^C） 和 SIGTERM
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// 启动 goroutine 来等待信号
	go func() {
		sig := <-c
		log.Printf("Received signal: %v. Stopping services...\n", sig)
		stopServices() // 在收到信号时停止服务
		os.Exit(0)     // 退出程序
	}()
}

func main() {
	handleSignal() // 启动信号处理

	if len(os.Args) > 1 && os.Args[1] == "stop" {
		stopServices()
	} else {
		startServices()
	}
}
