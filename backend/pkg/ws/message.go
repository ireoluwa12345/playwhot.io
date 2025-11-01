package ws

import (
	"fmt"

	"playwhot.io/pkg/game"
)

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	RoomID  string      `json:"room_id,omitempty"`
	UserID  string      `json:"user_id,omitempty"`
}

const (
	MessageTypeJoin  = "join"
	MessageTypeLeave = "leave"
	MessageTypeError = "error"

	MessageTypeCardPlay = "card_play"
	MessageTypeDraw     = "draw"
	MessageTypePass     = "pass"

	MessageTypeGameState    = "game_state"
	MessageTypePlayerUpdate = "player_update"
)

type JoinPayload struct {
	Username string `json:"username"`
	RoomID   string `json:"room_id"`
}

type CardPlayPayload struct {
	CardID string `json:"card_id"`
	Target string `json:"target,omitempty"`
}

type GameStatePayload struct {
	CurrentPlayer string      `json:"current_player"`
	Deck          []game.Card `json:"deck"`
	Players       []int       `json:"players"`
}

type ErrorPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type MessageHandler func(*Client, *Message) error

type MessageRouter struct {
	handlers map[string]MessageHandler
}

func NewMessageRouter() *MessageRouter {
	router := &MessageRouter{
		handlers: make(map[string]MessageHandler),
	}

	router.RegisterHandler(MessageTypeJoin, handleJoin)
	router.RegisterHandler(MessageTypeCardPlay, handleCardPlay)
	router.RegisterHandler(MessageTypeDraw, handleDraw)

	return router
}

func (r *MessageRouter) RegisterHandler(msgType string, handler MessageHandler) {
	r.handlers[msgType] = handler
}

func (r *MessageRouter) Route(client *Client, message *Message) error {
	handler, exists := r.handlers[message.Type]
	if !exists {
		return fmt.Errorf("unknown message type: %s", message.Type)
	}
	return handler(client, message)
}

func handleJoin(client *Client, message *Message) error {
	payload, ok := message.Payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid join payload")
	}

	roomID, ok := payload["room_id"].(string)
	if !ok {
		return fmt.Errorf("room_id required")
	}

	client.hub.JoinRoom(client, roomID)

	response := &Message{
		Type:    MessageTypePlayerUpdate,
		Payload: map[string]interface{}{"status": "joined", "room_id": roomID},
		RoomID:  roomID,
	}

	client.send <- client.hub.messageToBytes(response)
	return nil
}

func handleCardPlay(client *Client, message *Message) error {
	payload, ok := message.Payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid card_play payload")
	}

	cardID, ok := payload["card_id"].(string)
	if !ok {
		return fmt.Errorf("card_id required")
	}

	broadcastMsg := &Message{
		Type: MessageTypeGameState,
		Payload: map[string]interface{}{
			"action":    "card_played",
			"card_id":   cardID,
			"player_id": client.UserID,
		},
		RoomID: client.RoomID,
	}

	client.hub.broadcast <- broadcastMsg
	return nil
}

func handleDraw(client *Client, message *Message) error {
	payload, ok := message.Payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid draw payload")
	}

	count := 1
	if countVal, exists := payload["count"]; exists {
		if c, ok := countVal.(float64); ok {
			count = int(c)
		}
	}

	broadcastMsg := &Message{
		Type: MessageTypeGameState,
		Payload: map[string]interface{}{
			"action":    "cards_drawn",
			"count":     count,
			"player_id": client.UserID,
		},
		RoomID: client.RoomID,
	}

	client.hub.broadcast <- broadcastMsg
	return nil
}
