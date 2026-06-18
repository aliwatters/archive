package main

import "fmt"

// not idiomatic
type foo int

func main() {
	var age foo
	age = 42
	fmt.Printf("%T %v\n", age, age)

	ageB := int(100)

	// age + ageB -- not working
	fmt.Println(int(age) + ageB)
}
