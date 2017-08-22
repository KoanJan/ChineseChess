package daf

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/db"
	"ChineseChess/server/models"
)

// Insert can insert a model into db
func Insert(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.Insert(model)
	})
	return
}

// Update can update a model
func Update(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.UpdateId(model.GetID(), model)
	})
	return
}

// Delete can delete a model
func Delete(model models.Model) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.RemoveId(model.GetID())
	})
	return
}

// Find can find a model.
func FindOne(model models.Model, m bson.M) (err error) {
	db.Do(model.CollectionName(), func(c *mgo.Collection) {
		err = c.Find(m).One(model)
	})
	return
}
