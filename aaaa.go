package main

import "github.com/gin-gonic/gin"

type AAA struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	a := AAA{
		Name: "666",
		Age:  18,
	}
	r.GET("/aa", func(c *gin.Context) {
		c.JSON(200, a)
	})
	r.Run("localhost:8081")
}
