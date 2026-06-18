package main

import "fmt"

func halveIt(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ": ")
		fmt.Println(halveIt(i))
	}
}
