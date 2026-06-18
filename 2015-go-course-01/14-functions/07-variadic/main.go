package main

import "fmt"

func main() {
	// demo that a variadic function can recieve unlimited args
	fmt.Println(greet("Ali", "Michelle", "Olivia"))

	// demo that we can pass a slice into a variadic function
	data := []float64{43, 56.8, 3, 0.001}

	// the slice is "spread" out into args
	// average(data) would pass arg[0] as a slice - slightly different
	fmt.Println(average(data...))
	// average2 demos that
	fmt.Println(average2(data))
}

func greet(opt ...string) string {
	s := ""
	for _, v := range opt {
		s += v
	}
	return s
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func average2(sf []float64) float64 {
	// spread the sf
	return average(sf...)
}
