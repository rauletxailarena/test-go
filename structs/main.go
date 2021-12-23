package main

import "fmt"

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

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName

}

func main() {
	raulon := person{firstName: "Raul", lastName: "Raulon"}
	raulon.updateName("Enmayao")
	raulon.print()
}
