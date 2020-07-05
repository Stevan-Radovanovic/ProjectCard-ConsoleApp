package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var cards deck
	//cards = deck{cards.newCard(), cards.newCard(), cards.newCard(), cards.newCard(), cards.newCard()}
	cards = cards.initialize()
	cards.iterate()
}
