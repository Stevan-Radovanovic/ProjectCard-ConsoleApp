package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type deck []string

func (d deck) showDeck() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) newCard() string {
	attack := rand.Intn(15) + 1
	defence := rand.Intn(15) + 1
	return "Attack: " + strconv.Itoa(attack) + " | Defence: " + strconv.Itoa(defence)
}

func initializeUltimateDeck() deck {

	var d deck
	for i := 0; i < 50; i++ {
		d = append(d, d.newCard())
	}

	return d
}

func (d deck) generateHand(handSize int) (deck, deck) {
	return d[rand.Intn(50-handSize)+handSize:], d[rand.Intn(50-handSize)+handSize:]
}
