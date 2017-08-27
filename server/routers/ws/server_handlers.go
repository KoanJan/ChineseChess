package ws

import (
	"log"

	"github.com/golang/protobuf/proto"

	"ChineseChess/server/logic"
	"ChineseChess/server/routers/ws/msg"
)

// 监听来自服务端的消息
func handleServer(h *Hub) {

	go handleGameServer(h)
}

// 服务端发送的游戏消息
func handleGameServer(h *Hub) {

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
			h.Broadcast(data)
			return
		}
		for _, uid := range gsm.UIDs {

			h.Push(uid, data)
		}
	})
}
