package logic

import (
	"ChineseChess/server/logger"
	"ChineseChess/server/redi6"
)

// 用户是否在线
func UserIsOnline(userID string) bool {

	reply, err := redi6.Get(userID)
	if err != nil {
		logger.Error(err)
		return false
	}
	return reply != nil
}
