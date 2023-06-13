package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// structs can have other structs inside
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	p := person{
		firstName: "Jack",
		lastName:  "Sparrow",
		contactInfo: contactInfo{
			email:   "a@a.a",
			zipCode: 10000,
		},
	}

	p.updateName("Jim")

	p.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
