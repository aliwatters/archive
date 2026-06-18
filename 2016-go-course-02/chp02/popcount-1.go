package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliwatters/go-course-02/chp02/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount1: %v\n", err)
			os.Exit(1)
		}

		count := popcount.Popcount1(n)

		fmt.Printf("%d popcount is: %d\n", n, count)
	}
}
