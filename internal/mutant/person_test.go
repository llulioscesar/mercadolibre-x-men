package mutant_test

import (
	"github.com/llulioscesar/mercadolibre-x-men/internal/mutant"
	"testing"
)

func TestPerson_IsMutant(t *testing.T) {
	p := mutant.Person{DNA: mutant.DNAStructure{Nucleotides: []string{"AAATCGA", "PLKATG"}}}

	p.IsMutant()
}