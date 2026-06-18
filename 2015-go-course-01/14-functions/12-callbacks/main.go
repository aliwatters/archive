package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{} // empty slice

	fmt.Printf("%T", callback)
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
			// xs.push(n) -- can we push on a slice? doesn't look like it.
		}
	}

	return xs
}

func main() {
	nums := []int{1, 2, 3, 4}

	// callback - passing a function as an argument
	visit(nums, func(n int) {
		fmt.Println(n)
	})

	xs := filter(nums, func(n int) bool {
		return n > 2
	})

	fmt.Println(xs)
}
