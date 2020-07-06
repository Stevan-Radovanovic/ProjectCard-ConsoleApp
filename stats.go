package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func changePlayerName(newName string) {

	input, err := ioutil.ReadFile("stats.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	lines[1] = newName
	output := strings.Join(lines, "\n")
	ioutil.WriteFile("stats.txt", []byte(output), 0644)
}

func updateStats(victory bool) {
	input, err := ioutil.ReadFile("stats.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	numberOfGames, _ := strconv.Atoi(lines[3])
	numberOfWins, _ := strconv.Atoi(lines[5])
	numberOfLosses, _ := strconv.Atoi(lines[7])
	numberOfGames++
	if victory == true {
		numberOfWins++
	} else {
		numberOfLosses++
	}
	lines[3] = strconv.Itoa(numberOfGames)
	lines[5] = strconv.Itoa(numberOfWins)
	lines[7] = strconv.Itoa(numberOfLosses)
	output := strings.Join(lines, "\n")
	ioutil.WriteFile("stats.txt", []byte(output), 0644)
}
