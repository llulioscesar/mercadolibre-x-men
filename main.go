package main

import (
	"log"
	"prueba/internal"
)

func main() {
	p1 := internal.Person{
		DNA: internal.DNAStructure{
			Nucleotides: []string{
				"ATGCGA",
				"CAGTGC",
				"TTATGT",
				"AGAAGG",
				"CCCCTA",
				"TCACTG"},
		},
	}
	p2 := internal.Person{
		DNA: internal.DNAStructure{
			Nucleotides: []string{
				"AAAAGA",
				"CAGTGC",
				"TTATGT",
				"AGAAGG",
				"CCCCAA",
				"TCACTG"},
		},
	}
	p3 := internal.Person{
		DNA: internal.DNAStructure{
			Nucleotides: []string{
				"ATGCGA",
				"ATCGTA",
				"AGCGTA",
				"ATGCGA",
				"CCACAA",
				"CACACA"},
		},
	}
	log.Println(p1.IsMutant())
	log.Println("-----")
	log.Println(p2.IsMutant())
	log.Println("-----")
	log.Println(p3.IsMutant())
}
