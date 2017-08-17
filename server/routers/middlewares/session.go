package middlewares

import (
	"encoding/json"

	"ChineseChess/server/models"
)

// Session
type Session struct {
	UserID string `json:"user_id"` // 用户id
	Nick   string `json:"nick"`    // 用户昵称
}

func GenerateSessionString(user *models.User) string {

	session := new(Session)
	session.UserID = user.ID.Hex()
	session.Nick = user.Nick
	sessionString, _ := json.Marshal(session)
	return string(sessionString)
}
