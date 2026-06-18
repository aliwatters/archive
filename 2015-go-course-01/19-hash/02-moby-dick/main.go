package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func HashBucket(word string, buckets int) int {
	t := 0
	for _, v := range word {
		t += int(v)
	}
	return t % 12
}

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	check(err)
	defer res.Body.Close() // do that when finished

	words := make(map[string]bool)

	sc := bufio.NewScanner(res.Body) // strings.NewReader is a thing - just in case
	sc.Split(bufio.ScanWords)

	buckets := make([]int, 12) // slice of 12 things - could just be an array

	// sc.Scan returns a bool
	for sc.Scan() {
		words[sc.Text()] = true
		n := HashBucket(sc.Text(), 12)
		buckets[n]++
	}
	check(sc.Err())

	fmt.Println(buckets)
	fmt.Println(len(words), "unique words") // not trimmed lowercased etc
}
