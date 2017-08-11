package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/db"
)

type common struct {
	ID bson.ObjectId `bson: "_id", json: "id,string"`
}

// 集合名称
func (this *common) CollectionName() string {
	return ""
}

// 插入
func (this *common) Save() (err error) {
	db.Do(this.CollectionName(), func(c *mgo.Collection) {
		err = c.Insert(this)
	})
	return
}

// 修改
func (this *common) Update() (err error) {
	db.Do(this.CollectionName(), func(c *mgo.Collection) {
		err = c.UpdateId(this.ID, this)
	})
	return
}
