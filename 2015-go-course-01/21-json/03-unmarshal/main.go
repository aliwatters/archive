package main

import (
	"encoding/json"
	"fmt"
)

// note: only upper cased fields are exported - lowercased will not appear in json.
type Person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"years"`
}

func main() {

	bs := []byte(`{"first":"Jane","last":"Doe","years":26, "foo":"bar"}`) // note foo is ignored

	var p1 Person
	json.Unmarshal(bs, &p1) // p1 is a pointer to a person - read into the address
	fmt.Println(p1)

}
