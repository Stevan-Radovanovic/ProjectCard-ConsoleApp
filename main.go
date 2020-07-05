package main

import "fmt"

func main() {
	//var testCard string = "Attack 7 - Defence 5"
	card := initializeCard()
	fmt.Println(card)
}

func initializeCard() string {
	return "Attack X - Defence X"
}
