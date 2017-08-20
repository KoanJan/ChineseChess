package ws

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"ChineseChess/server/cache"
	"ChineseChess/server/routers/ws/api/game"
	"ChineseChess/server/routers/ws/msg"
)

// 消息处理
func handle(c *Conn) {

	for {
		data, err := c.Read()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		// switch msg type
		message := new(msg.Msg)
		if err = proto.Unmarshal(data, message); err != nil {
			log.Printf("error: %v", err)
			continue
		}
		switch message.Type {
		case msg.Msg_Chat:
			handleChat(message)
		case msg.Msg_Game:
			handleGame(message)
		default:
			log.Println("error: unknown message type")
		}
		log.Println(data)
	}
}

// 聊天消息处理
func handleChat(message *msg.Msg) {

	log.Println("error: chat message is not supported yet")
}

// 游戏内消息处理
func handleGame(message *msg.Msg) {

	gameMsg := new(msg.GameMsg)
	if err := proto.Unmarshal(message.Body, gameMsg); err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	resp := new(msg.RespMsg)
	respData, err := game.Dispatch(gameMsg)
	if err != nil {
		resp.Err = err.Error()
		resp.Code = msg.RespMsg_Failed
	} else {
		resp.Data = respData
		resp.Code = msg.RespMsg_Done
	}
	b, _ := proto.Marshal(resp)

	// 获取棋盘状态
	snapshot, err := cache.SnapshotBoard(gameMsg.BoardID)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}

	// 转发给所有观看者或者棋手
	for _, userID := range snapshot.Others {
		Push(userID, b)
	}
}
