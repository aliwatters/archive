package main

import "fmt"

func main() {

	// examples of array

	var x [58]string // array of string 58 items

	// shows a empty string for each item...
	fmt.Println(x)

	// 16-array/main.go:11: invalid array index 60 (out of bounds for 58-element array)
	// x[60] = "test"

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)

	var y [10]int
	fmt.Println(y) // note item inited to 0

	const size = 5 // note a var does not work.
	var z [size]int
	fmt.Println(z)

	// loops through the array
	var x2 [256]int
	// we can get the length of x2 - so use it in the loop
	for i, l := 0, len(x2); i < l; i++ {
		x2[i] = i
	}
	// range also works - i is the index, v value
	for i, v := range x2 {
		fmt.Printf("%d:\t%T %v\t%b\t%x\n", i, v, v, v, v)
	}

	// shows that a byte is an alias of uint8
	var x3 [256]byte
	for i := 0; i < 256; i++ {
		x3[i] = byte(i)
	}
	for i, v := range x3 {
		fmt.Printf("%d:\t%T %v\t%b\t%x\n", i, v, v, v, v)
	}

}
