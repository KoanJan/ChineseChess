package ws

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"ChineseChess/server/logic"
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
			handleGame(message, c.UID)
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
func handleGame(message *msg.Msg, uid string) {

	gameMsg := new(msg.GameMsg)
	if err := proto.Unmarshal(message.Body, gameMsg); err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	go logic.GameLogicFunc(gameMsg.Call)(gameMsg, uid)

}

// 服务端发送的游戏消息
func handleServer() {

	logic.HandleGameServerMsg(func(gsm *msg.GameServerMsg) {

		body, err := proto.Marshal(gsm.GameMsg)
		if err != nil {
			log.Println(err)
			return
		}
		m := new(msg.Msg)
		m.Type = msg.Msg_Game
		m.Body = body
		data, err := proto.Marshal(m)
		if err != nil {
			log.Println(err)
			return
		}
		if len(gsm.UIDs) == 0 {

			// broadcast
			Broadcast(data)
			return
		}
		for _, uid := range gsm.UIDs {

			Push(uid, data)
		}
	})
}
