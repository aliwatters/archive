package main

import (
	"fmt"
)

func main() {
	var name interface{} = 42 // "Sydney"

	str, ok := name.(string) // assert that the underlying type is a string

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("Value is not a string is %T\n", name)
	}
}
