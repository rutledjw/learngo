package main

import (
	arr "../learngo/lib"
	"fmt"
	"math"
	"strings"
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
	fmt.Println("My favorite number is: ", rand.Intn(50))
	fmt.Println("----  space   ----")
	fmt.Printf("you now have %g problems", math.Sqrt(7))
	fmt.Println()
	fmt.Println("Pi: ", math.Pi, "add func: ", add(42, 45))
	fmt.Println()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Printf("Weird split function: ")
	fmt.Println(strings.split(17))
}
