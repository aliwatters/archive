package main

import "fmt"

// composite type
type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func (p person) greet() string {
	return "Hello"
}

type doubleZero struct {
	person
	first         string
	licenseToKill bool
}

func (d doubleZero) greet() string {
	return d.last + ", " + d.person.first + " " + d.last
}

func main() {
	p1 := doubleZero{
		person:        person{"James", "Bond", 44},
		licenseToKill: true,
		first:         "007", // this comma is not optional
	}
	p2 := person{"Miss", "Moneypenny", 24}

	fmt.Println(p1)
	// note on doubleZero - fullname and fields are promoted unless there is a clash - but note that first is from the person
	fmt.Println(p1.fullName(), p2.fullName())
	fmt.Println(p2.greet(), p1.greet())
}
