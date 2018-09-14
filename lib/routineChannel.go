package testlibs

import "fmt"

//thread is a dumb function for calling threads
func thread(source string) {
	fmt.Println("called from:", source)
	fmt.Println("--------------")
}

//TestThreads is where we do this, no need to export above
func TestThreads() {
	thread("direct")

	go thread("goroutine")

	//anonymous call
	go func(msg string) {
		fmt.Println(msg)
	}("go routine")

	fmt.Scanln() // take input
	fmt.Println("done...")
}