package db

import (
	"gopkg.in/mgo.v2"

	"ChineseChess/server/conf"
	_ "ChineseChess/server/conf"
)

// 连接
type conn struct {
	sess *mgo.Session
	pool *pool
}

// 释放
func (c *conn) release() {

	c.pool.Lock()
	defer c.pool.Unlock()
	c.pool.conns = append(c.pool.conns, c)
}

// 每次使用时先自我检查
func (c *conn) check() {
	if c.sess.Ping() != nil {
		var err error
		if c.sess, err = mgo.Dial(*conf.AppConf.Mongodb.URL); err != nil {
			panic(err)
		}
	}
}
