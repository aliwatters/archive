package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	// launch our go routines!
	go func() {
		for i := 0; i < 10000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println("a", n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println("b", n)
		}
		done <- true
	}()

	// we have 2 dones to happen - so read and block twice
	<-done
	<-done

	close(done)
}
