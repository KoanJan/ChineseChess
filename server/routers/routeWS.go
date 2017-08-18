package routers

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/routers/middlewares"
	"ChineseChess/server/routers/ws"
)

func routeWS(e *gin.Engine) {

	e.GET("/api/ws", ws.WS).Use(middlewares.Handlers[1]) // 获取WebSocket连接
}
