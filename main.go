package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	clear()
	fmt.Println("\nStarting...")
	time.Sleep(2 * time.Second)
	fmt.Println("\nWelcome to ProjectCard!")
	fmt.Println("\nPress ENTER to continue...")
	fmt.Scanf("\n")
	for {
		clear()
		var choice int
		fmt.Print("Pro")
		time.Sleep(1 * time.Second)
		fmt.Print("ject")
		time.Sleep(1 * time.Second)
		fmt.Print("Card")
		time.Sleep(1 * time.Second)
		fmt.Print("!")
		time.Sleep(1 * time.Second)
		fmt.Println("\n1. Play an Attack Round")
		fmt.Println("2. Play a Defence Round")
		fmt.Println("3. See Stats")
		fmt.Println("4. Change Player Name")
		fmt.Println("5. Exit the Game")
		fmt.Print("Choose an option: ")
		fmt.Scanf("%d\n", &choice)
		clear()
		switch choice {
		case 1:
			round("a")
		case 2:
			round("d")
		case 3:
			printStats()
		case 4:
			changePlayerName("test")
		case 5:
			{
				fmt.Println("Bye!")
				fmt.Print("Press ENTER to exit...")
				fmt.Scanf("\n")
				return
			}
		default:
			{
				fmt.Print("Wrong entry, press ENTER to continue...")
				fmt.Scanf("\n")
			}
		}
	}
}

func round(mode string) {

	clear()
	handsize := 5
	currentHandSize := 5
	var cardIndex int
	var card1 card
	var card2 card
	score1 := 0
	score2 := 0

	fmt.Println("Initializing Ultimate Deck...")
	var cards deck
	cards.initializeUltimateDeck()
	time.Sleep(2 * time.Second)

	fmt.Println("\nGenerating hands...")
	deck1, deck2 := cards.generateHand(handsize)
	time.Sleep(1 * time.Second)

	fmt.Println("Game is starting...")
	time.Sleep(1 * time.Second)

	fmt.Println("Configuring game mode...")
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
		result := playGame(card1, card2, mode)
		if result == 1 {
			score1++
		} else if result == -1 {
			score2++
		}
		time.Sleep(3 * time.Second)
		clear()
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
	fmt.Print("Press ENTER to continue...")
	fmt.Scanf("\n")
}
