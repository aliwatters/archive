package main

import "fmt"

func main() {
	// semaphore pattern
	c := make(chan int)
	done := make(chan bool) // this is our completion channel

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() { // RIGHT WAY
		<-done
		<-done
		close(c)
	}()

	// WRONG WAY -- <-done blocks so range isn't reached so c <- i is a deadlock
	// <-done
	// <-done
	// close(c)

	for n := range c {
		fmt.Println(n)
	}

	close(done)
}
