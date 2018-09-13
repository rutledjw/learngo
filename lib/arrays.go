package testlibs

import (
	"fmt"
)

//PointerTest - getting some experiece with Go pointers
func PointerTest() {
	i := 42
	p := &i
	fmt.Println("pointer to i ", *p)
}

// PrintSlice test our slice functions
func PrintSlice(i []string) {
	fmt.Printf("length=%d capacity=%d value is %v \n", len(i), len(i), i)
}

// RangeTest in another function
func RangeTest(input []string) {
	//var result []string
	for k, v := range input {
		fmt.Println("String values, k = ", k, " and v = ", v)
	}
}

// LoopTest is a function
func LoopTest(number int) {
	for j := 0; j < number; j++ {
		fmt.Println("And were back in 3rd grade - ", j)
	}
}

//BuildResponse blah
func BuildResponse(input string) string {
	return ("Hi there, " + input)
}

//RetMult test out multiple return values and have an example
func RetMult(a, b int) (int, int) {
	c := a * b
	return c, b
}

//VarTest tests variadic functions
func VarTest(nums ...int) {
	sum, count := 0, 0

	for _, num := range nums {
		sum += num
		count++
	}

	fmt.Printf("Variadic numbers: ")
	fmt.Println(nums)
	fmt.Printf("The sum of %d numbers is: %d\n", count, sum)
}
