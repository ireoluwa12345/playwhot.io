package game

import "errors"

const (
	FirstHand = 6 // Updated to 6 cards as per Whot rules
)

type Status string

const (
	Ongoing Status = "ongoing"
)

// Game represents the current state of a Whot game
type Game struct {
	Deck              *Deck          `json:"deck"`
	PlayerHands       map[int][]Card `json:"player_hands"`
	CurrentCard       Card           `json:"current_card"`
	CurrentTurn       int            `json:"current_turn"`
	RoomID            int            `json:"room_id"`
	GameStatus        string         `json:"game_status"`
	PendingDraws      int            `json:"pending_draws"`
	SuspendedPlayer   int            `json:"suspended_player"`
	ChosenSuit        string         `json:"chosen_suit"`
	LastCardAnnounced map[int]bool   `json:"last_card_announced"`
}

// StartGame initializes a new game with shuffled deck and deals cards to players
func (g *Game) StartGame(playerIDs []int, room_id int) {
	g.Deck = NewDeck()
	g.CurrentCard = g.Deck.Draw()
	g.RoomID = room_id
	g.CurrentTurn = playerIDs[0]
	g.PendingDraws = 0
	g.SuspendedPlayer = 0
	g.ChosenSuit = ""
	g.LastCardAnnounced = make(map[int]bool)

	// Initialize player hands
	g.PlayerHands = make(map[int][]Card)
	for _, id := range playerIDs {
		g.PlayerHands[id] = g.Deck.Deal(FirstHand)
	}

	g.GameStatus = string(Ongoing)
}

func (g *Game) PlayCard(playerID int, card Card, chosenSuit string) error {
	return errors.New("hello")
}
