package main

import (
	"errors"
	"golang.org/x/exp/constraints"
)

// The methods below also exist in the "slices" package, this is just for demo purposes.
func Min[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		return *new(T), errors.New("array is empty")
	}

	minimumValue := s[0]
	for _, v := range s {
		if v < minimumValue {
			minimumValue = v
		}
	}
	return minimumValue, nil
}

func Max[T constraints.Ordered](s []T) (T, error) {
	if len(s) == 0 {
		return *new(T), errors.New("array is empty")
	}

	maximumValue := s[0]
	for _, v := range s {
		if v > maximumValue {
			maximumValue = v
		}
	}
	return maximumValue, nil
}
