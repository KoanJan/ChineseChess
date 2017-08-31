package models

import (
	"gopkg.in/mgo.v2/bson"
)

// 模型接口
type Model interface {
	CN() string           // collection name
	GetID() bson.ObjectId // ID
}
