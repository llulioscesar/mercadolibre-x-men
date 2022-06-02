package main

type Person struct {
	DNA
}

func (person *Person) isMutant() bool {
	err := person.DNA.CheckMatrix()
	if err != nil {
		return false
	}
	
	
}

func (person *Person) ()  {
	
}