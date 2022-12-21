package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

func Game(op, you string) int {
	var turnScore int
	switch you {
	case "X":
		turnScore = 1
	case "Y":
		turnScore = 2
	case "Z":
		turnScore = 3
	default:
		panic("invalid value")
	}

	const (
		win  = 6
		draw = 3
		lost = 0
	)

	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	switch op {
	case "A": // rock
		switch you {
		case "X": // rock => draw
			return turnScore + draw
		case "Y": // paper => lost
			return turnScore + win
		case "Z": // scissor => win
			return turnScore + lost
		}
	case "B": // paper
		switch you {
		case "X": // rock => lost
			return turnScore + lost
		case "Y": // paper => draw
			return turnScore + draw
		case "Z": // scissor => lost
			return turnScore + win
		}
	case "C": // scissors
		switch you {
		case "X": // rock => lost
			return turnScore + win
		case "Y": // paper =>  win
			return turnScore + lost
		case "Z": // draw
			return turnScore + draw
		}
	}
	return 0
}

func main() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	total := 0
	for sc.Scan() {
		line := sc.Text()
		words := strings.Split(line, " ")
		if len(words) != 2 {
			log.Printf("Invalid data for line")
			continue
		}
		op, you := words[0], words[1]
		total += Game(op, you)
	}

	log.Printf("Score is: %d\n", total)
}
