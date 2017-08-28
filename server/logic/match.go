package logic

import (
	"encoding/json"
	"errors"
	"log"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
	"ChineseChess/server/routers/ws/msg"
)

var (
	matchingQueue chan string                  = make(chan string, 1024)            // 匹配队列
	matchResults  map[string]chan *matchResult = make(map[string]chan *matchResult) // 匹配结果
)

// 匹配结果
type matchResult struct {
	board *models.ChessBoard
	err   error
}

// 监听匹配队列
func handleQueue() {

	for {
		a := <-matchingQueue
		b := <-matchingQueue
		go func() {

			// 通知棋手匹配结果
			board := models.NewChessBoard(a, b)
			if err := daf.Insert(board); err != nil {
				log.Printf("error: %s和%s匹配失败\n", a, b)
				matchResults[a] <- &matchResult{nil, errors.New("匹配失败")}
				matchResults[b] <- &matchResult{nil, errors.New("匹配失败")}
				return
			}
			if err := cache.AddBoardCache(board); err != nil {
				log.Printf("error: %s和%s匹配失败\n", a, b)
				matchResults[a] <- &matchResult{nil, errors.New("匹配失败")}
				matchResults[b] <- &matchResult{nil, errors.New("匹配失败")}
				return
			}
			matchResults[a] <- &matchResult{board, errors.New("匹配失败")}
			matchResults[b] <- &matchResult{board, errors.New("匹配失败")}
		}()
	}
}

// 匹配
func Match(gameMsg *msg.GameMsg, uid string) {

	matchResults[uid] = make(chan *matchResult, 1)
	matchingQueue <- uid
	r := <-matchResults[uid]
	close(matchResults[uid])
	delete(matchResults, uid)
	var (
		data []byte = []byte{}
		err  error  = nil
	)

	if r.err == nil {
		data, err = json.Marshal(r.board)
	}
	// send game server msg
	PushGameServerMsg(gameMsg.Call, data, err, uid)
}

// 初始化
func initMatch() {
	go handleQueue()
}
