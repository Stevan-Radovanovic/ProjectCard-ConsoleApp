package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	handsize := 5
	var cardIndex int
	var card1 string
	var card2 string
	score1 := 0
	score2 := 0

	fmt.Println("\nStarting...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nInitializing Ultimate Deck...")
	cards := initializeUltimateDeck()

	fmt.Println("\nGenerating hands...")
	deck1, deck2 := cards.generateHand(handsize)

	fmt.Println("\nGame is starting...")

	for i := 0; i < handsize; i++ {
		deck1.showDeck()
		fmt.Println("\nChoose a card by index: ")
		fmt.Scanf("%d\n", &cardIndex)
		fmt.Println("\nCard index: ", cardIndex)
		fmt.Println("\nYour opponent is choosing...")
		deck1, card1 = deck1.removeCardByIndex(cardIndex - 1)
		deck2, card2 = deck2.removeCardByIndex(rand.Intn(handsize))
		if playGame(card1, card2) == true {
			score1++
		} else {
			score2++
		}
	}
}
