package main

import "fmt"

const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

func main() {
	var powers = [...]float64{KB, MB, GB, TB, PB, EB, ZB, YB}
	for i, v := range powers {
		fmt.Printf("%d = %f\n", i, v)
	}
}
