package v1

import (
	"encoding/json"

	"ChineseChess/server/daf"
	"ChineseChess/server/models"
	. "ChineseChess/server/routers/common"
	"fmt"
)

type chessBoardForm struct {
	RedUserID   string `json:"red_user_id"`
	BlackUserID string `json:"black_user_id"`
}

// 创建棋局
func CreateChessBoard(data []byte) []byte {

	form := new(chessBoardForm)
	if err := json.Unmarshal(data, form); err != nil {
		return RespErr(err)
	}
	fmt.Println(form)

	// 增加棋局
	board := models.NewChessBoard(form.RedUserID, form.BlackUserID)
	if err := daf.Insert(board); err != nil {
		return RespErr(err)
	}
	return RespOK()
}
