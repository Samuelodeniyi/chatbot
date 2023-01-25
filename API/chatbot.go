package API

import "github.com/gorilla/websocket"

type clientmesage {
client *Client`josn:"client"`
message []byte`json:"message"`
}

type Hub struct {
	chatbot bot .Bot`json:"chatbot"`
	clients map[*Client]bool`json:"clients"`
	broadcastmsg chan *clientmesage`json:"broadcastmsg"`
	register chan * Client`json:"register"`
	unregister chan * Client`json:"unregister"`



}

type Client struct {
	h *Hub `json:"h"`
	j *websocket.conntype`json:"j"`
}

func NewHub(chatbot bot .Bot)* Hub  {
	return &Hub{
		chatbot: chatbot,
		broadcastmsg: make(chan
		*clientmesage),register:
			make(chan * Client),
		unregister: make(chan*Client),
		clients:
			make(map[*Client]bool),
	}


}
func (h*Hub) chatBot()bot.Bot {
	return h.chatbot

}
func (h*Hub) sendMessage(client*Client,message[]byte) {
	client.send <-message

}

	/*type clientmesage struct {
	client *Client
	message []byte

	}

	/*type Hub struct {
	chatbot bot .Bot
	clients map[*Client]bool
	broadcastmsg chan *clientmesage
	register chan * Client
	unregister chan * Client


	}*/
