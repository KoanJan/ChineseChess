package v1

import (
	"github.com/kataras/iris/context"

	"ChineseChess/server/daf"
	"ChineseChess/server/models"
)

type userForm struct {
	Username string `json:"username"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

// CreateUser creates a user
func CreateUser(ctx context.Context) {

	form := new(userForm)
	if err := ctx.ReadJSON(form); err != nil {
		ctx.WriteString(err.Error())
		return
	}
	user := models.NewUser()
	user.Username = form.Username
	user.Nick = form.Nick
	user.Password = form.Password
	if err := daf.Insert(user); err != nil {
		ctx.WriteString(err.Error())
	}
}
