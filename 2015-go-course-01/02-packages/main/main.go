package main

import (
	"fmt"

	"github.com/aliwatters/go-course-01/02-packages/stringutil"
)

// Main function runs automatically - C like
func main() {
	fmt.Printf("Hello %s!\n", stringutil.MyName)

	var str = "Hello World!"
	fmt.Printf("%s\t%s\n", str, stringutil.Reverse(str))
}
