package chat

import "github.com/gorilla/websocket"

type BloadcastManager struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func New() *BloadcastManager {
	return &BloadcastManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (b *BloadcastManager) Run() {
	for {
		select {
		case client := <-b.register:
			b.clients[client] = true

		case client := <-b.unregister:
			close(client.message)
			delete(b.clients, client)
		case message := <-b.broadcast:
			clients := b.clients
			for client := range clients {
				select {
				case client.message <- message.Data:
				default:
					close(client.message)
					delete(b.clients, client)
				}
			}
		}
	}
}

func (b *BloadcastManager) RegisterNewClient(conn *websocket.Conn) {
	client := &Client{
		manager: b,
		conn:    conn,
		message: make(chan []byte, 256),
	}
	client.manager.register <- client

	go client.reader()
	go client.writer()
}
