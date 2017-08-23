package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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
	UID string

	hub  *Hub
	conn *websocket.Conn
	wch  chan []byte
}

// Write some data into a connection
func (c *Conn) Write(data []byte) {
	c.wch <- data
}

// Read data from a connection
func (c *Conn) Read() ([]byte, error) {
	_, data, err := c.conn.ReadMessage()
	return data, err
}

func (c *Conn) readPump() {
	defer c.conn.Close()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	handle(c)
}

func (c *Conn) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() { ticker.Stop(); c.conn.Close() }()

	for {
		select {
		case msg, ok := <-c.wch:
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
			n := len(c.wch)
			for i := 0; i < n; i++ {
				w.Write(<-c.wch)
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
	c := &Conn{uid, wsHub, conn, make(chan []byte, 256)}
	c.hub.conns[uid] = c

	go c.readPump()
	go c.writePump()
}

// Push data to one connection
func Push(uid string, data []byte) {

	if c, ol := wsHub.conns[uid]; ol {
		c.Write(data)
	}
}

// Broadcast data
func Broadcast(data []byte) {

	for uid, _ := range wsHub.conns {
		Push(uid, data)
	}
}
