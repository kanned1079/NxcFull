package main

import (
	"log"
	"os"
)

func initLogWriter() {
	// 打开一个文件来写日志
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将日志输出目标设置为文件
	log.SetOutput(file)

	// 写日志
	log.Println("This is a log message")
	log.Println("Another log message")
}
