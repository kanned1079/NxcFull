package main

import (
	"receive/model"
	"sync"
)

func init() {
	model.InitMq()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		model.StartConsumeOrders()

	}()

	wg.Wait()
}
