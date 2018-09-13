package testlibs

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// Circle is a shape
type Circle struct {
	Radius float64
}

//Rect geometric features
type Rect struct {
	name          string
	Height, Width float64
}

func (r Rect) area() float64 {
	return r.Height * r.Width
}

func (r Rect) perim() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}

//Measure is exported
func Measure(g geometry) {
	fmt.Println("Geometry", g)
	fmt.Println("Area", g.area())
	fmt.Println("Perimeter", g.perim())
}

//ClosTest tests Closure functions, functions that return anonymous functions
func ClosTest(x, y float64) func() float64 {

	return func() float64 {
		oddSum := 0.0
		for x <= y {
			oddSum += y
			x++
		}
		return oddSum
	}
}

//Fact performs factoral computation
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
