package cache

import (
	"github.com/garyburd/redigo/redis"

	"ChineseChess/server/models/cache/utils"
	"ChineseChess/server/redi6"
)

// 用户状态
type SessionStatus = string

const (
	SessionStatusOK   SessionStatus = "ok"   // 空闲
	SessionStatusGame SessionStatus = "game" // 游戏中
)

const SessionExpiry int32 = 7 * 24 * 3600 // 登陆默认保存7天

// Session
type Session struct {
	UserID string        `redis:"user_id"` // 用户id
	Nick   string        `redis:"nick"`    // 用户昵称
	Status SessionStatus `redis:"status"`  // 用户状态
}

// Save a session into redis
func (this *Session) Save() error {

	token, err := utils.Lock(this.UserID)
	defer utils.Unlock(this.UserID, token)
	if err != nil {
		return err
	}
	_, err = redi6.Hmset(utils.SessionKey(this.UserID), utils.SessionFieldNick, this.Nick, utils.SessionFieldStatus, this.Status)
	if err != nil {
		return err
	}
	redi6.Expire(utils.SessionKey(this.UserID), SessionExpiry)
	return nil
}

// UpdateSession update a session with it's one field
func UpdateSession(userID string, field utils.SessionField, value interface{}) error {

	token, err := utils.Lock(userID)
	defer utils.Unlock(userID, token)
	if err != nil {
		return err
	}
	_, err = redi6.Hset(userID, field, value)
	return err
}

// DelSession delete a session from redis
func DelSession(userID string) error {

	_, err := redi6.Del(userID)
	return err
}

// FindSession finds a session by userID from redis
func FindSession(userID string) (*Session, error) {

	src, err := redis.Values(redi6.Hgetall(userID))
	if err != nil {
		return nil, err
	}
	session := new(Session)
	if err = redis.ScanStruct(src, session); err != nil {
		return nil, err
	}
	return session, nil
}

// NewSession returns a new session
func NewSession(userID, nick string, status SessionStatus) *Session {

	return &Session{userID, nick, status}
}
