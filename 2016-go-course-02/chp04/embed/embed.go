// example of struct embedding
package main

import "fmt"

// Point x, y coord
type Point struct {
	X int
	Y int
}

// Circle a point with a radius
type Circle struct {
	Point
	Radius int
}

// Wheel a circle with spokes
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}
