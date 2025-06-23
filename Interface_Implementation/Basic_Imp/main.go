package main

import (
	"math"
)

// Interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

// Circle struct
type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func costt(s Shape) float64 {
	switch sh := s.(type) {
	case *Circle:
		return sh.radius * sh.Area()
	case *Rectangle:
		return 0.04 * sh.Area()
	default:
		return 0.0
	}
}

//func main() {
//	// Circle
//	//c := &Circle{radius: 5}
//	//fmt.Println("Circle cost:", costt(c))
//	//
//	//r := &Rectangle{width: 10, height: 5}
//	//fmt.Println("Rectangle cost:", costt(r))
//
//}
