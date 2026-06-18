package main

import "fmt"

const (
	_ = iota // 0 - _ discarded
	// KB kilobytes
	KB = 1 << (iota * 10) // iota === 1
	// MB megabytes
	MB = 1 << (iota * 10) // iota === 2
	// GB gigbytes
	GB = 1 << (iota * 10) // iota === 3
	// TB terabytes
	TB = 1 << (iota * 10) // iota === 4
)

func main() {
	fmt.Println("Binary\tDecimal")
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)
}
