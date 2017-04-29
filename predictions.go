package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func readGuesses(player string) []Series {
	filename := fmt.Sprintf("predictions/%s.yml", player)

	var guesses []Series
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &guesses)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return guesses
}
