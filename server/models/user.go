package models

/*
用户
*/
type User struct {
	Common
	Username string `bson:"username",json:"username"` // 登录名
	Nick     string `bson:"nick",json:"nick"`         // 昵称
}

func (this *User) CollectionName() string {
	return "users"
}
