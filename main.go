package main

import (
	"online-store/registry"

	"github.com/gin-gonic/gin"

	//load endpoints by import feature routes
	_ "online-store/api/user/gateway/route"

	_ "online-store/api/item/gateway/route"

	_ "online-store/api/cart/gateway/route"
)

func main() {
	e := gin.Default()

	//load all attached route by init
	for _, router := range registry.LoadRouter() {
		router().Endpoints(e)
	}

	e.Run(":3000")
}
