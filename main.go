package main

import "fmt"

func main() {
	//var testCard string = "Attack 7 - Defence 5"
	cards := []string{initializeCard(), initializeCard()}
	iterateCards(cards)
}

func initializeCard() string {
	return "Attack X - Defence X"
}

func iterateCards(cards []string) {

	for i, card := range cards {
		fmt.Println(i, card)
	}

}
