package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	f, err := os.Open("/usr/share/dict/words")
	check(err)
	defer f.Close() // do that when finished

	words := make(map[string]string)

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input", err)
	}

	for k, _ := range words {
		fmt.Println(k)
	}

}
