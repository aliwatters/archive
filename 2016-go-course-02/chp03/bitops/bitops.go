package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)  // intersection 00000010
	fmt.Printf("%08b\n", x|y)  // union 00100110
	fmt.Printf("%08b\n", x^y)  // 00100100 symetric difference
	fmt.Printf("%08b\n", x&^y) // 00100000 difference

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // "1" "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // 01000100 [2,6]
	fmt.Printf("%08b\n", x>>1) // 00010001 [0,4]

}
