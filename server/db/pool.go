package db

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"

	"ChineseChess/server/conf"
)

const (
	poolSize = 20
	timeout  = 3
)

// 连接池
type pool struct {
	conns []*conn
	sync.Mutex
}

// 获取连接
func (p *pool) get() *conn {
	p.Lock()
	defer p.Unlock()
	conn := p.conns[0]
	p.conns = p.conns[1:]
	return conn
}

func newPool() *pool {

	p := new(pool)
	p.conns = make([]*conn, poolSize)

	sess, err := mgo.Dial(*conf.AppConf.Mongodb.URL)
	defer sess.Close()
	sess.SetSocketTimeout(timeout * time.Second)
	if err != nil {
		panic(err)
	}
	for i := 0; i < poolSize; i++ {

		p.conns[i] = &conn{sess.Clone(), p}
	}
	return p
}
