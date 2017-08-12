package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 模型接口
type Model interface {
	CollectionName() string
	GetID() bson.ObjectId
}

// 公共属性集合
type Common struct {
	ID        bson.ObjectId `bson:"_id",json:"id,string"`         // ID
	CreatedAt time.Time     `bson:"created_at",json:"created_at"` // 创建时间
	UpdatedAt time.Time     `bson:"updated_at",json:"updated_at"` // 修改时间
}

// 获取ID
func (this *Common) GetID() bson.ObjectId {
	return this.ID
}
