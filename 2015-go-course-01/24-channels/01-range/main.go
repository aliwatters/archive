package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // setup a channel

	go func() {
		// anonymous func a - sends to a channel c <- val
		for i := 0; i < 10; i++ {
			c <- i // send i to the channel - note code stops here until something picks up the value
		}
		close(c)
	}() // and execute

	for n := range c { // here values are picked up from the channel
		fmt.Println(n)
	}

	// range will run until c is closed AND empty - no longer need a hacky delay
}
