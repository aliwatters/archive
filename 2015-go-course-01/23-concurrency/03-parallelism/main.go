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

func main() {
	wg.Add(2)
	go foo("foo:")
	go foo("bar:")
	wg.Wait()
}

func foo(msg string) {
	for i := 0; i < 10000; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}
