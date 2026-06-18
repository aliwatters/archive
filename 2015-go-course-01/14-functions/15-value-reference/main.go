package main

import "fmt"

func main() {

	// everything is passed by value in go
	// BUT - slices, maps and channels are reference types
	// so modifying something modified like - z[0] = "new val" means that reference is modified

	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) // [Ali]

}

func changeMe(z []string) {
	z[0] = "Ali"
}
