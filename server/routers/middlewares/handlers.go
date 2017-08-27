package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"

	"ChineseChess/server/models/cache"
	. "ChineseChess/server/routers/render"
)

const (
	JwtTokenHttpHeaderName = "J-T"           // the name of jwt-token inside Http header
	CurrentUserIDKey       = "currentUserID" // the ID of current user
)

var Handlers []gin.HandlerFunc = []gin.HandlerFunc{

	jsonHandler,
	sessionHandler,
}

// jwtHandler is a middleware handler which validate jwt-token
func sessionHandler(c *gin.Context) {

	token := c.GetHeader(JwtTokenHttpHeaderName)
	if err := ValidateToken(token); err != nil {
		c.Error(err)
		RenderErr(c, errors.New("未登录或登录已过期"), 401)
		return
	}
	session, err := cache.FindSession(GetUserID(c))

	if err != nil {
		c.Error(err)
		RenderErr(c, errors.New("未登录或登录已过期"), 401)

	} else {
		c.Set(CurrentUserIDKey, session.UserID)
	}
}

func jsonHandler(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.Header("Content-Encoding", "UTF-8")
}
