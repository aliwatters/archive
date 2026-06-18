package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// fanout pattern
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)

	for n := range merge(f1, f2, f3, f4, f5) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			out <- (i % 10) + 3 // put 3-12 on the channel 100 times
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

// merge channels - does the fan in!
func merge(cs ...<-chan int) chan int {
	fmt.Printf("Type of CS: %T\n", cs)

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
