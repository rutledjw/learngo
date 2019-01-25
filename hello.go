/*
*  To run (since you've forgotten)
*  > go build hello.go
*  > ./hello
 */

package main

import (
	"fmt"

	arr "../learngo/lib"
)

func main() {

	fmt.Println("Factoral of 5:", arr.Fact(5))

	fmt.Printf("Multiple Return: ")
	fmt.Println(arr.RetMult(3, 4))
	arr.VarTest(3, 6, 33, 894)
	cTest := arr.ClosTest(2, 6)
	fmt.Println("ClosureTest 2, 6: ", cTest())

	// Interfaces
	r := arr.Rect{Name: "Rectangle1", Width: 3, Height: 6.2}
	c := arr.Circle{Radius: 5.5}
	arr.Measure(r)
	arr.Measure(c)

	//Threads
	arr.TestThreads()
	if i, e := arr.ProperError(42); e != nil {
		fmt.Println("ProperError thrown", e, "and I =", i)
	}

	//Goroutines
	arr.RoutineTest("Direct Call")
	go arr.RoutineTest("Go routine")
	go func(message string) {
		fmt.Println("Go anon routine" + message)
	}("anon test")

}
