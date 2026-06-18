package main

import "fmt"

const welcome = "Hello, 世界"

func main() {
	for i, r := range welcome {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
