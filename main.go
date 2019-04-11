package main

import (
	"fmt"
)

func main() {
	players := []string{"trey", "leigh", "tas", "skeets", "slomek"}

	results := readResults()

	for _, p := range players {
		fmt.Printf("%s: %d\n", p, playerScore(p, results))
	}
}

func playerScore(player string, playoffs Playoffs) int {
	scores := readGuesses(player)

	return calculatePoints(playoffs, scores)
}
