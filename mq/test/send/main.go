package main

import (
	"encoding/json"
	"log"
	"send/model"
)

func init() {
	model.InitMq()
}

func main() {
	jsonData, _ := json.Marshal(map[string]interface{}{
		"name": "kanna",
	})
	for i := 0; i < 10; i++ {

		if err := model.PublishOrderMessage(jsonData); err != nil {
			log.Println(err)
		}
	}
}
