package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{33, 44, 33, 1, 2, 33, 2, 87, 54, 12, 90, 44, 3, 544, 2}

	fmt.Println(n)
	// sort.Sort(sort.IntSlice(n)) // implements the interface for slices of ints
	// sort.Sort(sort.Reverse(sort.IntSlice(n)))
	sort.Ints(n)
	fmt.Println(n)
}
