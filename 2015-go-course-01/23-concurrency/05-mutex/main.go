package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go foo("foo:")
	go foo("bar:")
	wg.Wait()
}

func foo(msg string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))
		fmt.Println(msg, i, counter)

		mutex.Lock()
		counter++
		mutex.Unlock()
		// BUT whatever is between lock and unlock delays - so make it quick and minimal :)
	}
	wg.Done()
}
