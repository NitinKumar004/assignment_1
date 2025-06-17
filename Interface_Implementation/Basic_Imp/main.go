package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 10, height: 5}
	fmt.Println("Circle Area is ", c.Area())
	fmt.Println("Rectangle Area is ", r.Area())

}
