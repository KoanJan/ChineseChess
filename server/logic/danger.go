package logic

import "ChineseChess/server/models"

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
