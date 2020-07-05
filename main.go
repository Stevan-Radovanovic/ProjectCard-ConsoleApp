package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cards := initializeUltimateDeck()
	deck1, deck2 := cards.generateHand(5)
	deck1.showDeck()
	deck2.showDeck()
}
