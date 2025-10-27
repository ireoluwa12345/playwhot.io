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
