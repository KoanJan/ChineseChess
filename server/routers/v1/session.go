package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/daf"
	"ChineseChess/server/logger"
	"ChineseChess/server/models"
	"ChineseChess/server/redis"
	"ChineseChess/server/routers/middlewares"
	. "ChineseChess/server/routers/render"
)

const (
	RedisUserSessionExpire = 7 * 24 * 60 * 60
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login function will be called while user logining
func Login(c *gin.Context) {

	logger.Debug("oil login")

	form := new(loginForm)
	if err := c.BindJSON(form); err != nil {
		RenderErr(c, err, 400)
		return
	}
	user := new(models.User)
	if err := daf.FindOne(user, bson.M{"username": form.Username, "password": form.Password}); err != nil {
		RenderErr(c, err)
		return
	}
	if !user.ID.Valid() {
		RenderErr(c, errors.New("用户名或密码错误"), 401)
		return
	}
	token, err := middlewares.GenerateToken(user.ID.Hex())
	if err != nil {
		c.Error(err)
		RenderErr(c, errors.New("登陆失败"))
		return
	}
	c.Header(middlewares.JwtTokenHttpHeaderName, token)
	//
	redis.Set(user.ID.Hex(), middlewares.GenerateSessionString(user))
	redis.Expire(user.ID.Hex(), RedisUserSessionExpire)
	RenderOk(c)
}

// Logout function will be called while user logout
func Logout(c *gin.Context) {

	redis.Del(c.GetString(middlewares.CurrentUserIDKey))
	RenderOk(c)
}
