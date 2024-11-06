package main

import (
	"book-mgr-backend/dao"
	"book-mgr-backend/routers"
)

func init() {
	dao.InitMysqlServer()
}

func main() {
	var app routers.App
	app.RunServer()
}
