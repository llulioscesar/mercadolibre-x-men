package internal

import (
	"errors"
	"regexp"
	"strings"
)

type DNAStructure struct {
	Nucleotides []string
}

func (dna *DNAStructure) CheckMatrix() error {
	adnSize := len(dna.Nucleotides)
	re := regexp.MustCompile(`(?s)[ACGT]+`)

	for _, sequence := range dna.Nucleotides {
		if len(sequence) != adnSize {
			return errors.New("Invalid Matrix size")
		}
		if re.FindString(sequence) != sequence {
			return errors.New("Invalid DNAStructure sequence")
		}
	}
	return nil
}

func (dna *DNAStructure) getRightDiagonal(nucleotides []string) (result []string) {
	var all [][]string
	for s := 0; s < len(nucleotides); s++ {
		var temp []string
		for y, x := s, 0; y >= 0; y, x = y-1, x+1 {
			temp = append(temp, string([]rune(nucleotides[y])[x]))
		}
		all = append(all, temp)
	}
	for s := 1; s < len(nucleotides[0]); s++ {
		var temp []string
		for y, x := len(nucleotides)-1, s; x < len(nucleotides[0]); y, x = y-1, x+1 {
			temp = append(temp, string([]rune(nucleotides[y])[x]))
		}
		all = append(all, temp)
	}
	for _, sequence := range all {
		result = append(result, strings.Join(sequence, ""))
	}
	return result
}

func (dna *DNAStructure) GetRightDiagonal() []string {
	return dna.getRightDiagonal(dna.Nucleotides)
}

func (dna *DNAStructure) GetLeftDiagonal() []string {
	reverse := dna.reverseMatix(dna.Nucleotides)
	return dna.getRightDiagonal(reverse)
}

func (dna *DNAStructure) reverseMatix(nucleotides []string) (result []string) {
	for _, v := range nucleotides {
		result = append(result, dna.reverseSting(v))
	}
	return
}

func (dna *DNAStructure) reverseSting(sequence string) (result string) {
	for _, v := range sequence {
		result = string(v) + result
	}
	return
}
