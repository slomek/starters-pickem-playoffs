package main

import (
	"math"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func calculatePoints(playoffs Playoffs, guesses []Series) int {
	total := 0

	for _, s := range guesses {
		result, err := playoffs.SeriesResult(s.Teams)
		if err != nil {
			log.Debugf("No points for '%s': %v", s.Teams, err)
			continue
		}

		total = total + seriesPoints(result, s.Score)
	}

	return total
}

func seriesPoints(score, guess string) int {
	sA, sB := getWinsByTeam(score)
	gA, gB := getWinsByTeam(guess)

	return byTeamPoints(sA, gA) + byTeamPoints(sB, gB)
}

func getWinsByTeam(score string) (int, int) {
	scores := strings.Split(score, "-")
	a, _ := strconv.Atoi(scores[0])
	b, _ := strconv.Atoi(scores[1])
	return a, b
}

func byTeamPoints(wins, guess int) int {
	if wins == 4 && guess == 4 {
		return 5
	}

	if wins >= guess {
		return int(math.Min(float64(wins), float64(guess)))
	}

	return wins - guess
}
