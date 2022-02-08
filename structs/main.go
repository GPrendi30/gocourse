package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Guy",
		contactInfo: contactInfo{
			zipCode: 45500,
			email:   "alex@gmail.com",
		},
	}

	alex.print()
	alex.updateName("Joe")
	alex.print()
	//fmt.Printf("%+v", alex)
}

func (p *person) updateName(firstname string) {
	p.firstName = firstname
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
