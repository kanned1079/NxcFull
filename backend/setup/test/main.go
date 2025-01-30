package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println(context.Background())
	log.Println(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	log.Println(ctx)
}
