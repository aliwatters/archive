package main

import (
	"encoding/json"
	"os"
)

// note: only upper cased fields are exported - lowercased will not appear in json.
type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	p1 := Person{"James", "Bond", 44}
	json.NewEncoder(os.Stdout).Encode(p1)
}
