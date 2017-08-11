package routers

import "ChineseChess/server/routers/v1"

var routerConf = []handler{

	// hello
	{"hello", hello},

	{"CreateChessBoard", v1.CreateChessBoard}, // 创建棋局
}

//hello
func hello(param []byte) []byte {
	return []byte{'f', 'o', 'f', 'f'}
}
