package structures

import (
	"fmt"
	"sync"
)

// This is a simple queue that is written to be thread safe.
type Queue[T any] struct {
	items []T
	lock  sync.Mutex
}

func (q *Queue[T]) Enqueue(item T) {
	// When you call "lock". it will block until the lock comes available, so it can lock in the current call.
	// This guarantees that only 1 function can access the struct at a given time. It is how you can make your structs thread safe
	// if you need to pass them onto multiple go routines for some reason.
	q.lock.Lock()
	defer q.lock.Unlock() // When you lock, you also have to unlock because if you don't it'll keep blocking on the next lock call.
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
