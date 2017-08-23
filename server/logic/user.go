package logic

import (
	"ChineseChess/server/logger"
	"ChineseChess/server/redis"
)

// 用户是否在线
func UserIsOnline(userID string) bool {

	reply, err := redis.Get(userID)
	if err != nil {
		logger.Error(err)
		return false
	}
	return reply != nil
}
