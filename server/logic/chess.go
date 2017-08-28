package logic

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
	"ChineseChess/server/routers/ws/msg"
)

// PlayForm contains the all parameters that the Play API needs
type PlayForm struct {
	X1      int32  `json:"x1"`
	Y1      int32  `json:"y1"`
	X2      int32  `json:"x2"`
	Y2      int32  `json:"y2"`
	Step    int    `json:"step"` // 步数验证
	BoardID string `json:"board_id"`
}

// PlayResp
type PlayResp struct {
	X1   int32 `json:"x1"`
	Y1   int32 `json:"y1"`
	X2   int32 `json:"x2"`
	Y2   int32 `json:"y2"`
	Step int   `json:"step"` // 步数验证
}

// Play 下子
func Play(gameMsg *msg.GameMsg, uid string) {

	form := new(PlayForm)
	if err := json.Unmarshal(gameMsg.Data, form); err != nil {

		PushGameServerMsg(GameLogicFuncPlay, []byte{}, errors.New("数据解析失败!"), uid)
		return
	}

	// 验证是否符合规则
	if err := cache.UpdateBoardCache(form.BoardID, func(board *models.ChessBoard) error {

		if board.WinnerID != nil {
			return errors.New("比赛已经结束")
		}

		if form.Step != len(board.Steps)+1 {
			return errors.New("你穿越了")
		}

		if !AllowedUnderRules(form.X1, form.Y1, form.X2, form.Y2, board, uid) {
			return errors.New("不符合游戏规则")
		}

		// 保存执行前的参数
		var cacheV1, cacheV2 int32 = board.Get(form.X1, form.Y1), board.Get(form.X2, form.Y2)

		// 执行
		board.Set(form.X2, form.Y2, board.Get(form.X1, form.Y1))
		board.Set(form.X1, form.Y1, models.PieceNo)

		if cacheV2 == models.PieceJiang || cacheV2 == models.PieceShuai {

			// 吃将
			winnerID := bson.ObjectIdHex(uid)
			board.WinnerID = &winnerID
		} else if IsInDanger(board, uid) {

			// 如果走完这步会被吃将,则撤销本次操作
			board.Set(form.X1, form.Y1, cacheV1)
			board.Set(form.X2, form.Y2, cacheV2)
			return errors.New("正在被将军!")
		}

		// 更新下子记录
		board.Steps = append(board.Steps, models.Step{form.X1, form.Y1, form.X2, form.Y2})

		// 更新到数据库
		if err := daf.Update(board); err != nil {
			fmt.Errorf("更新到数据库失败: %v\n", err)
		}

		return nil
	}); err != nil {

		PushGameServerMsg(GameLogicFuncPlay, []byte{}, err, uid)
		return
	} else {
		resp := &PlayResp{form.X1, form.Y1, form.X2, form.Y2, form.Step}
		respData, e := json.Marshal(resp)
		if e != nil {
			PushGameServerMsg(GameLogicFuncPlay, []byte{}, errors.New("操作成功, 同步棋局失败"), uid)
			return
		}
		PushGameServerMsg(GameLogicFuncPlay, respData, nil, uid)
		return
	}
}

