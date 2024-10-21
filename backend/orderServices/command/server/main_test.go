package main

import (
	"log"
	"orderServices/internal/config/remote"
	"orderServices/internal/dao"
	"testing"
)

func TestName(t *testing.T) {
	log.Println("test")
	dao.InitMq()
	log.Println(remote.MyMqConfig)
}
