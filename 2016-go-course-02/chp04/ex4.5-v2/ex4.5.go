package main

import "fmt"

// EliminateDups removes duplicate adjacent strings in a slice of strings
func EliminateDups(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	out := strings[:1] // slice of the underlying array (subtle!)
	a := out[0]
	for _, s := range strings {
		if s != a {
			out = append(out, s)
			a = s
		}
	}

	return out
}

func main() {
	// s := []string{"b", "b", "b"}
	s := []string{"a", "b", "b", "b", "c", "b", "d"}

	fmt.Println("In: ", s)
	s = EliminateDups(s)
	fmt.Println("Out:", s)
}
