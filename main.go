package main

import (
	"fmt"
	"sort"
)

func main() {
	players := []string{"trey", "leigh", "tas", "skeets", "slomek"}
	results := readResults()

	scores := playersScores(players, results)
	sort.Slice(scores, func(i, j int) bool { return scores[i].score > scores[j].score })

	for idx, s := range scores {
		fmt.Printf("%10s: %d%s\n", s.player, s.score, medal(idx, scores))
	}
}

type score struct {
	player string
	score  int
}

func playersScores(pp []string, results Playoffs) []score {
	var ss []score
	for _, p := range pp {
		ss = append(ss, score{player: p, score: playerScore(p, results)})
	}
	return ss
}

func playerScore(player string, playoffs Playoffs) int {
	scores := readGuesses(player)

	return calculatePoints(playoffs, scores)
}

func medal(index int, scores []score) string {
	const (
		first  = " 🥇"
		second = " 🥈"
		third  = " 🥉"
	)

	if len(scores) == 0 {
		return ""
	}

	score := scores[index]
	if index == 0 || score.score == scores[0].score {
		return first
	}
	if index == 1 || score.score == scores[1].score {
		return second
	}
	if index == 2 || score.score == scores[2].score {
		return third
	}
	return ""
}
