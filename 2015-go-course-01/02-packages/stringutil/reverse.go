package stringutil

// Reverse returns a reversed string
func Reverse(s string) string {
	// uses a unexported function
	return reverseTwo(s)
}
