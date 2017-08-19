package game

import (
	"github.com/golang/protobuf/proto"

	"ChineseChess/server/logic"
	"ChineseChess/server/routers/ws/msg"
)

func Dispatch(data []byte) ([]byte, error) {

	gameMsg := new(msg.GameMsg)
	if err := proto.Unmarshal(data, gameMsg); err != nil {
		return nil, err
	}

	return logic.GameLogicFunc(gameMsg.Call)(gameMsg.Data)
}
