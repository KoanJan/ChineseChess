package game

import (
	"ChineseChess/server/logic"
	"ChineseChess/server/routers/ws/msg"
)

// Dispatch can dispatch all the request of game
func Dispatch(gameMsg *msg.GameMsg) ([]byte, error) {

	return logic.GameLogicFunc(gameMsg.Call)(gameMsg.Data)
}
