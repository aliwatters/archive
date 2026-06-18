package main

import (
	"crypto/sha256"
	"fmt"
)

func countBitDiff(h1, h2 [32]byte) uint {
	out := uint(0)
	for i := range h1 {
		if h1[i] != h2[i] {
			out++
		}
	}
	return out
}

func main() {
	c1 := sha256.Sum256([]byte("hello world!"))
	c2 := sha256.Sum256([]byte("Hello world!"))

	fmt.Printf("%x\n%x\n%d of %d differ\n", c1, c2, countBitDiff(c1, c2), len(c1))
}
