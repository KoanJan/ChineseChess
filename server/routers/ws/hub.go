package ws

import "github.com/gorilla/websocket"

var wsHub *Hub

// Hub maintains all WebSocket connections
type Hub struct {
	conns     map[string]*Conn
	broadcast chan []byte
}

func (this *Hub) run() {
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
