package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)

	// Fan out

	c1 := sq(in)
	c2 := sq(in)

	// Fan in

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("Type of nums: %T\n", nums)

	out := make(chan int)
	go func() {
		for _, n := range nums { // range over a slice idx, val
			out <- n
		}
		close(out)
	}()
	return out
}

var workers int // global to allow counting of worker processes

func sq(in chan int) chan int {
	fmt.Printf("Type of in: %T\n", in)

	workers++
	out := make(chan int)
	go func(w int) {
		for n := range in {
			fmt.Println("Worker:", w, "Picked up", n)
			out <- n * n
		}
		close(out)
	}(workers)
	return out
}

// merge channels - does the fan in!
func merge(cs ...chan int) chan int {
	fmt.Printf("Type of CS: %T\n", cs)

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
