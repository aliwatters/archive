package main

import "fmt"

// returns a function that returns a string
func makeGreeter(who string) func() string {
	// this is the closure scope
	salutation := "Hello"
	return func() string {
		return salutation + " " + who
	}
}

func main() {
	// extending to a curried example
	greetWorld := makeGreeter("World")
	greetAli := makeGreeter("Ali")

	fmt.Println(greetWorld())
	fmt.Println(greetAli())

	// make a new scope - no control flow/function needed
	{
		y := 100
		fmt.Println(y)
	}
	// but here - y is not in scope!
	// fmt.Println(y)
}
