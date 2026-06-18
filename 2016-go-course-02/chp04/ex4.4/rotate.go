package main

import "fmt"

func rotate(s []int, n int) {
	l := len(s)
	if l == 0 || l <= n {
		return
	}

	// create tmp slice of last n entries
	tmp := make([]int, n)
	copy(tmp, s[:n]) // copy(dst, src)!

	copy(s, s[n:])
	copy(s[l-n:], tmp)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	z := 5 // must be max len(s)

	fmt.Println(s)
	rotate(s, z)
	fmt.Println(s)
}
