// Boilng prints the boiling point of water. (at sea level)
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g 'f or %g 'c\n", f, c)
}
