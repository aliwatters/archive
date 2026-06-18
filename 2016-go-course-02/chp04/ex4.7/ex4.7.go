package main

import (
	"fmt"
	"unicode/utf8"
)

// Reverse reverses unicode characters in a slice of bytes
func Reverse(s []byte) []byte {

	// out := s[:0] // this fails with "mirrored" text - slice is modified in place.
	out := []byte{}
	i, j := len(s), len(s)

	for i > 0 {
		i--
		for !utf8.Valid(s[i:j]) {
			i--
		}

		// fmt.Printf("\ni: %d, j: %d, rune: %#U", i, j, r)
		out = append(out, s[i:j]...)
		j = i
	}

	return out
}

func main() {
	a := []byte("Hello World! 日本語")
	fmt.Printf("% x %s\n", a, string(a[:]))
	a = Reverse(a)
	fmt.Printf("% x %s\n", a, string(a[:]))
}
