package main

import "fmt"

// EliminateDups removes duplicate adjacent strings in a slice of strings
func EliminateDups(s []string) []string {
	l := len(s)
	if l == 0 {
		return s
	}

	a := s[0]
	for i := 1; i < l; i++ {
		b := s[i]
		// fmt.Println("testing", a, b)
		if a == b {
			// fmt.Println("switching", s[i:l-1], s[i+1:l])
			copy(s[i:l-1], s[i+1:l])
			l-- // removed an element
			i-- // test current position again - catches triples
		}
		a = b
	}

	return s[:l]
}

func main() {
	// s := []string{"b", "b", "b"}
	s := []string{"a", "b", "b", "b", "c", "b", "d"}

	fmt.Println("In: ", s)
	s = EliminateDups(s)
	fmt.Println("Out:", s)
}
