package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	// slice mySlice from position 2 -> 4 (not including 4)
	fmt.Println(mySlice[2:4])

	// bring back value at index 2
	fmt.Println(mySlice[2])

	// string are slices of runes (bytes)!
	// so slice operations work
	fmt.Println("myString"[2:6])

}
