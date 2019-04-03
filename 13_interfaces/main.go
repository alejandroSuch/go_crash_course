package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

type Circle struct {
	x, y, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	r := Rectangle{4, 3}
	c := Circle{5, 5, 12}

	fmt.Printf("Rectangle area: %f\n", getArea(r))
	fmt.Printf("Circle area: %f\n", getArea(c))
}
