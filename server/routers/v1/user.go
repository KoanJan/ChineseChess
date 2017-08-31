package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/daf"
	"ChineseChess/server/logger"
	"ChineseChess/server/models"
	"ChineseChess/server/routers/middlewares"
	. "ChineseChess/server/routers/render"
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

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	if middlewares.GetUserID(c) != id {
		RenderErr(c, errors.New("权限不足"), 401)
		return
	}
	if !bson.IsObjectIdHex(id) {
		RenderErr(c, errors.New("id不合法"), 400)
		return
	}
	data := bson.M{}
	if err := c.BindJSON(&data); err != nil {
		logger.Error(err)
		RenderErr(c, errors.New("数据解析失败"), 400)
		return
	}
	// 过滤掉禁止修改的属性
	forbidden := []string{"username", "password"}
	for _, key := range forbidden {
		delete(data, key)
	}
	// 更新
	if err := daf.UpdateM(models.UserCN(), bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": data}); err != nil {
		RenderErr(c, errors.New("请求失败"), 500)
		return
	}
	RenderOk(c)
}
