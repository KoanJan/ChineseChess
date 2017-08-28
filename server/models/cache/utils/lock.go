package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"

	"ChineseChess/server/redi6"
)

const (
	// key format of lock
	LockKeyF = "redis_lock.%s"

	// lua script of unlock
	UnlockScript = `
		if redis.call("get", KEYS[1]) == ARGV[1]
		then
    		return redis.call("del", KEYS[1])
		else
    		return 0
		end
	`

	// default repeat times of lock
	LockReqeat int = 9

	// default expiry of lock
	LockExpiry int32 = 300
)

var RedisSetNxError error = errors.New("redis: SET with 'NX' failed")

// Lock will repeat locking a key of redis, and return a lock token and error
func Lock(key string, expire ...int32) (int64, error) {

	// default expiry of the lock is 3000 millseconds
	var _expire int32 = LockExpiry
	if len(expire) > 0 && expire[0] > 0 {
		_expire = expire[0]
	}
	for i := 0; i < LockReqeat; i++ {
		token := time.Now().Unix()
		c, e := redis.String(redi6.Set(LockKey(key), token, "PX", _expire, "NX"))
		if e == nil {
			if c == "OK" {
				return token, nil
			}
		}
		// retry
		time.Sleep(30 * time.Millisecond)
	}
	return 0, RedisSetNxError
}

// TryLock will try lock a key of redis once
func TryLock(key string, expire ...int32) (int64, error) {

	// default expiry of the lock is 300 millseconds
	var _expire int32 = LockExpiry
	if len(expire) > 0 && expire[0] > 0 {
		_expire = expire[0]
	}
	token := time.Now().Unix()
	c, e := redis.String(redi6.Set(LockKey(key), token, "PX", _expire, "NX"))
	if e == nil {
		if c == "OK" {
			return token, nil
		}
	}
	return 0, RedisSetNxError
}

// Unlock will unlock a key of redis atomicly
func Unlock(key string, token int64) error {

	_, err := redi6.EvalLuaScript(1, UnlockScript, LockKey(key), token)
	return err
}

// LockKey returns a key of redis formatting Lock type
func LockKey(key string) string {
	return fmt.Sprintf(LockKeyF, key)
}
