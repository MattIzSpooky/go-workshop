package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("cannot divide x by zero")

func main() {
	// Errors in Go do not work like Exceptions in Java, C# or C++
	// Instead, errors are treated like objects and can be wrapped with other errors
	testErr := errors.New("kaput!")

	wrappedErrors := errors.Join(testErr, errors.New("now there are 2 of them!"))

	fmt.Println(wrappedErrors.Error())

	// Functions do not throw errors, they return errors, a good example of this is a function we defined. Divide.

	// You can also ignore returned values by using an _. This also works for ignoring errors.
	// Whether that's a good idea I leave up to you.
	_, err := divide(1, 0)

	// This is how you check if an error is of a specific type.
	if errors.Is(err, DivideByZeroError) {
		fmt.Println(fmt.Sprintf("error!: %s\n", err.Error()))
	}

	// For unrecoverable errors you can also choose to panic. This will exit the program and dump a trace.
	panic("darn I panicked.")
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}
	return x / y, nil
}
