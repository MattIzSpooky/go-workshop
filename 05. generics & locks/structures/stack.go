package structures

import (
	"fmt"
	"sync"
)

type Stack[T any] struct {
	items []T
	lock  sync.Mutex
}

func (s *Stack[T]) Push(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	fmt.Println(fmt.Sprintf("Push: [%v]", item))
	s.items = append([]T{item}, s.items...)
}

func (s *Stack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	item := s.items[0]
	fmt.Println(fmt.Sprintf("Pop: [%v]", item))
	s.items = s.items[1:]

	return item
}

func (s *Stack[T]) Top() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.items[0]
}

func (s *Stack[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items) == 0
}
