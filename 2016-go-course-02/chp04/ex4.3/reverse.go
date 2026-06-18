package main

import "fmt"

// looks like the authors were trying to demonstrate the limitation
// of this approach - very limited functions
func reverse(a *[10]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	reverse(&a)
	fmt.Println(a)
}
