package main

import "fmt"

func main() {
	// make a map where the key is string, val is int
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Printf("Type %T\n", m)
	fmt.Println("len:", len(m))

	{
		// key exists check - throw away value
		_, ok := m["k2"]
		fmt.Println("ok:", ok)
	}

	// delete the k2 key
	delete(m, "k2")

	{
		// k2 no longer exists - note: val is the zero value! - we need to use a test so that we know if values are set vs not inited.
		val, ok := m["k2"]
		fmt.Println("ok:", val, ok)
	}

	if val, ok := m["k2"]; ok { // keeps scope of ok to just this block!
		fmt.Println("Value:", val)
	} else {
		fmt.Println("k2 not set", val) // note has zero value
	}

	{
		// problems with this "no elements may be added" - why is this a thing?
		var m map[string]string
		fmt.Println(m)
		// m["key"] = "hello"
		// panic: assignment to entry in nil map
	}

	{
		// composite literal
		m := map[string]int{"k1": 1, "k2": 2}
		fmt.Println(m)
	}
}
