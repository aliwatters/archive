package main

// this is FILE level scoped - needs to be imported in each file.
import (
	"fmt"

	"github.com/aliwatters/go-course-01/04-scope/bar"
)

// demo of scope in golang - universe,  package, file and block.
// scope is inferrred from declaration position

// package level scope - outside of a block.
var x int = 42

func main() {
	fmt.Println(x)
	foo()

	y := 17
	fmt.Println(y)

	// variable shadowing - BAD practice
	max := max(7)
	fmt.Println(max)
	// fmt.Println(max(100)) // at this point - can't be done - the variable as shadowed the func
	msg := bar.You("Ali")
	fmt.Println(msg)

	increment := bar.Wrapper()
	fmt.Printf("Increment call 1: %d\n", increment())
	fmt.Printf("Increment call 2: %d\n", increment())
}

func foo() {
	fmt.Println(x)
	// fmt.Println(y) // - y not in scope

}

func max(x int) int {
	counter := 0

	increment := func() int {
		counter++
		return counter
	}

	return increment() + x
}
