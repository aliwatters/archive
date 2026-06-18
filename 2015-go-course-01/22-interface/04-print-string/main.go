package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	// runs automatically when fmt.print of this struct is called
	return fmt.Sprintf("%v (%d)", p.name, p.age)
}

func main() {
	p := []person{
		{"Zeno", 23},
		{"Harry", 36},
		{"Pete", 14},
	}
	fmt.Println(p)
}
