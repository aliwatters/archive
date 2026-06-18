package main

import "fmt"

func world() {
	fmt.Println("World")
}

func hello() {
	fmt.Print("Hello ")
}

func testDefer() {
	defer world() // this will not run til the end of func
	hello()
}

func main() {
	defer world() // same
	testDefer()
	// we're expecting Hello World\n World\n
}
