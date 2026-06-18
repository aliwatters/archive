package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int32 // prepare for atomic

func main() {
	wg.Add(2)
	go foo("foo:")
	go foo("bar:")
	wg.Wait()
}

func foo(msg string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))

		atomic.AddInt32(&counter, 1) // faster than mutex - but only very simple ops
		fmt.Println(msg, i, counter)
	}
	wg.Done()
}
