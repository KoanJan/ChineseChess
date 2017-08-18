package ws

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/routers/middlewares"
)

// WS handler
func WS(c *gin.Context) {

	serveWS(c.GetString(middlewares.CurrentUserIDKey), c.Writer, c.Request)
}
