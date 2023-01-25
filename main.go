package main

import
{"fmt"
"net/http"
"os"
"github.com/gorilla/mux"
}
type Client struct {
	h* Hub
	j * websocket.conn

}
type clientmesage struct {
	client *Client
	message []byte

}
type Hub struct {
	chatbot bot .Bot
	clients map[*Client]bool
	broadcastmsg chan *clientmesage
	register chan * Client
	unregister chan * Client


}