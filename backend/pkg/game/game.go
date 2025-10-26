package game

const (
	FirstHand = 4
)

type Status string

const (
	Ongoing Status = "ongoing"
)

type Game struct {
	Deck        *Deck          `json:"deck"`
	PlayerHands map[int][]Card `json:"player_hands"`
	CurrentCard Card           `json:"current_card"`
	CurrentTurn int            `json:"current_turn"`
	RoomID      int            `json:"room_id"`
	Direction   int            `json:"direction"`
	GameStatus  string         `json:"game_status"`
}

func (g *Game) StartGame(playerIDs []int, room_id int) {
	g.Deck = NewDeck()
	g.CurrentCard = g.Deck.Draw()
	g.RoomID = room_id
	g.CurrentTurn = playerIDs[0]

	for _, id := range playerIDs {
		g.PlayerHands[id] = g.Deck.Deal(FirstHand)
	}

	g.GameStatus = string(Ongoing)

}
