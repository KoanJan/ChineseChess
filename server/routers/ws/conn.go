package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"ChineseChess/server/routers/ws/api/game"
	"ChineseChess/server/routers/ws/msg"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Conn is a WebSocket connection
type Conn struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Conn) readPump() {
	defer c.conn.Close()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, data, err := c.conn.ReadMessage()
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
			log.Println("error: chat message is not supported yet")
		case msg.Msg_Game:
			resp := new(msg.RespMsg)
			respData, err := game.Dispatch(message.Body)
			if err != nil {
				resp.Err = err.Error()
				resp.Code = msg.RespMsg_Failed
			} else {
				resp.Data = respData
				resp.Code = msg.RespMsg_Done
			}
			b, _ := proto.Marshal(resp)
			c.conn.WriteMessage(websocket.TextMessage, b)
		default:
			log.Println("error: unknown message type")
		}
		log.Println(data)
	}
}

func (c *Conn) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() { ticker.Stop(); c.conn.Close() }()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			// write all
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWS(uid string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &Conn{wsHub, conn, make(chan []byte, 256)}
	c.hub.conns[uid] = c

	go c.readPump()
	go c.writePump()
}
