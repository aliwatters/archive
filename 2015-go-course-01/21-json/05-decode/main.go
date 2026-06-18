package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// note: only upper cased fields are exported - lowercased will not appear in json.
type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"first":"James","last":"Bond","age":44}`)

	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1)
}
