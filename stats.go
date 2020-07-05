package main

import (
	"io/ioutil"
	"strings"
)

func changePlayerName(newName string) {

	input, _ := ioutil.ReadFile("stats.txt")
	lines := strings.Split(string(input), "\n")
	lines[1] = newName
	output := strings.Join(lines, "\n")
	ioutil.WriteFile("stats.txt", []byte(output), 0644)

}

func updateStats() {

}
