package main

import "fmt"

// composite type
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 44}
	p2 := person{"Miss", "Moneypenny", 24}

	fmt.Println(p1, p2.first, p2.last, p2.age)

}
