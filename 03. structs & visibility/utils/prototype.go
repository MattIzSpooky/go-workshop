package utils

type Prototype[T any] interface {
	Clone() T
}
