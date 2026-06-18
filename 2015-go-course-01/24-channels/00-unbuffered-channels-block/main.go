package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // setup a channel

	go func() {
		// anonymous func a - sends to a channel c <- val
		for i := 0; i < 10; i++ {
			c <- i // send i to the channel - note code stops here until something picks up the value
		}
	}() // and execute

	go func() {
		// anonymous func b - acts as a reciever for the channel val <- c
		for {
			fmt.Println(<-c) // here values are picked up from the channel - code stops until something else is back on!
		}
	}()

	time.Sleep(time.Second)
}
