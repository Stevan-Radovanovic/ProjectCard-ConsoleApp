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
	range1 := rand.Intn(50 - handSize)
	range2 := rand.Intn(50 - handSize)
	return d[range1 : range1+handSize], d[range2 : range2+handSize]
}

func (d deck) removeCardByIndex(index int) (deck, string) {

	removed := d[index]
	copy(d[index:], d[index+1:])
	d[len(d)-1] = ""
	d = d[:len(d)-1]
	return d, removed
}

func playGame(card1 string, card2 string) bool {
	return true
}
