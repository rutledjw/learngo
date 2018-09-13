package main

import (
	arr "../learngo/lib"
	"fmt"
)

func main() {

	fmt.Println("Factoral of 5:", arr.Fact(5))

	fmt.Printf("Multiple Return: ")
	fmt.Println(arr.RetMult(3, 4))
	arr.VarTest(3, 6, 33, 894)
	cTest := arr.ClosTest(2, 6)
	fmt.Println("ClosureTest 2, 6: ", cTest())

	// Interfaces
	r := arr.Rect{Width: 3, Height: 6.2}
	c := arr.Circle{Radius: 5.5}

	arr.Measure(r)
	arr.Measure(c)
}
