package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"time"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Message

	// Registered user for this connection
	user User
}

func (c *connection) reader() {
	for {
		var m Message
		err := websocket.JSON.Receive(c.ws, &m)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Received:", m)
		m.Sender = c.user
		m.Time = time.Now()
		if len(m.Text) > 140 {
			m.Text = m.Text[:140] + "…"
		}
		h.broadcast <- m
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.JSON.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
	c := &connection{send: make(chan Message, 256), ws: ws, user: User{"axel", "avatar.jpg"}}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
