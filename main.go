package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	handsize := 5
	currentHandSize := 5
	var cardIndex int
	var card1 string
	var card2 string
	score1 := 0
	score2 := 0

	fmt.Println("\nStarting...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nInitializing Ultimate Deck...")
	cards := initializeUltimateDeck()
	time.Sleep(4 * time.Second)

	fmt.Println("\nGenerating hands...")
	deck1, deck2 := cards.generateHand(handsize)
	time.Sleep(4 * time.Second)

	fmt.Println("\nGame is starting...")
	time.Sleep(3 * time.Second)

	for i := 0; i < handsize; i++ {
		deck1.showDeck()
		time.Sleep(2 * time.Second)
		fmt.Println("\nChoose a card by index: ")
		fmt.Scanf("%d\n", &cardIndex)
		fmt.Println("\nYour opponent is choosing...")
		time.Sleep(3 * time.Second)
		deck1, card1 = deck1.removeCardByIndex(cardIndex - 1)
		deck2, card2 = deck2.removeCardByIndex(rand.Intn(currentHandSize))
		currentHandSize--
		result := playGame(card1, card2)
		if result == 1 {
			score1++
		} else if result == 0 {
			score2++
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("\nCalculating results...")
	time.Sleep(4 * time.Second)
	if score1 > score2 {
		fmt.Println("\nCongratulations, you've won!")
	} else {
		fmt.Println("\nOh no, you've lost!")
	}
}
