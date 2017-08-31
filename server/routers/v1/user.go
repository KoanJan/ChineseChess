package v1

import (
	"github.com/gin-gonic/gin"

	"ChineseChess/server/daf"
	"ChineseChess/server/models"
	. "ChineseChess/server/routers/render"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type userForm struct {
	Username string `json:"username"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

// CreateUser creates a user
func CreateUser(c *gin.Context) {

	form := new(userForm)
	if err := c.BindJSON(form); err != nil {
		RenderErr(c, err, 400)
		return
	}
	user := models.NewUser()
	user.Username = form.Username
	user.Nick = form.Nick
	user.Password = form.Password
	if err := daf.Insert(user); err != nil {
		RenderErr(c, err)
		return
	}
	RenderOk(c)
}

// GetUser get user's infomation
func GetUser(c *gin.Context) {

	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		RenderErr(c, errors.New("id不合法"), 400)
		return
	}
	user := new(models.User)
	if err := daf.FindOne(user, bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		RenderErr(c, err)
		return
	}
	RenderOk(c, map[string]interface{}{
		"details": user,
	})
}
