package route

import (
	"online-store/api/user"
	"online-store/api/user/gateway/handler"
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
	e.POST("/login", svc.Handler.Login)
	e.GET("/user", middleware.Auth, svc.Handler.GetList)
	e.POST("/user", svc.Handler.Create)
	e.PUT("/user/:id", middleware.Auth, svc.Handler.Update)
	e.DELETE("/user/:id", middleware.Auth, svc.Handler.Delete)
}

func New() registry.RouteContract {
	return &service{
		Handler: user.ApiHandler(),
	}
}

func init() {
	registry.RegisterRouter(New)
}
