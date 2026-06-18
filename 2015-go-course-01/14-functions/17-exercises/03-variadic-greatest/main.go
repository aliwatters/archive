package main

import "fmt"

func greatest(x ...int) int {
	var max int
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(greatest(6, 3, 2, 5, 6, 7, 22, 3, 1, 3, 4, -6, -8, -100, 99, 10, 2, 3))
}
