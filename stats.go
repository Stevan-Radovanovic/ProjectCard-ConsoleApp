package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

}
