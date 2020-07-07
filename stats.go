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

func updateStats(finalValue int) {
	input, err := ioutil.ReadFile("stats.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	numberOfGames, _ := strconv.Atoi(lines[3])
	numberOfWins, _ := strconv.Atoi(lines[5])
	numberOfDraws, _ := strconv.Atoi(lines[7])
	numberOfLosses, _ := strconv.Atoi(lines[9])
	totalValue, _ := strconv.Atoi(lines[11])
	numberOfGames++
	if finalValue > 0 {
		numberOfWins++
	} else if finalValue < 0 {
		numberOfLosses++
	} else {
		numberOfDraws++
	}
	totalValue += finalValue
	lines[3] = strconv.Itoa(numberOfGames)
	lines[5] = strconv.Itoa(numberOfWins)
	lines[7] = strconv.Itoa(numberOfDraws)
	lines[9] = strconv.Itoa(numberOfLosses)
	lines[11] = strconv.Itoa(totalValue)
	output := strings.Join(lines, "\n")
	ioutil.WriteFile("stats.txt", []byte(output), 0644)
}

func printStats() {
	input, err := ioutil.ReadFile("stats.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(string(input))
	fmt.Println("Pres ENTER to continue...")
	fmt.Scanf("\n")
}
