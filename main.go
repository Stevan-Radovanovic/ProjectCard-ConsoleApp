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
	var card1 card
	var card2 card
	score1 := 0
	score2 := 0
	var gameMode string

	fmt.Println("\nStarting...")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nInitializing Ultimate Deck...")
	cards := initializeUltimateDeck()
	time.Sleep(1 * time.Second)

	fmt.Println("\nGenerating hands...")
	deck1, deck2 := cards.generateHand(handsize)
	time.Sleep(1 * time.Second)

	fmt.Println("\nGame is starting...")
	time.Sleep(1 * time.Second)

	fmt.Println("\nChoose your game mode! A/a for Attack, D/d for Defence")
	fmt.Scanf("%v\n", &gameMode)

	fmt.Println("\nConfiguring game mode...")
	time.Sleep(1 * time.Second)

	for i := 0; i < handsize; i++ {
		deck1.showDeck()
		time.Sleep(1 * time.Second)
		fmt.Println("\nChoose a card by index: ")
		fmt.Scanf("%d\n", &cardIndex)
		fmt.Println("\nYour opponent is choosing...")
		time.Sleep(1 * time.Second)
		deck1, card1 = deck1.removeCardByIndex(cardIndex - 1)
		deck2, card2 = deck2.removeCardByIndex(rand.Intn(currentHandSize))
		currentHandSize--
		result := playGame(card1, card2, gameMode)
		if result == 1 {
			score1++
		} else if result == -1 {
			score2++
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nCalculating results...")
	time.Sleep(1 * time.Second)

	finalValue := score1 - score2

	if finalValue > 0 {
		fmt.Println("\nCongratulations, you've won!")
	} else if finalValue < 0 {
		fmt.Println("\nOh no, you've lost!")
	} else {
		fmt.Println("\nIt's a draw!")
	}

	fmt.Println("\nUpdating High Scores...")
	updateStats(finalValue)
}
