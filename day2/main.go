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

func matchScore(left, right string) int {
	const (
		win  = 6
		draw = 3
		lost = 0
	)

	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	isDraw := left == "A" && right == "X" || left == "B" && right == "Y" || left == "C" && right == "Z"

	isWin := left == "A" && right == "Y" ||
		left == "B" && right == "Z" ||
		left == "C" && right == "X"

	if isDraw {
		return draw
	}

	if isWin {
		return win
	}
	return lost
}

type turn string

func (t turn) MyScore() int {
	switch t {
	case "X": // rock
		return 1
	case "Y":
		return 2
	case "Z":
		return 3

	}
	panic("invalid value")
}

func main() {
	f, err := os.Open("./data.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)

	total := 0
	for sc.Scan() {
		line := sc.Text()
		splitedLine := strings.Split(strings.TrimSuffix(line, "\n"), " ")
		left, right := strings.TrimSpace(splitedLine[0]), strings.TrimSpace(splitedLine[1])
		winScore := matchScore(left, right)

		total += winScore + turn(right).MyScore()

	}

	log.Printf("Score is: %d\n", total)
}
