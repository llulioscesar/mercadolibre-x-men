package main

import (
	"log"
)

func main() {
	adn1 := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
	adn2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	err := checkMatrix(adn1)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(isMutant(adn1))

	err = checkMatrix(adn2)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(isMutant(adn2))
}
