package v1

import (
	"encoding/json"

	"ChineseChess/server/logic"
	"ChineseChess/server/models"
	. "ChineseChess/server/routers/common"
)

type stepForm struct {
	ChessBoardID string      `json:"chess_board_id"` // 棋局id
	Step         models.Step `json:"step"`           // 具体走法
	UserID       string      `json:"user_id"`        // 用户id
}

// 走子
func CreateStep(data []byte) []byte {

	// 解析数据
	form := new(stepForm)
	if err := json.Unmarshal(data, form); err != nil {
		return RespErr(err)
	}

	// 执行
	if err := logic.Play(form.Step[0], form.Step[1], form.Step[2], form.Step[3], form.ChessBoardID, form.UserID); err != nil {
		return RespErr(err)
	}

	return RespOK()
}
