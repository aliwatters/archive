package main

import "fmt"

// The Fibonacci sequence is defined by the recurrence relation:
//
// Fn = Fn?1 + Fn?2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:
//
// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
// The 12th term, F12, is the first term to contain three digits.
//
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

func fibonacci(digitCount int) int {
	// digit count - when we have this number of digits - return

	i := 1
	j := 1

	for j/10 < digitCount {
		k := i + j
		i = j
		j = k
	}
	return j

}

func main() {
	fmt.Println("running")
	test := fibonacci(3)
	fmt.Println(test)

	// unfortunately - int64 isn't large enough to solve 1000 digits
	real := fibonacci(10)
	fmt.Println(real)
}
