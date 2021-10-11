package route

import (
	user "online-store/api/item"
	"online-store/api/item/gateway/handler"
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
	e.GET("/item", middleware.Auth, svc.Handler.GetList)
	e.POST("/item", middleware.Auth, svc.Handler.Create)
	e.PUT("/item/:id", middleware.Auth, svc.Handler.Update)
	e.DELETE("/item/:id", middleware.Auth, svc.Handler.Delete)
}

//fill injected hander to route
func New() registry.RouteContract {
	return &service{
		Handler: user.ApiHandler(),
	}
}

//make sure to call this package to register this route
func init() {
	registry.RegisterRouter(New)
}
