package main

import "fmt"

func zero(x int) *int {
	fmt.Println("in zero x's address", &x, "and value", x)
	x = 0
	fmt.Println("in zero x's address", &x, "and value", x)
	x = 42
	return &x
}

func main() {
	x := 5 // initialize x
	fmt.Println("in main x's address", &x, "and value", x)
	y := zero(x) // passes the value of x (it's copied into a new address)

	fmt.Println("in main x's address", &x, "and value", x)

	fmt.Println("and y returned from 0", y, "has value", *y)

	// d has a new address but
	z := *y
	fmt.Println("z address and value", &z, z)

}
