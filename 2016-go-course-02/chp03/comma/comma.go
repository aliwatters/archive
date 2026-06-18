package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	arg := os.Args[1]

	fmt.Printf("Comma: %s\n", comma(arg))
}

// basename1 removes directory components and a .suffix
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
