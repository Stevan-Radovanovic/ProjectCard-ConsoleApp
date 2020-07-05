package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//cards = deck{cards.newCard(), cards.newCard(), cards.newCard(), cards.newCard(), cards.newCard()}
	cards := initializeDeck()
	cards.showDeck()
}
