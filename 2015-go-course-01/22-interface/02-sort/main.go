package main

import (
	"fmt"
	"sort"
)

type people []string

// next thrree funcs are required to implement the sort interface
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	p := people{"Zeno", "Harry", "Pete", "Jo", "Adam", "Sally"}
	fmt.Println(p)
	sort.Sort(p)
	// sort.Sort(sort.Reverse(p))
	fmt.Println(p)
}
