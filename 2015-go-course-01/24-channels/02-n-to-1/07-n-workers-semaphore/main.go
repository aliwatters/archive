package main

import (
	"fmt"
	"runtime"
)

func main() {

	n := runtime.NumCPU() // for CPU intensive things - this would be a sensible num
	c := make(chan int)
	done := make(chan bool)

	// launch our go routines!
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	// workers
	for i := 0; i < n; i++ {
		go func(j int) {
			for n := range c {
				fmt.Println("worker", j, n)
			}
			done <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<-done // waiting for n of these
	}

	close(done)
}
