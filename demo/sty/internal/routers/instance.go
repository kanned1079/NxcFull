package routers

import "github.com/gin-gonic/gin"

type ApiInstance struct {
	Id     int64
	Router *gin.Engine
}

func NewApiInstance(id int64) *ApiInstance {
	return &ApiInstance{
		Id:     id,
		Router: gin.Default(),
	}
}
