package redi6

import (
	"errors"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisKeyLockError error = errors.New("redis锁失败")
)

// Do executes a redis command
func Do(commandName string, args ...interface{}) (interface{}, error) {

	conn := pool.Get()
	return conn.Do(commandName, args...)
}

// Hset executes the command 'HSET' of Redis
func Hset(key, field, v interface{}) (interface{}, error) {

	return Do("hset", key, field, v)
}

// Hget executes the command 'HGET' of Redis
func Hget(key, field interface{}) (interface{}, error) {

	return Do("hget", key, field)
}

// Hdel executes the command 'HDEL' of Redis
func Hdel(key, field interface{}) (interface{}, error) {

	return Do("hdel", key, field)
}

// Hmset executes the command 'HMSET' of Redis
func Hmset(args ...interface{}) (interface{}, error) {

	return Do("hmset", args...)
}

// Hgetall returns a 'interface{}'
func Hgetall(key interface{}) (interface{}, error) {

	return Do("hgetall", key)
}

// Set executes the command 'SET' of Redis
func Set(key, v interface{}) (interface{}, error) {

	return Do("set", key, v)
}

// Get executes the command 'GET' of Redis
func Get(key interface{}) (interface{}, error) {

	return Do("get", key)
}

// Del executes the command 'DEL' of Redis
func Del(key interface{}) (interface{}, error) {

	return Do("del", key)
}

// Setnx executes the command 'SET [key] [value] NX PX [expire]' of Redis
func Setnxpx(key, v interface{}, expire int32) error {

	if r, e := Do("set", key, v, "nx px", expire); e == nil {
		if rv, ok := r.(int); ok {
			if rv == 1 {
				return nil
			}
		}
	}
	return RedisKeyLockError
}

// Expire executes the command 'EXPIRE' of Redis
func Expire(key, seconds interface{}) (interface{}, error) {

	return Do("expire", key, seconds)
}

// EvalLuaScript executes the command 'EVAL' of Redis with a lua script
func EvalLuaScript(keyCount int, src string, keysAndArgs ...interface{}) (interface{}, error) {

	script := redis.NewScript(keyCount, src)
	return script.Do(pool.Get(), keysAndArgs...)
}
