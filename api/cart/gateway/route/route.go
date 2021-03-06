package route

import (
	"online-store/api/cart"
	"online-store/api/cart/gateway/handler"
	"online-store/middleware"
	"online-store/registry"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Engine *gin.Engine
}

type service struct {
	Handler handler.HandlerContract
}

func (svc *service) Endpoints(e *gin.Engine) {
	e.GET("/cart", middleware.Auth, svc.Handler.GetList)
	e.POST("/cart", middleware.Auth, svc.Handler.Create)
	e.DELETE("/cart/:id", middleware.Auth, svc.Handler.Void)
}

//fill injected hander to route
func New() registry.RouteContract {
	return &service{
		Handler: cart.ApiHandler(),
	}
}

//make sure to call this package to register this route
func init() {
	registry.RegisterRouter(New)
}
