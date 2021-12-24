package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (person *person) updateName(newFirstName string) {
	person.firstName = newFirstName
}

func main() {
	raulon := person{firstName: "Raul", lastName: "Raulon"}
	raulon.updateName("Enmayao")
	raulon.print()
}
