package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 用户
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id,string"`         // ID
	Username  string        `bson:"username" json:"username"`     // 登录名
	Nick      string        `bson:"nick" json:"nick"`             // 昵称
	Password  string        `bson:"password" json:"-"`            // 密码
	CreatedAt time.Time     `bson:"created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"` // 修改时间
}

func (this *User) CollectionName() string {
	return "users"
}

func (this *User) GetID() bson.ObjectId {
	return this.ID
}

// NewUser returns a new user
func NewUser() *User {

	user := new(User)
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return user
}
