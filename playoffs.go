package main

import "fmt"

type Series struct {
	Teams string `yaml:"teams"`
	Score string `yaml:"score"`
}

type Playoffs []Series

func (p Playoffs) SeriesResult(matchup string) (string, error) {
	for _, r := range p {
		if r.Teams == matchup {
			return r.Score, nil
		}
	}
	return "", fmt.Errorf("No result for matchup '%s'", matchup)
}
