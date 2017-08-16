package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

/*
用户
*/
type User struct {
	Common
	Username string `bson:"username",json:"username"` // 登录名
	Nick     string `bson:"nick",json:"nick"`         // 昵称
	Password string `bson:"password",json:"-"`        // 密码
}

func (this *User) CollectionName() string {
	return "users"
}

// NewUser returns a new user
func NewUser() *User {

	user := new(User)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user
}
