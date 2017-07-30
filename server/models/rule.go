package models

import "ChineseChess/server/utils"

// 棋子
const (
	PieceShuai  = iota // 帅
	PieceJiang         // 将
	PieceShiR          // 士(红)
	PieceShiB          // 士(黑)
	PieceXiangR        // 相
	PieceXiangB        // 象
	PieceMaR           // 马(红)
	PieceMaB           // 马(黑)
	PieceJuR           // 车(红)
	PieceJuB           // 车(黑)
	PiecePaoR          // 炮(红)
	PiecePaoB          // 炮(黑)
	PieceBing          // 兵
	PieceZu            // 卒
)

// 棋子走法规则验证
var rules map[int32]func(int32, int32, int32, int32, *ChessBoard) bool = map[int32]func(int32, int32, int32, int32, *ChessBoard) bool{

	PieceShuai: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if (3 <= x1) && (5 >= x1) && (3 <= x2) && (5 >= x2) && (0 <= y1) && (2 >= y1) && (0 <= y2) && (2 >= y2) {

			// 只能横向或者纵向移动
			if (x1 == x2) != (y1 == y2) {

				// 只能移动1单元距离
				if utils.SquaredEucDist(x1, y1, x2, y2) == 1 {

					return true
				}
			}
		}
		return false
	},
	PieceJiang: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if (3 <= x1) && (5 >= x1) && (3 <= x2) && (5 >= x2) && (7 <= y1) && (9 >= y1) && (7 <= y2) && (9 >= y2) {

			// 只能横向或者纵向移动
			if (x1 == x2) != (y1 == y2) {

				// 只能移动1单元距离
				if utils.SquaredEucDist(x1, y1, x2, y2) == 1 {

					return true
				}
			}
		}
		return false
	},
	PieceShiR: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if (((x1 == 3 || x1 == 5) && (y1 == 0 || y1 == 2)) || (x1 == 4 && y1 == 1)) && (((x2 == 3 || x2 == 5) && (y2 == 0 || y2 == 2)) || (x2 == 4 && y2 == 1)) {

			// 只能对角线方向移动一步
			if utils.SquaredEucDist(x1, y1, x2, y2) == 2 {

				return true
			}
		}
		return false
	},
	PieceShiB: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if (((x1 == 3 || x1 == 5) && (y1 == 7 || y1 == 9)) || (x1 == 4 && y1 == 8)) && (((x2 == 3 || x2 == 5) && (y2 == 7 || y2 == 9)) || (x2 == 4 && y2 == 8)) {

			// 只能对角线方向移动1步
			if utils.SquaredEucDist(x1, y1, x2, y2) == 2 {

				return true
			}
		}
		return false
	},
	PieceXiangR: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if ((y1 == 2 && x1%4 == 0) || ((y1 == 0 || y1 == 4) && x1%4 == 2)) && ((y2 == 2 && x2%4 == 0) || ((y2 == 0 || y2 == 4) && x2%4 == 2)) {

			// 只能对角线方向移动2步
			if utils.SquaredEucDist(x1, y1, x2, y2) == 8 {

				// 中间不能被棋子挡住
				if board.Get((x1+x2)/2, (y1+y2)/2) == -1 {

					return true
				}
			}
		}
		return false
	},
	PieceXiangB: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 验证是否在合法坐标范围内
		if ((y1 == 7 && x1%4 == 0) || ((y1 == 9 || y1 == 5) && x1%4 == 2)) && ((y2 == 7 && x2%4 == 0) || ((y2 == 9 || y2 == 5) && x2%4 == 2)) {

			// 只能对角线方向移动2步
			if utils.SquaredEucDist(x1, y1, x2, y2) == 8 {

				// 中间不能被棋子挡住
				if board.Get((x1+x2)/2, (y1+y2)/2) == -1 {

					return true
				}
			}
		}
		return false
	},
	PieceMaR: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		// 马走'日'
		if utils.SquaredEucDist(x1, y1, x2, y2) == 5 {

			// 中间不能被棋子挡住
			if (x1-x2 == -2 && board.Get(x1+1, y1) == -1) || (x1-x2 == 2 && board.Get(x1-1, y1) == -1) || (y1-y2 == -2 && board.Get(x1, y1+1) == -1) || (y1-y2 == 2 && board.Get(x1, y1-1) == -1) {

				return true
			}
		}
		return false
	},
	PieceMaB: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		// 马走'日'
		if utils.SquaredEucDist(x1, y1, x2, y2) == 5 {

			// 中间不能被棋子挡住
			if (x1-x2 == -2 && board.Get(x1+1, y1) == -1) || (x1-x2 == 2 && board.Get(x1-1, y1) == -1) || (y1-y2 == -2 && board.Get(x1, y1+1) == -1) || (y1-y2 == 2 && board.Get(x1, y1-1) == -1) {

				return true
			}
		}
		return false
	},
	PieceJuR: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		// 车可移动任意距离, 但只能横向或纵向
		if (x1 == x2) != (y1 == y2) {

			// 中间不能被棋子挡住
			if x1 == x2 {
				if y1 < y2 {
					for i := y1 + 1; i < y2; i++ {
						if board.Get(x1, i) != -1 {
							return false
						}
					}
				} else {
					for i := y1 - 1; i > y2; i-- {
						if board.Get(x1, i) != -1 {
							return false
						}
					}
				}

			} else {
				if x1 < x2 {
					for i := x1 + 1; i < x2; i++ {
						if board.Get(i, y1) != -1 {
							return false
						}
					}
				} else {
					for i := x1 - 1; i > x2; i-- {
						if board.Get(i, y1) != -1 {
							return false
						}
					}
				}
			}

			return true
		}
		return false
	},
	PieceJuB: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		// 车可移动任意距离, 但只能横向或纵向
		if (x1 == x2) != (y1 == y2) {

			// 中间不能被棋子挡住
			if x1 == x2 {
				if y1 < y2 {
					for i := y1 + 1; i < y2; i++ {
						if board.Get(x1, i) != -1 {
							return false
						}
					}
				} else {
					for i := y1 - 1; i > y2; i-- {
						if board.Get(x1, i) != -1 {
							return false
						}
					}
				}

			} else {
				if x1 < x2 {
					for i := x1 + 1; i < x2; i++ {
						if board.Get(i, y1) != -1 {
							return false
						}
					}
				} else {
					for i := x1 - 1; i > x2; i-- {
						if board.Get(i, y1) != -1 {
							return false
						}
					}
				}
			}

			return true
		}
		return false
	},
	PiecePaoR: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		if board.Get(x2, y2) == -1 {

			// 移动(任意距离, 但只能横向或纵向)
			if (x1 == x2) != (y1 == y2) {

				// 中间不能被棋子挡住
				if x1 == x2 {
					if y1 < y2 {
						for i := y1 + 1; i < y2; i++ {
							if board.Get(x1, i) != -1 {
								return false
							}
						}
					} else {
						for i := y1 - 1; i > y2; i-- {
							if board.Get(x1, i) != -1 {
								return false
							}
						}
					}

				} else {
					if x1 < x2 {
						for i := x1 + 1; i < x2; i++ {
							if board.Get(i, y1) != -1 {
								return false
							}
						}
					} else {
						for i := x1 - 1; i > x2; i-- {
							if board.Get(i, y1) != -1 {
								return false
							}
						}
					}
				}

				return true
			}
		} else {

			// 吃子(任意距离, 但只能横向或纵向,且中间必须有且仅有一子)
			if (x1 == x2) != (y1 == y2) {

				// 中间必须有且仅有一子
				if x1 == x2 {
					if y1 < y2 {
						barrier := 0
						for i := y1 + 1; i < y2; i++ {
							if board.Get(x1, i) != -1 {
								barrier++
							}
						}
						return barrier == 1
					} else {
						barrier := 0
						for i := y1 - 1; i > y2; i-- {
							if board.Get(x1, i) != -1 {
								barrier++
							}
						}
						return barrier == 1
					}

				} else {
					if x1 < x2 {
						barrier := 0
						for i := x1 + 1; i < x2; i++ {
							if board.Get(i, y1) != -1 {
								barrier++
							}
						}
						return barrier == 1
					} else {
						barrier := 0
						for i := x1 - 1; i > x2; i-- {
							if board.Get(i, y1) != -1 {
								barrier++
							}
						}
						return barrier == 1
					}
				}
			}
		}
		return false
	},
	PiecePaoB: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 活动范围不受限制
		if board.Get(x2, y2) == -1 {

			// 移动(任意距离, 但只能横向或纵向)
			if (x1 == x2) != (y1 == y2) {

				// 中间不能被棋子挡住
				if x1 == x2 {
					if y1 < y2 {
						for i := y1 + 1; i < y2; i++ {
							if board.Get(x1, i) != -1 {
								return false
							}
						}
					} else {
						for i := y1 - 1; i > y2; i-- {
							if board.Get(x1, i) != -1 {
								return false
							}
						}
					}

				} else {
					if x1 < x2 {
						for i := x1 + 1; i < x2; i++ {
							if board.Get(i, y1) != -1 {
								return false
							}
						}
					} else {
						for i := x1 - 1; i > x2; i-- {
							if board.Get(i, y1) != -1 {
								return false
							}
						}
					}
				}

				return true
			}
		} else {

			// 吃子(任意距离, 但只能横向或纵向,且中间必须有且仅有一子)
			if (x1 == x2) != (y1 == y2) {

				// 中间必须有且仅有一子
				if x1 == x2 {
					if y1 < y2 {
						barrier := 0
						for i := y1 + 1; i < y2; i++ {
							if board.Get(x1, i) != -1 {
								barrier++
							}
						}
						return barrier == 1
					} else {
						barrier := 0
						for i := y1 - 1; i > y2; i-- {
							if board.Get(x1, i) != -1 {
								barrier++
							}
						}
						return barrier == 1
					}

				} else {
					if x1 < x2 {
						barrier := 0
						for i := x1 + 1; i < x2; i++ {
							if board.Get(i, y1) != -1 {
								barrier++
							}
						}
						return barrier == 1
					} else {
						barrier := 0
						for i := x1 - 1; i > x2; i-- {
							if board.Get(i, y1) != -1 {
								barrier++
							}
						}
						return barrier == 1
					}
				}
			}
		}
		return false
	},
	PieceBing: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 兵过河之前只能往前，过河之后可以往前或者平移, 不可后退
		if y1 == 3 || y1 == 4 {

			// 过河之前
			return x1 == x2 && y2-y1 == 1
		} else if y1 > 4 {

			// 过河之后
			return (x1 == x2 && y2-y1 == 1) || ((x1-x2 == -1 || x1-x2 == 1) && y1 == y2)
		}
		return false
	},
	PieceZu: func(x1, y1, x2, y2 int32, board *ChessBoard) bool {

		if isBlockedBySameColor(x1, y1, x2, y2, board) {
			return false
		}

		// 卒过河之前只能往前，过河之后可以往前或者平移, 不可后退
		if y1 == 6 || y1 == 5 {

			// 过河之前
			return x1 == x2 && y2-y1 == -1
		} else if y1 < 5 {

			// 过河之后
			return (x1 == x2 && y2-y1 == -1) || ((x1-x2 == -1 || x1-x2 == 1) && y1 == y2)
		}
		return false
	},
}

// 检查目标坐标是否被己方棋子占据
func isBlockedBySameColor(x1, y1, x2, y2 int32, board *ChessBoard) bool {

	from := board.Get(x1, y1)
	to := board.Get(x2, y2)
	return to == -1 || (to%2 != from%2)
}

/*
验证是否被游戏规则允许
*/
func AllowedUnderRules(x1, y1, x2, y2 int32, board *ChessBoard) bool {

	return rules[board.Get(x1, y1)](x1, y1, x2, y2, board)
}
