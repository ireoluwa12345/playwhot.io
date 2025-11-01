package models

import (
	"errors"
	"time"
)

var (
	ErrDuplicateEmail     = errors.New("account exists, try logging in")
	ErrDuplicateUsername  = errors.New("username exists, chose a different username")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PlayerData struct {
	ID       int
	Username string
}

type Room struct {
	ID         int          `json:"room_id"`
	Name       string       `json:"name"`
	CreatedBy  int          `json:"user_id"`
	Timestamp  time.Time    `json:"-"`
	Players    []PlayerData `json:"players"`
	MaxPlayers int          `json:"max_players"`
	Password   string       `json:"room_password"`
	IsPrivate  bool         `json:"is_private"`
}
