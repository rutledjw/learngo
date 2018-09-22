package testlibs

import (
	"errors"
	"fmt"
)

//ArgError is an error struct
type ArgError struct {
	arg  int
	prob string
}

//ProperError throws a Go-native exception
func ProperError(arg int) (int, error) {
	if arg == 42 {
		fmt.Println("can't work with 42 any longer")
		return -1, errors.New("42 Error")
	}

	return arg * 3, nil
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//F42Err is a simple error test function
func F42Err(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "can't work with 42 any more"}
	}
	return arg + 3, nil
}
