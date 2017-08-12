package v1

import (
	"encoding/json"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
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

	if err := cache.UpdateBoardCache(form.ChessBoardID, func(board *models.ChessBoard) error {

		// 走子
		if err := board.Go(form.Step); err != nil {
			return err
		}

		// 更新数据
		if err := daf.Update(board); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return RespErr(err)
	}

	// TODO 判断输赢

	return RespOK()
}
