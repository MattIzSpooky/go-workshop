## 6. Error handling

Go does not handle errors in the way C#, Java, JavaScript, etc. does. 
Throwing an error does not exist in Go, however, errors can be returned by functions instead.

Functions can return a value and an error, example:
```go
var DivideByZeroError = errors.New("cannot divide x by zero")

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}
	return x / y, nil
}
```

Here you can see that "divide" can return a value and an error. If you try to divide by `0` it will return an error.
When you divide by a value that is not `0` it will divide `x` and `y` and return `nil` for an error.

See [main.go](main.go) for a full example.

Why did the developers of Go design it this way? Read https://go.dev/doc/faq#exceptions