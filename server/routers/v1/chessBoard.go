package v1

import (
	"encoding/json"

	"ChineseChess/server/models"
	. "ChineseChess/server/routers/common"
	"gopkg.in/mgo.v2/bson"
)

func CreateChessBoard(data []byte) []byte {

	board := new(models.ChessBoard)
	if err := json.Unmarshal(data, board); err != nil {
		return RespErr(err)
	}

	// 增加棋局
	board.ID = bson.NewObjectId()
	if err := board.Save(); err != nil {
		return RespErr(err)
	}
	return RespOK()
}
