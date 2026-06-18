package main

import (
	"encoding/json"
	"fmt"
)

// note: only upper cased fields are exported - lowercased will not appear in json.
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{"James", "Bond", 44}

	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

}
