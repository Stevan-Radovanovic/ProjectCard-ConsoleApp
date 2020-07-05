package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Starting...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Initializing Ultimate Deck...")
	cards := initializeUltimateDeck()

	fmt.Println("Generating hands...")
	deck1, deck2 := cards.generateHand(5)

	fmt.Println("Deck 1:")
	deck1.showDeck()

	fmt.Println("Deck 2:")
	deck2.showDeck()
}
