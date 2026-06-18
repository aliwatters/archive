package main

import "fmt"

func main() {
	halveIt := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i, ": ")
		fmt.Println(halveIt(i))
	}
}
