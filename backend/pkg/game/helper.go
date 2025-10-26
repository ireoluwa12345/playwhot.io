package game

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func fillDeck(deckCards []Card, index int, suit string, validSuitNumbers []int) int {
	if len(validSuitNumbers) == 0 {
		return index
	}
	if index+len(validSuitNumbers) > len(deckCards) {
		return index
	}
	for i, number := range validSuitNumbers {
		deckCards[index+i] = Card{
			Suit:   suit,
			Number: fmt.Sprintf("%d", number),
			Whot:   suit == SuitWhot,
		}
	}

	return index + len(validSuitNumbers)
}

func randomInt(min, max int) int {
	return r.Intn(max-min+1) + min
}
