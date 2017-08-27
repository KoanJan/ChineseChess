package ws

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"ChineseChess/server/logic"
	"ChineseChess/server/routers/ws/msg"
)

// 监听来自客户端的消息
func handleClient(c *Conn) {

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
			handleChatClient(message)
		case msg.Msg_Game:
			handleGameClient(message, c.UID)
		default:
			log.Println("error: unknown message type")
		}
		log.Println(data)
	}
}

// 聊天消息处理
func handleChatClient(message *msg.Msg) {

	log.Println("error: chat message is not supported yet")
}

// 游戏内消息处理
func handleGameClient(message *msg.Msg, uid string) {

	gameMsg := new(msg.GameMsg)
	if err := proto.Unmarshal(message.Body, gameMsg); err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	go logic.GameLogicFunc(gameMsg.Call)(gameMsg, uid)

}
