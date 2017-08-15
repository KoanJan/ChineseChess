package logic

import "ChineseChess/server/models"

// 是否被将军
func IsInDanger(board *models.ChessBoard, userID string) bool {

	switch userID {
	// 红方
	case board.RedUserID.Hex():
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
		// TODO 车炮卒四个方向
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

		// TODO 马
	// 黑方
	case board.BlackUserID.Hex():
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
		// TODO 判断车,马,炮.兵
	default:
		return false // 又不是你下棋别bb那么多
	}

}
