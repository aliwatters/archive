package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// channel based solution - in theory faster as takes advantage of all CPUs

	cpus := runtime.NumCPU() // 2 CPUs on this system

	// note factorial quickly exceeds int64 - really should be done with math/big
	var max int
	fmt.Printf("Detected %d CPU, Factorial of? ", cpus)
	fmt.Scan(&max)

	c := make(chan int)
	done := make(chan int)

	go func() {
		// setup the factorial calculation
		for i := max; i > 0; i-- {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < cpus; i++ {
		go func(j int) {
			fmt.Println("Setup worker", j)
			total := 1
			for n := range c {
				fmt.Println("worker", j, "picked up", n)
				time.Sleep(200 * time.Millisecond) // for demo
				total *= n
			}
			done <- total
		}(i)
	}

	total := 1
	for i := 0; i < cpus; i++ {
		x := <-done
		fmt.Println("interim total", i, x)
		total *= x // last multiplications in series
	}

	fmt.Println("Factorial of ", max, " is:", total)
	close(done)
}
