package structures

import "sync"

type Stack[T any] struct {
	items []T
	lock  sync.Mutex
}

func (s *Stack[T]) Push(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append([]T{item}, s.items...)
}

func (s *Stack[T]) Pop() T {
	s.lock.Lock()
	defer s.lock.Unlock()

	x := s.items[0]
	s.items = s.items[1:]

	return x
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
