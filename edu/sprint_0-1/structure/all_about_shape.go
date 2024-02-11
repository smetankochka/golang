package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (self Circle) Area() float64 {
	return math.Pi * self.radius * self.radius
}

func (self Rectangle) Area() float64 {
	return self.height * self.width
}

func main() {
	circle := Circle{radius: 3.0}
	rectangle := Rectangle{width: 4.0, height: 5.0}
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area of the shape: %f\n", shape.Area())
	}
}
