package main

import "fmt"

func main() {
	{ // ex1
		c := make(chan int)

		// c <- 1 // causes deadlock

		go func() {
			c <- 1
		}()

		fmt.Println(<-c)
		close(c)
	}
	{ // ex2
		c := make(chan int)

		// c <- 1 // causes deadlock

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()

		for n := range c {
			fmt.Println(n)
		}
	}

}
