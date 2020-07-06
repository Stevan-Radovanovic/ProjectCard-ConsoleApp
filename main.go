package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	handsize := 5
	var cardIndex int
	score1 := 0
	score2 := 0

	fmt.Println("\nStarting...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nInitializing Ultimate Deck...")
	cards := initializeUltimateDeck()

	fmt.Println("\nGenerating hands...")
	deck1, deck2 := cards.generateHand(handsize)

	fmt.Println("\nGame is starting...")
	deck1.showDeck()
	deck2.showDeck()

	for i := 0; i < handsize; i++ {
		fmt.Println("Choose a card by index: ")
		fmt.Scanf("%d\n", &cardIndex)
		fmt.Println("Card index: ", cardIndex)
		fmt.Println("Your opponent is choosing...")
		card1 := deck1.removeCardByIndex(cardIndex - 1)
		card2 := deck2.removeCardByIndex(rand.Intn(handsize))
		if playGame(card1, card2) == true {
			score1++
		} else {
			score2++
		}
	}
}
