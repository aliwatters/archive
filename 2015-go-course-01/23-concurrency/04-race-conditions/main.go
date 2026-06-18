package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	// init runs before anything else - setup
	// before 1.5 default was 1 core - 1.5 default is all.
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Max CPUs:", runtime.NumCPU())
}

var counter int

func main() {
	wg.Add(2)
	go foo("foo:")
	go foo("bar:")
	wg.Wait()
}

func foo(msg string) {
	for i := 0; i < 100; i++ {
		// x := counter // artificial example - counter to 200?
		fmt.Println(msg, i, counter)
		time.Sleep(time.Duration(5 * time.Millisecond))
		counter++ // = x + 1 // this is where a race condition is caused...
		// interesting - -race flag detects this as a race condition!
	}
	wg.Done()
}
