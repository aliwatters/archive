package main

import "fmt"

func main() {

	records := make([][]string, 0)

	s1 := make([]string, 4)
	s1[0] = "Doe"
	s1[1] = "John"
	s1[2] = "100.00"
	s1[3] = "64.00"

	records = append(records, s1)

	s2 := make([]string, 4)
	s2[0] = "Blogs"
	s2[1] = "Joe"
	s2[2] = "70.00"
	s2[3] = "11.00"

	records = append(records, s2)

	fmt.Println(records)
}
