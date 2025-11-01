package ws

import "encoding/json"

type Hub struct {
	clients    map[*Client]bool
	rooms      map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	router     *MessageRouter
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[string]map[*Client]bool),
		router:     NewMessageRouter(),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				// Remove from all rooms
				for roomID, roomClients := range h.rooms {
					delete(roomClients, client)
					if len(roomClients) == 0 {
						delete(h.rooms, roomID)
					}
				}
				close(client.send)
			}

		case message := <-h.broadcast:
			// Broadcast to specific room
			if roomClients, exists := h.rooms[message.RoomID]; exists {
				for client := range roomClients {
					select {
					case client.send <- h.messageToBytes(message):
					default:
						// Failed to send, remove client
						delete(roomClients, client)
						if len(roomClients) == 0 {
							delete(h.rooms, message.RoomID)
						}
					}
				}
			}
		}
	}
}

func (h *Hub) JoinRoom(client *Client, roomID string) {
	if h.rooms[roomID] == nil {
		h.rooms[roomID] = make(map[*Client]bool)
	}
	h.rooms[roomID][client] = true
	client.RoomID = roomID
}

func (h *Hub) LeaveRoom(client *Client) {
	if roomClients, exists := h.rooms[client.RoomID]; exists {
		delete(roomClients, client)
		if len(roomClients) == 0 {
			delete(h.rooms, client.RoomID)
		}
	}
	client.RoomID = ""
}

func (h *Hub) messageToBytes(message *Message) []byte {
	data, err := json.Marshal(message)
	if err != nil {
		return []byte(`{"type":"error","payload":{"code":"marshal_error","message":"Failed to marshal message"}}`)
	}
	return data
}
