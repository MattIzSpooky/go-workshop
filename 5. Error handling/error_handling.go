package main

import "errors"

func main() {
	// Errors in Go do not work like Exceptions in Java, C# or C++
	// Instead, errors are treated like objects and can be wrapped with other errors
	testErr := errors.New("kaput!")

	errors.Join(errors.New("now there are 2 of them!"))
	testErr.Error()
}
