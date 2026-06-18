package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// UnicodeSpaceToASCII removes duplicate adjacent strings in a slice of strings
func UnicodeSpaceToASCII(s []byte) []byte {
	out := s[:0] // 0 length slice of s - uses same underlying array
	lastIsSpace := false

	for start, i := 0, 0; i < len(s)+1; i++ {
		b := s[start:i]
		// fmt.Printf("\n%d\tTesting: % x", i, b)
		if utf8.Valid(b) {
			r, _ := utf8.DecodeRune(b)
			if unicode.IsSpace(r) {
				if lastIsSpace == false {
					out = append(out, 32) // ascii space
					lastIsSpace = true
				}
			} else {
				out = append(out, b...) // ... expands out the slice to values
				lastIsSpace = false
			}
			start = i
		}
	}

	return out
}

func main() {
	// samples - see tests for exahustive list
	// var bytes = []byte("[᠎ 日本語 ][語᠎日]") // "[᠎ 日本語 ][語 日] (hidden ws in second set)
	// var bytes = []byte("hello")
	var bytes = []byte("[\u0020\u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000]")

	fmt.Println("In: ", string(bytes))
	bytes = UnicodeSpaceToASCII(bytes)
	fmt.Println("Out:", string(bytes))
}
