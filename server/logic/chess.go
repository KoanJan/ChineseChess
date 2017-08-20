package logic

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
)

// Play 下子
func Play(x1, y1, x2, y2 int32, step int, boardID, userID string) ([]byte, error) {

	// 验证是否符合规则
	if err := cache.UpdateBoardCache(boardID, func(board *models.ChessBoard) error {

		if board.WinnerID != nil {
			return errors.New("比赛已经结束")
		}

		if step != len(board.Steps)+1 {
			return errors.New("你穿越了")
		}

		if !AllowedUnderRules(x1, y1, x2, y2, board, userID) {
			return errors.New("不符合游戏规则")
		}

		// 保存执行前的参数
		var cacheV1, cacheV2 int32 = board.Get(x1, y1), board.Get(x2, y2)

		// 执行
		board.Set(x2, y2, board.Get(x1, y1))
		board.Set(x1, y1, models.PieceNo)

		if cacheV2 == models.PieceJiang || cacheV2 == models.PieceShuai {

			// 吃将
			winnerID := bson.ObjectIdHex(userID)
			board.WinnerID = &winnerID
		} else if IsInDanger(board, userID) {

			// 如果走完这步会被吃将,则撤销本次操作
			board.Set(x1, y1, cacheV1)
			board.Set(x2, y2, cacheV2)
			return errors.New("正在被将军!")
		}

		// 更新下子记录
		board.Steps = append(board.Steps, models.Step{x1, y1, x2, y2})

		// 更新到数据库
		if err := daf.Update(board); err != nil {
			fmt.Errorf("更新到数据库失败: %v\n", err)
		}

		return nil
	}); err != nil {
		return []byte{}, err
	} else {
		resp := &PlayResp{x1, x2, y1, y2, step}
		return json.Marshal(resp)
	}
}
