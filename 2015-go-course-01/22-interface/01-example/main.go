package main

import (
	"fmt"
	"math"
)

// Square - example shape 1
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

// Circle - example shape 2
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// Shape - interface for our examples
type Shape interface {
	area() float64
}

func info(z Shape) string {
	s := fmt.Sprintf("The area of this shape is: %f", z.area())
	return s
}

func main() {
	c := Circle{5}
	s := Square{10}
	fmt.Println("Circle:", info(c))
	fmt.Println("Square:", info(s))
}
