package main

import "fmt"

func main() {
	fmt.Println("Decimal: ", 42)
	fmt.Printf("Binary: %d - %b\n", 42, 42)
	fmt.Printf("Octal: %d - %#o\n", 42, 42) // %o -- without leading 0
	fmt.Printf("Hexidecimal: %d - %#x\n", 42, 42)
	fmt.Printf("Hexidecimal: %d - %#X\n", 42, 42)
	fmt.Printf("Hexidecimal: %d - %#X\n", 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d\t%b\t%#o\t%#x\t%q\n", i, i, i, i, i)
	}

	fmt.Println("Finished")
}
