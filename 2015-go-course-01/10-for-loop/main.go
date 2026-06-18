package main

import "fmt"

func main() {

	a := []string{"g", "f", "e", "d", "c", "b", "a"}

	fmt.Println("Array:", a)

	// scope of i and j are BLOCK level
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("Array:", a)

	// while version
	i := 0
	for i < 10 {
		fmt.Printf("%d, ", i)
		i++
	}
	fmt.Print("\n")

	// forever - break and continue
	i = 0
	for {
		i++

		// print odd numbers using continue
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d, ", i)
		if i > 50 {
			break // out of the loop
		}
	}
	fmt.Println("\nFinished")
}
