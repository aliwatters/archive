package bar

// You demos use of block scoped vars returned and exported functions
func You(name string) string {
	message := ", you are a future " + what() + " programmer!"
	return name + message
}

// This function is not exported - lowercase name
func what() string {
	return "rock star"
}

// Wrapper here is designed to be an IIFE - called with wrapper()
func Wrapper() func() int {
	// the returned func is a closure of this scope (a ref to x is maintained)
	x := 0
	return func() int {
		x++
		return x
	}
}
