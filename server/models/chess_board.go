package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

const (
	ChessBoardMaxX = 8
	ChessBoardMaxY = 9
)

// 棋子
const (
	PieceNo     = -1   // 没有棋子占位
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

// 棋子
const (
	_                       = iota
	ChessBoardStatusReady   // 准备
	ChessBoardStatusPlaying // 比赛中
	ChessBoardStatusEnded   // 结束
)

/*
棋盘
*/
type ChessBoard struct {
	Common
	Steps       []Step         `bson:"steps",json:"steps"`                        // 走子历史
	RedUserID   bson.ObjectId  `bson:"red_user_id",json:"red_user_id,string"`     // 红方用户ID
	BlackUserID bson.ObjectId  `bson:"black_user_id",json:"black_user_id,string"` // 黑方用户ID
	WinnerID    *bson.ObjectId `bson:"winner_id",json:"winner_id,string"`         // 获胜方用户ID(如果是和局则该局无值)
	Status      int32          `bson:"status",json:"status"`                      // 棋局状态

	Others []string `bson:"-",json:"-"` // 观战者

	board [ChessBoardMaxX + 1][ChessBoardMaxY + 1]int32 // 棋盘

}

func (this *ChessBoard) CollectionName() string {
	return "chess_board"
}

/*
走子记录
*/
type Step [4]int32 // [x1, y1, x2, y2]

/*
获取指定坐标上的信息
*/
func (this *ChessBoard) Get(x, y int32) int32 {
	if validLocation(x, y) {
		return this.board[x][y]
	}
	return -1
}

/*
设定指定坐标上的值
*/
func (this *ChessBoard) Set(x, y, v int32) error {
	if validLocation(x, y) {
		this.board[x][y] = v
		return nil
	}
	return errors.New("坐标不合法")
}

// 检验坐标是否合法
func validLocation(x, y int32) bool {

	return (x <= ChessBoardMaxX) && (x >= 0) && (y <= ChessBoardMaxX) && (y >= 0)
}

/*
 开启新的棋局
*/
func NewChessBoard(redUserID, blackUserID string) *ChessBoard {

	chessBoard := new(ChessBoard)

	board := [ChessBoardMaxX + 1][ChessBoardMaxY + 1]int32{}
	// 初始化红方
	board[0][0], board[8][0] = PieceJuR, PieceJuR
	board[1][0], board[7][0] = PieceMaR, PieceMaR
	board[2][0], board[6][0] = PieceXiangR, PieceXiangR
	board[3][0], board[5][0] = PieceShiR, PieceShiR
	board[4][0] = PieceShuai
	board[1][2], board[7][2] = PiecePaoR, PiecePaoR
	board[0][3], board[2][3], board[4][3], board[6][3], board[8][3] = PieceBing, PieceBing, PieceBing, PieceBing, PieceBing
	// 初始化黑方
	board[0][9], board[8][9] = PieceJuB, PieceJuB
	board[1][9], board[7][9] = PieceMaB, PieceMaB
	board[2][9], board[6][9] = PieceXiangB, PieceXiangB
	board[3][9], board[5][9] = PieceShiB, PieceShiB
	board[4][9] = PieceJiang
	board[1][7], board[7][7] = PiecePaoB, PiecePaoB
	board[0][6], board[2][6], board[4][6], board[6][6], board[8][6] = PieceZu, PieceZu, PieceZu, PieceZu, PieceZu

	chessBoard.ID = bson.NewObjectId()
	chessBoard.board = board
	chessBoard.Steps = []Step{}
	chessBoard.RedUserID = bson.ObjectIdHex(redUserID)
	chessBoard.BlackUserID = bson.ObjectIdHex(blackUserID)
	chessBoard.Status = ChessBoardStatusReady

	return chessBoard
}
