package game

/*
Helper Functions

randomInt
fillDeck

*/

const (
	SuitCircle     = "circle"
	SuitTriangle   = "triangle"
	SuitCross      = "cross"
	SuitSquare     = "square"
	SuitStar       = "star"
	SuitWhot       = "whot"
	MaxDeckSize    = 54
	WhotCardNumber = 4
)

type Card struct {
	Suit   string `json:"suit"`
	Number string `json:"number"`
	Whot   bool   `json:"whot"`
}

type SuitCards struct {
	Suit    string
	Numbers []int
}

type Deck struct {
	Cards   [MaxDeckSize]Card
	TopCard int
}

func NewDeck() *Deck {
	deck := &Deck{TopCard: 0}
	index := 0

	suits := []SuitCards{
		{Suit: SuitCircle, Numbers: []int{1, 2, 3, 4, 5, 7, 8, 10, 11, 12, 13, 14}},
		{Suit: SuitTriangle, Numbers: []int{1, 2, 3, 4, 5, 7, 8, 10, 11, 12, 13, 14}},
		{Suit: SuitCross, Numbers: []int{1, 2, 3, 5, 7, 10, 11, 13, 14}},
		{Suit: SuitSquare, Numbers: []int{1, 2, 3, 5, 7, 10, 11, 13, 14}},
		{Suit: SuitStar, Numbers: []int{1, 2, 3, 4, 5, 7, 8, 10, 11, 12, 13, 14}},
	}

	for _, suitCards := range suits {
		index = fillDeck(deck.Cards[:], index, suitCards.Suit, suitCards.Numbers)
	}

	for i := 0; i < WhotCardNumber; i++ {
		j := randomInt(0, MaxDeckSize)
		deck.createWhotCard(j)
	}

	deck.Shuffle()

	return deck
}

func (d *Deck) createWhotCard(key int) {
	d.Cards[key] = Card{
		Suit:   SuitWhot,
		Number: "20",
		Whot:   true,
	}
}

func (d *Deck) Shuffle() *Deck {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := randomInt(0, i)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

	return d
}

func (d *Deck) Draw() Card {
	if d.TopCard >= len(d.Cards) {
		return Card{}
	}
	card := d.Cards[d.TopCard]
	d.TopCard++
	return card
}

func (d *Deck) Deal(count int) []Card {
	if d.TopCard+count > len(d.Cards) {
		count = len(d.Cards) - d.TopCard
	}

	cards := make([]Card, count)
	for i := 0; i < count; i++ {
		cards[i] = d.Cards[d.TopCard+i]
	}
	d.TopCard += count
	return cards
}

func (d *Deck) Remaining() int {
	return len(d.Cards) - d.TopCard
}
