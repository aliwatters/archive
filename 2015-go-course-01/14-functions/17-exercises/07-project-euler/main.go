package main

import (
	"fmt"
	"math/big"
)

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

// https://golang.org/pkg/math/big - basics from here!
func main() {

	fmt.Print("Fibonacci: How many digits do you want to find? : ")
	var x int64
	fmt.Scan(&x)

	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(x-1), nil)

	fmt.Println("\nFinding a", x, "digit number")

	idx := 0
	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		idx++
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a // this is a neat trick - swap in place!
	}

	if x < 10000 {
		fmt.Println("Looking for a number larger than")
		fmt.Println(&limit) // get the value at the pointer
		fmt.Println(a)      // digit Fibonacci number
	}
	fmt.Println("found at position", idx)

	// it's fast! -- 1 million digit took a minute or so! - 4784969
	// $ time ./main
	// Fibonacci: How many digits do you want to find? : 100000
	//
	// Finding a 100000 digit number
	// found at position 478495

	// real	0m4.476s
	// user	0m0.700s
	// sys	0m0.004s

}
