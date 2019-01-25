package testlibs

import "fmt"

//thread is a dumb function for calling threads
func thread(source string) {
	fmt.Println("called from:", source)
	fmt.Println("--------------")
}

//TestThreads is where we do this, no need to export above
func TestThreads() {

	var inputString string

	thread("direct")

	go thread("goroutine")

	//anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("go routine")

	fmt.Scanln(&inputString) // take input
	fmt.Println("Scanln() done...", inputString)
}

//RoutineTest is a loop function
func RoutineTest(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from+" : ", i)
	}
}
