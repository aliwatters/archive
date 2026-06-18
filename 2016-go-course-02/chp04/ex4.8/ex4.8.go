package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// UnicodeStat is a structure to represent stats about unicdoe text
type UnicodeStat struct {
	Count    map[rune]int
	UtfLen   map[int]int
	Category map[string]int
	Invalid  int
}

// CharType returns a string representing the unicode type of a rune
func CharType(r rune) string {
	switch {
	case unicode.IsLetter(r):
		return "letter"
	case unicode.IsSpace(r):
		return "space"
	case unicode.IsPunct(r):
		return "punct"
	case unicode.IsNumber(r):
		return "number"
	case unicode.IsSymbol(r):
		return "symbol"
	case unicode.IsMark(r):
		return "mark"
	case unicode.IsDigit(r):
		return "digit"
	case unicode.IsPrint(r):
		return "print"
	case unicode.IsControl(r):
		return "control"
	case unicode.IsGraphic(r):
		return "graphic"
	default:
		return "invalid"
	}
}

// CharCount reads input and returns stats about unicode frequency
func CharCount(in io.RuneReader) UnicodeStat {
	var results UnicodeStat
	results.Count = make(map[rune]int)      // counts of Unicode characters
	results.UtfLen = make(map[int]int)      // count of lengths of UTF-8 encodings
	results.Category = make(map[string]int) // number of categorys like Unicode.IsNumber
	results.Invalid = 0                     // count of invalid UTF-8 characters

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			results.Invalid++
			continue
		}
		results.Count[r]++
		results.Category[CharType(r)]++
		results.UtfLen[n]++
	}

	return results
}

func main() {
	in := bufio.NewReader(os.Stdin)
	counts := CharCount(in)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts.Count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\ncategory\tcount\n")
	for t, n := range counts.Category {
		fmt.Printf("%s\t%d\n", t, n)
	}
	fmt.Printf("\nutflen\tcount\n")
	for i, n := range counts.UtfLen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if counts.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters", counts.Invalid)
	}
}
