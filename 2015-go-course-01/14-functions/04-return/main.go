package main

import "fmt"

func main() {
	a := "Ali"
	b := "Watters"
	fmt.Println(greet1(a, b))
	fmt.Println(greet2(a, b))
	fmt.Println(greet3(a, b))
}

// simple return
func greet1(fname, lname string) string {
	return fmt.Sprint(fname, lname) // returns a single string
}

// named return - not liking it
func greet2(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return // s is returned due to the named return above...
}

// multiple returns
func greet3(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
