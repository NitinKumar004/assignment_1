package main

import (
	"math"
	"testing"
)

func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}

func TestRectangle(t *testing.T) {
	r := &Rectangle{width: 10, height: 5}
	var exp float64 = 2
	y := costt(r)
	if !withinTolerance(exp, y, 1e-12) {
		t.Errorf("Expected %.18f, got %.18f", exp, y)
	}
}
func TestCircle(t *testing.T) {
	r := &Circle{5}
	var exp float64 = 392
	y := costt(r)

	if 1e-2 < (exp - y) {
		t.Errorf("Expected %.18f, got %.18f", exp, y)
	}
	//if !withinTolerance(exp, y, 1e-12) {
	//	t.Errorf("Expected %.18f, got %.18f", exp, y)
	//}
}
