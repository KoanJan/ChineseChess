package daf

import (
	"gopkg.in/mgo.v2"

	"ChineseChess/server/db"
	"ChineseChess/server/models"
)

// 插入
func Insert(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.Insert(model)
	})
	return
}

// 更新
func Update(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.UpdateId(model.GetID(), model)
	})
	return
}

// 删除
func Delete(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.RemoveId(model.GetID())
	})
	return
}

// 根据ID查找
func Find(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.FindId(model.GetID()).One(model)
	})
	return
}
