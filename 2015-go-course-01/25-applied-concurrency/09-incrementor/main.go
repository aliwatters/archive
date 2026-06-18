package main

import "fmt"

func main() {

	c := incrementor(20)

	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for x := 0; x < 20; x++ {
				c <- fmt.Sprint("GoRoutine:", i, " Count:", x)
			}
			done <- true
		}(i)
	}

	// Nice pattern - semaphore - blocks until all DONE, then closes c - which unblocks in main()
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
