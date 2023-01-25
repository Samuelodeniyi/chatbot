package client

import (
	"awesomeProject4/API"
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	hub *API.Hub
	conn *websocket.Conn send chan []byte
}

func (c*client) readPump()  {
	defer func() {
		c.hub.unregister <-c
		c.conn.Close()
	}()
c.conn.SetReadLimit( int64)
	
c.conn.SetReadDeadline(time.Now().
	Add(pongWait))
c.conn.SetPongHandler(func(appData string) error {
	c.conn.SetReadDeadline(time.Now().Add)
	
})	
	
}