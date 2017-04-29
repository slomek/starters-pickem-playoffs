package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func readResults() Playoffs {
	var results []Series
	yamlFile, err := ioutil.ReadFile("scores.yml")
	if err != nil {
		log.Fatalf("Cannot read scores.yml: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &results)
	if err != nil {
		log.Fatalf("Corrupt scores.yml: %v", err)
	}

	return results
}
