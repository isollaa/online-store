package registry

import "github.com/gin-gonic/gin"

type RouteContract interface {
	Endpoints(e *gin.Engine)
}

var _routes []RouterFactory

type RouterFactory func() RouteContract

func RegisterRouter(router RouterFactory) {
	_routes = append(_routes, router)
}

func LoadRouter() []RouterFactory {
	return _routes
}
