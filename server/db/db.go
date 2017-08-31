package db

import (
	"gopkg.in/mgo.v2"

	"ChineseChess/server/conf"
)

var sessPool *pool = newPool()

// 执行数据库操作
func Do(collectionName string, f func(*mgo.Collection)) {

	c := sessPool.get()
	defer c.release()
	if err := c.sess.Ping(); err != nil {
		panic(err)
	}
	c.check()
	f(c.sess.DB(*conf.AppConf.Mongodb.DBName).C(collectionName))
}
