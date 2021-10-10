package main

import (
	"online-store/registry"

	"github.com/gin-gonic/gin"

	//init endpoints
	_ "online-store/api/user/gateway/route"

	_ "online-store/api/item/gateway/route"

	_ "online-store/api/cart/gateway/route"
)

func main() {
	e := gin.Default()
	for _, router := range registry.LoadRouter() {
		router().Endpoints(e)
	}

	e.Run(":3000")
}
