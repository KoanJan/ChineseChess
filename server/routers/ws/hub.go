package ws

import "github.com/gorilla/websocket"

var wsHub *Hub

// Hub maintains all WebSocket connections
type Hub struct {
	broadcast chan []byte
	conns     map[string]*Conn
}

func (this *Hub) run() {

	handleServer(this) // 监听来自服务端的消息

	for {
		select {
		case msg := <-this.broadcast:
			for uid, c := range this.conns {
				if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					close(c.wch)
					delete(this.conns, uid)
				}
			}
		}
	}
}

// Push data to one connection
func (this *Hub) Push(uid string, data []byte) {

	if c, ol := this.conns[uid]; ol {
		c.Write(data)
	}
}

// Broadcast data
func (this *Hub) Broadcast(data []byte) {

	this.broadcast <- data
}

func newHub() *Hub {

	hub := new(Hub)
	hub.conns = make(map[string]*Conn)
	hub.broadcast = make(chan []byte)
	return hub
}

func init() {
	wsHub = newHub()
	go wsHub.run()
}
