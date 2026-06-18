package main

import "fmt"

func main() {
	greet := map[int]string{0: "Hello", 2: "Hola", 3: "Selemat Pagi", 4: "Szia"}

	for k, v := range greet {
		fmt.Println(k, " - ", v)
	}
}
