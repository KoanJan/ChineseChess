package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"

	. "ChineseChess/server/routers/render"
)

const JwtTokenHttpHeaderName = "token"

var Handlers []gin.HandlerFunc = []gin.HandlerFunc{

	jwtHandler,
}

// jwtHandler is a middleware handler which validate jwt-token
func jwtHandler(c *gin.Context) {

	token := c.GetHeader(JwtTokenHttpHeaderName)
	if err := ValidateToken(token); err != nil {
		c.Error(err)
		RenderErr(c, errors.New("未登录或登录已过期"), 401)
		return
	}
}
