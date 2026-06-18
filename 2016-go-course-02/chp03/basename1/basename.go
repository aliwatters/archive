package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
	}
	arg := os.Args[1]

	fmt.Printf("Basename1: %s\n", basename1(arg))
	fmt.Printf("Basename2: %s\n", basename2(arg))
}

// basename1 removes directory components and a .suffix
func basename1(s string) string {
	// discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before the last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// basename2 removes directory components and a .suffix
func basename2(s string) string {
	// discard last '/' and everything before
	slash := strings.LastIndex(s, "/") // -i if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
