package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var ds = flag.Int("ds", 256, "Digest size to use; [256, 384, 512]")

func init() {
	flag.Parse() // needed to actually put values in *ds
	switch *ds {
	case 256, 384, 512:
		// do nothing
	default:
		fmt.Println("Error: Invalid digest size")
		os.Exit(1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: reading stdin")
		os.Exit(1)
	}

	switch *ds {
	case 256:
		c := sha256.Sum256([]byte(text))
		fmt.Printf("%x\n", c)
	case 384:
		c := sha512.Sum384([]byte(text))
		fmt.Printf("%x\n", c)
	case 512:
		c := sha512.Sum512([]byte(text))
		fmt.Printf("%x\n", c)
	}
}
