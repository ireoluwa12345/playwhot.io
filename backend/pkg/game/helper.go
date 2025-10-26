package game

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// fillDeck creates cards for a specific suit with special card properties
func fillDeck(deckCards []Card, index int, suit string, validSuitNumbers []int) int {
	if len(validSuitNumbers) == 0 {
		return index
	}
	if index+len(validSuitNumbers) > len(deckCards) {
		return index
	}
	for i, number := range validSuitNumbers {
		card := Card{
			Suit:   suit,
			Number: fmt.Sprintf("%d", number),
			Whot:   suit == SuitWhot,
			Value:  number,
		}

		// Set special types based on card number
		switch number {
		case 1:
			card.SpecialType = string(SpecialHoldOn)
		case 2:
			card.SpecialType = string(SpecialPickTwo)
		case 5:
			card.SpecialType = string(SpecialPickThree)
		case 8:
			card.SpecialType = string(SpecialSuspension)
		}

		deckCards[index+i] = card
	}

	return index + len(validSuitNumbers)
}

func randomInt(min, max int) int {
	return r.Intn(max-min+1) + min
}
