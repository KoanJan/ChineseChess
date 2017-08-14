package logic

import (
	"errors"
	"fmt"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
)

// Play 下子
func Play(x1, y1, x2, y2 int32, boardID, userID string) error {

	// 验证是否符合规则
	return cache.UpdateBoardCache(boardID, func(board *models.ChessBoard) error {

		if !AllowedUnderRules(x1, y1, x2, y2, board, userID) {
			return errors.New("不符合游戏规则")
		}
		board.Set(x2, y2, board.Get(x1, y1))
		board.Set(x1, y1, models.PieceNo)
		if err := daf.Update(board); err != nil {
			fmt.Errorf("更新到数据库失败: %v\n", err)
		}
		return nil
	})
}
