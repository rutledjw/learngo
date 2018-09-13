package main

import (
	arr "../learngo/lib"
	"fmt"
)

func main() {

	fmt.Println("Factoral of 5:", arr.fact(5))

	fmt.Printf("Multiple Return: ")
	fmt.Println(arr.RetMult(3, 4))
	arr.VarTest(3, 6, 33, 894)
	cTest := arr.ClosTest(2, 6)
	fmt.Println("ClosureTest 2, 6: ", cTest())
}

/*	fmt.Println(arr.BuildResponse("Jeff"))

	var i, j int = 1, 2
	var c, java, python = true, false, "no!"
	fmt.Println(i, j, c, java, python)

	list := []string{"a", "c", "d", "e"}
	list = list[2:3]
	fmt.Println("slice size 3")
	arr.PrintSlice(list)
	arr.RangeTest(list)

	zz.LoopTest(5)
	fmt.Println("My favorite number is: ", rand.Intn(50))
	fmt.Println("----  space   ----")
	fmt.Printf("you now have %g problems", math.Sqrt(7))
	fmt.Println()
	fmt.Println("Pi: ", math.Pi, "add func: ", add(42, 45))
	fmt.Println()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Printf("Weird split function: ")
	fmt.Println(split(17))
}
func add(x, y int) int {
	return x + y
}
func swap(a, b string) (string, string) {
	return b, a
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} */
