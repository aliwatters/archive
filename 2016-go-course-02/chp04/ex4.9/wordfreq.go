package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

func isNotLetter(r rune) bool {
	return !unicode.IsLetter(r)
}

// WordFreq reads input and returns stats about unicode frequency
func WordFreq(in io.Reader) map[string]int {
	var counts = make(map[string]int)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(strings.TrimFunc(scanner.Text(), isNotLetter))
		counts[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "WordFreq: %v\n", err)
		os.Exit(1)
	}

	return counts
}

func main() {
	in := bufio.NewReader(os.Stdin)
	counts := WordFreq(in)

	fmt.Printf("word\tfreq\n")

	// sort the words by key for display
	var keys []string
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s\t%d\n", key, counts[key])
	}
}
