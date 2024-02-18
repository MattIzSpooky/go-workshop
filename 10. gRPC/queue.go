package queue

import "sync"

type Queue[T any] struct {
	items []T
	lock  sync.Mutex
}

func (q *Queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)

}

func (q *Queue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	x := q.items[0]

	// Discard top element
	q.items = q.items[1:]

	return x
}

func (q *Queue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) > 0
}
