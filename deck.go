package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type deck []string

func (d deck) iterate() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) newCard() string {
	attack := rand.Intn(15) + 1
	defence := rand.Intn(15) + 1
	return "Attack: " + strconv.Itoa(attack) + " | Defence: " + strconv.Itoa(defence)
}

func (d deck) initialize() deck {
	d = append(d, d.newCard())
	d = append(d, d.newCard())
	d = append(d, d.newCard())
	d = append(d, d.newCard())
	d = append(d, d.newCard())
	return d
}
