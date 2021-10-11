package registry

import "github.com/gin-gonic/gin"

//declare route interface
type RouteContract interface {
	Endpoints(e *gin.Engine)
}

//attached route will be stored here
var routes []RouterFactory

type RouterFactory func() RouteContract

//attach route into routes
func RegisterRouter(route RouterFactory) {
	routes = append(routes, route)
}

//get list attached route
func LoadRouter() []RouterFactory {
	return routes
}
