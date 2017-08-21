package matchingQueue

import (
	"log"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
)

var queue chan string

func Enqueue(userID string) {

	queue <- userID
}

func handleQueue() {
	for {
		a := <-queue
		b := <-queue
		go func() {
			// TODO 通知棋手匹配失败
			board := models.NewChessBoard(a, b)
			if err := daf.Insert(board); err != nil {
				log.Printf("error: %s和%s匹配失败\n", a, b)
				return
			}
			if err := cache.AddBoardCache(board); err != nil {
				log.Printf("error: %s和%s匹配失败\n", a, b)
				return
			}
		}()
	}
}

func init() {

	queue = make(chan string, 1024)
}
