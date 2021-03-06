package redi6

import (
	"time"

	"github.com/garyburd/redigo/redis"

	"ChineseChess/server/conf"
)

var pool *redis.Pool

func init() {

	pool = &redis.Pool{

		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", conf.AppConf.Redis.Address)
			if err != nil {
				return nil, err
			}
			if _, err = c.Do("SELECT", conf.AppConf.Redis.Select); err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {

			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     20,
		MaxActive:   5,
		IdleTimeout: 30 * time.Second,
		Wait:        true,
	}
}
