package routers

import (
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {

	routeV1(engine) // v1

	routeWS(engine) // websocket
}
