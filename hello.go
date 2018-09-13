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
}
