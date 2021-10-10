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
	// e.GET("/item/:id", middleware.Auth, svc.Handler.Get)
	e.GET("/item", middleware.Auth, svc.Handler.GetList)
	e.POST("/item", middleware.Auth, svc.Handler.Create)
	e.PUT("/item/:id", middleware.Auth, svc.Handler.Update)
	e.DELETE("/item/:id", middleware.Auth, svc.Handler.Delete)
}

func New() registry.RouteContract {
	return &service{
		Handler: user.ApiHandler(),
	}
}

func init() {
	registry.RegisterRouter(New)
}
