package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	a := 42
	fmt.Println("a is", a)
	fmt.Println("a's address is", &a) // returns hex location in memory of a

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // reads into var
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")

	// pointers again

	var b *int = &a // b is of type "pointer to int" *int
	fmt.Println("b is", b)
	fmt.Println("value at b", *b) // dereference b - gives value.

	// now change the values in memory address (affects b and a)

	*b = 101
	fmt.Println("a and b address", &a, b)
	fmt.Println("a and b value", a, *b)

	c := 42
	b = &c
	fmt.Println("now b is", b, *b)
}
