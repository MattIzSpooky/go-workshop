package structures

import (
	"fmt"
	"sync"
)

type Queue[T any] struct {
	items []T
	lock  sync.Mutex
}

func (q *Queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	fmt.Println(fmt.Sprintf("Enqueue: [%v]", item))
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	item := q.items[0]
	fmt.Println(fmt.Sprintf("Dequeue: [%v]", item))

	// Discard top element
	q.items = q.items[1:]

	return item
}

func (q *Queue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}
