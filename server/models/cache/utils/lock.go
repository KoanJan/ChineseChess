package utils

import (
	"errors"
	"fmt"
	"time"

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
	LockReqeat int = 5

	// default expiry of lock
	LockExpiry int32 = 3000
)

// Lock will repeat locking a key of redis, and return a lock token and error
func Lock(key string, expire ...int32) (int64, error) {

	// default expiry of the lock is 3000 millseconds
	var _expire int32 = LockExpiry
	if len(expire) > 0 && expire[0] > 0 {
		_expire = expire[0]
	}
	for i := 0; i < LockReqeat; i++ {
		token := time.Now().Unix()
		if e := redi6.Setnxpx(LockKey(key), token, _expire); e == nil {
			return token, e
		}
		// 每秒重试一次
		time.Sleep(time.Second)
	}
	return 0, errors.New("请求失败")
}

// TryLock will try lock a key of redis once
func TryLock(key string, expire ...int32) (int64, error) {

	// default expiry of the lock is 3000 millseconds
	var _expire int32 = LockExpiry
	if len(expire) > 0 && expire[0] > 0 {
		_expire = expire[0]
	}
	token := time.Now().Unix()
	return token, redi6.Setnxpx(LockKey(key), token, _expire)
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
