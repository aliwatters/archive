package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println([]byte("hello"))
	fmt.Println(string([]byte("hello")))
	{
		i, err := strconv.Atoi("42")
		fmt.Printf("%T %v %v\n", i, i, err)
	}
	{
		i := strconv.Itoa(42)
		fmt.Printf("%T %v\n", i, i)
	}
	{
		i, err := strconv.ParseInt("-42", 10, 64) // 64 bit!
		fmt.Printf("%T %v %v\n", i, i, err)
	}
}
