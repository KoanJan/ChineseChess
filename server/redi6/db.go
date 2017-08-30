package redi6

import (
	"github.com/garyburd/redigo/redis"
)

// Do executes a redis command
func Do(commandName string, args ...interface{}) (interface{}, error) {

	conn := pool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

// Hset executes the command 'HSET' of Redis
func Hset(key, field, v interface{}) (interface{}, error) {

	return Do("HSET", key, field, v)
}

// Hget executes the command 'HGET' of Redis
func Hget(key, field interface{}) (interface{}, error) {

	return Do("HGET", key, field)
}

// Hdel executes the command 'HDEL' of Redis
func Hdel(key, field interface{}) (interface{}, error) {

	return Do("HDEL", key, field)
}

// Hmset executes the command 'HMSET' of Redis
func Hmset(args ...interface{}) (interface{}, error) {

	return Do("HMSET", args...)
}

// Hgetall returns a 'interface{}'
func Hgetall(key interface{}) (interface{}, error) {

	return Do("HGETALL", key)
}

// Set executes the command 'SET' of Redis
func Set(args ...interface{}) (interface{}, error) {

	return Do("SET", args...)
}

// Get executes the command 'GET' of Redis
func Get(key interface{}) (interface{}, error) {

	return Do("GET", key)
}

// Del executes the command 'DEL' of Redis
func Del(key interface{}) (interface{}, error) {

	return Do("DEL", key)
}

// Expire executes the command 'EXPIRE' of Redis
func Expire(key, seconds interface{}) (interface{}, error) {

	return Do("EXPIRE", key, seconds)
}

// EvalLuaScript executes the command 'EVAL' of Redis with a lua script
func EvalLuaScript(keyCount int, src string, keysAndArgs ...interface{}) (interface{}, error) {

	script := redis.NewScript(keyCount, src)
	return script.Do(pool.Get(), keysAndArgs...)
}
