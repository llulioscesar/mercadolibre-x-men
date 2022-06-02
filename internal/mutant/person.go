package mutant

import (
	"regexp"
)

type Person struct {
	DNA DNAStructure
}

func (person *Person) IsMutant() bool {
	err := person.DNA.CheckMatrix()
	if err != nil {
		return false
	}
	return person.findMutantBlocks()
}

func (person *Person) findMutantBlocks() bool {
	var result []string

	re := regexp.MustCompile(`A{4,}|C{4,}|G{4,}|T{4,}`)
	for _, sequence := range person.DNA.Nucleotides {
		test := re.FindString(sequence)
		if len(test) > 0 {
			result = append(result, test)
		}
	}

	right := person.DNA.GetRightDiagonal()
	for _, sequence := range right {
		test := re.FindString(sequence)
		if len(test) > 0 {
			result = append(result, test)
		}
	}

	left := person.DNA.GetLeftDiagonal()
	for _, sequence := range left {
		test := re.FindString(sequence)
		if len(test) > 0 {
			result = append(result, test)
		}
	}
	return len(result) > 1
}
