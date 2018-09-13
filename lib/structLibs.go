package testlibs

type rectangle struct {
	name          string
	height, width int
}

func (p *rectangle) area() int {
	return p.height * p.width
}

func (p *rectangle) half() (int, int) {
	return (p.height / 2), (p.width / 2)
}

//ClosTest tests Closure functions, functions that return anonymous functions
func ClosTest(x, y int) func() int {

	return func() int {
		oddSum := 0
		for x <= y {
			oddSum += y
			x++
		}
		return oddSum
	}
}

//fact performs factoral computation
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
