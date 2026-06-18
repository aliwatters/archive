package main

import "fmt"

func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // sq of sq! 16, 81
	}

	// set up a pipeline
	//	c := gen(2, 3)
	//	out := sq(c)

	// Consume output
	//	fmt.Println(<-out) // 4
	//	fmt.Println(<-out) // 9
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
