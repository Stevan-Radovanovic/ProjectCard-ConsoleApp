package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := []string{initializeCard(), initializeCard(), initializeCard(), initializeCard(), initializeCard()}
	iterateCards(cards)
}

func initializeCard() string {
	attack := rand.Intn(15) + 1
	defence := rand.Intn(15) + 1
	return "Attack: " + strconv.Itoa(attack) + " | Defence: " + strconv.Itoa(defence)
}

func iterateCards(cards []string) {

	for i, card := range cards {
		fmt.Println(i, card)
	}

}
