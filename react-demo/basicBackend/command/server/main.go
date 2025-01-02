package main

import "basicBackend/routers"

func init() {
	routers.App = routers.NewGateway(1)
}

func main() {
	routers.App.RunApiGateway()
}