// 是否被将军
func IsInDanger(board *models.ChessBoard, userID string) bool {

	switch userID {
	case board.RedUserID.Hex():
		// 红方
		var (
			x, y  int32
			found bool = false
		)
		for i := int32(3); i <= 5; i++ {
			for j := int32(0); j <= 2; j++ {
				if board.Get(i, j) == models.PieceShuai {
					x, y, found = i, j, true
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false // 你棋子掉了
		}
		// 车炮卒四个方向
		// 右
		blocks := 0
		for i := x + 1; i <= models.ChessBoardMaxX && blocks <= 1; i++ {
			v := board.Get(i, y)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuB {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceZu && i == x+1 {
				// 卒
				return true
			} else if blocks == 1 && v == models.PiecePaoB {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 左
		blocks = 0
		for i := x - 1; i >= 0 && blocks <= 1; i-- {
			v := board.Get(i, y)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuB {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceZu && i == x-1 {
				// 卒
				return true
			} else if blocks == 1 && v == models.PiecePaoB {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 上
		blocks = 0
		for j := y + 1; j <= models.ChessBoardMaxY && blocks <= 1; j++ {
			v := board.Get(x, j)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuB {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceZu && j == y+1 {
				// 卒
				return true
			} else if blocks == 1 && v == models.PiecePaoB {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 下
		blocks = 0
		for j := y - 1; j >= 0 && blocks <= 1; j-- {
			v := board.Get(x, j)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuB {
				// 车
				return true
				// 卒不能往后走, 所以这个方向不考虑
			} else if blocks == 1 && v == models.PiecePaoB {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}

		// 马(最多可以有8个方位, 4个马脚)
		if board.Get(x+1, y+1) == models.PieceNo {
			// 右上角
			if board.Get(x+1, y+2) == models.PieceMaB || board.Get(x+2, y+1) == models.PieceMaB {
				return true
			}
		} else if board.Get(x-1, y+1) == models.PieceNo {
			// 左上角
			if board.Get(x-2, y+1) == models.PieceMaB || board.Get(x-1, y+2) == models.PieceMaB {
				return true
			}
		} else if y > 0 {
			// 如果有左下角和右下角
			if board.Get(x-1, y-1) == models.PieceNo {
				// 左下角
				if board.Get(x-2, y-1) == models.PieceMaB || (y > 1 && board.Get(x-1, y-2) == models.PieceMaB) {
					return true
				}
			} else if board.Get(x+1, y-1) == models.PieceNo {
				// 右下角
				if board.Get(x+2, y-1) == models.PieceMaB || (y > 1 && board.Get(x+1, y-2) == models.PieceMaB) {
					return true
				}
			}
		}
		return false
	case board.BlackUserID.Hex():
		// 黑方
		var (
			x, y  int32
			found bool = false
		)
		for i := int32(3); i <= 5; i++ {
			for j := int32(7); j <= 9; j++ {
				if board.Get(i, j) == models.PieceJiang {
					x, y, found = i, j, true
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false // 你棋子掉了
		}
		// 车炮兵四个方向
		// 右
		blocks := 0
		for i := x + 1; i <= models.ChessBoardMaxX && blocks <= 1; i++ {
			v := board.Get(i, y)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuR {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceBing && i == x+1 {
				// 兵
				return true
			} else if blocks == 1 && v == models.PiecePaoR {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 左
		blocks = 0
		for i := x - 1; i >= 0 && blocks <= 1; i-- {
			v := board.Get(i, y)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuR {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceBing && i == x-1 {
				// 兵
				return true
			} else if blocks == 1 && v == models.PiecePaoR {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 上
		blocks = 0
		for j := y + 1; j <= models.ChessBoardMaxY && blocks <= 1; j++ {
			v := board.Get(x, j)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuR {
				// 车
				return true
				// 兵不能往后走, 所以这个方向不考虑
			} else if blocks == 1 && v == models.PiecePaoR {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}
		// 下
		blocks = 0
		for j := y - 1; j >= 0 && blocks <= 1; j-- {
			v := board.Get(x, j)
			if v == models.PieceNo {
				continue
			}
			if blocks == 0 && v == models.PieceJuR {
				// 车
				return true
			} else if blocks == 0 && v == models.PieceBing && j == y+1 {
				// 卒
				return true
			} else if blocks == 1 && v == models.PiecePaoR {
				// 炮
				return true
			} else {
				blocks += 1
			}
		}

		// 马(最多可以有8个方位, 4个马脚)
		if board.Get(x-1, y-1) == models.PieceNo {
			// 左下角
			if board.Get(x-1, y-2) == models.PieceMaR || board.Get(x-2, y-1) == models.PieceMaR {
				return true
			}
		} else if board.Get(x+1, y-1) == models.PieceNo {
			// 右下角
			if board.Get(x+2, y-1) == models.PieceMaR || board.Get(x+1, y-2) == models.PieceMaR {
				return true
			}
		} else if y < 9 {
			// 如果有左上角和右上角
			if board.Get(x-1, y+1) == models.PieceNo {
				// 左上角
				if board.Get(x-2, y+1) == models.PieceMaR || (y < 8 && board.Get(x-1, y+2) == models.PieceMaR) {
					return true
				}
			} else if board.Get(x+1, y+1) == models.PieceNo {
				// 右上角
				if board.Get(x+2, y+1) == models.PieceMaR || (y < 8 && board.Get(x+1, y+2) == models.PieceMaR) {
					return true
				}
			}
		}
		return false
	default:
		return false // 又不是你下棋别bb那么多
	}

}

// 初始化
func initChess() {

}
