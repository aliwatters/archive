package main

import "fmt"

func main() {
	// declare
	var f string
	// assign
	f = "hello"
	// initialize
	var g = 3.1417
	var h int

	// idiomatic initialize - the preferred way
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", c, c)
	fmt.Printf("%T\t%v\n", d, d)

	fmt.Printf("%T\t%v\n", f, f)
	fmt.Printf("%T\t%v\n", g, g)
	fmt.Printf("%T\t%v\n", h, h)
}
