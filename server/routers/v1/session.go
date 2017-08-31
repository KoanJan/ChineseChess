package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/daf"
	"ChineseChess/server/logger"
	"ChineseChess/server/models"
	"ChineseChess/server/models/cache"
	"ChineseChess/server/routers/middlewares"
	. "ChineseChess/server/routers/render"
)

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login function will be called while user logining
func Login(c *gin.Context) {

	form := new(loginForm)
	if err := c.BindJSON(form); err != nil {
		logger.Error(err)
		RenderErr(c, err, 400)
		return
	}
	user := new(models.User)
	if err := daf.FindOne(user, bson.M{"username": form.Username, "password": form.Password}); err != nil {
		logger.Error(err)
		RenderErr(c, err)
		return
	}
	if !user.ID.Valid() {
		RenderErr(c, errors.New("用户名或密码错误"), 401)
		return
	}
	token, err := middlewares.GenerateToken(user.ID.Hex())
	if err != nil {
		logger.Error(err)
		RenderErr(c, errors.New("登陆失败"))
		return
	}
	c.Header(middlewares.JwtTokenHttpHeaderName, token)
	//
	session := cache.NewSession(user.ID.Hex(), user.Nick, cache.SessionStatusOK)
	if err := session.Save(); err != nil {
		logger.Error(err)
		RenderErr(c, err)
		return
	}
	RenderOk(c, map[string]interface{}{
		"details": session,
	})
}

// Logout function will be called while user logout
func Logout(c *gin.Context) {

	if err := cache.DelSession(c.GetString(middlewares.CurrentUserIDKey)); err != nil {
		RenderErr(c, err)
		return
	}
	RenderOk(c)
}
