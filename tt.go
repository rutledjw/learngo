//Comment so I can do a push
package main

import (
	"fmt"
	"math"
	"math/rand"

	arr "../learngo/lib"
)

func listRuns() {
	fmt.Println(arr.BuildResponse("Jeff"))

	var i, j int = 1, 2
	var c, java, python = true, false, "no!"
	fmt.Println(i, j, c, java, python)

	list := []string{"a", "c", "d", "e"}
	list = list[2:3]
	fmt.Println("slice size 3")
	arr.PrintSlice(list)
	arr.RangeTest(list)

	arr.LoopTest(5)

}

func oddMath() {
	fmt.Println("My favorite number is: ", rand.Intn(50))
	fmt.Println("----  space   ----")
	fmt.Printf("you now have %g problems", math.Sqrt(7))
	fmt.Println()
	fmt.Println("Pi: ", math.Pi)
	fmt.Println()
	a, b := arr.Swap("hello", "world")
	fmt.Println(a, b)
	fmt.Printf("Weird split function: ")
	fmt.Println(arr.Split(17))
}

/* I don't remember what this stuff does

 */
