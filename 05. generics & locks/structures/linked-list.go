package structures

import (
	"fmt"
	"sync"
)

type link[T any] struct {
	next    *link[T]
	content T
}

type LinkedList[T any] struct {
	head *link[T]
	lock sync.Mutex
}

func (l *LinkedList[T]) Insert(item T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	newLink := l.createLink(item)

	current := l.head

	if l.head == nil {
		l.head = newLink
	} else {
		for current.next != nil {
			current = current.next
		}
		current.next = newLink
	}
}

func (l *LinkedList[T]) Print() {
	l.lock.Lock()
	defer l.lock.Unlock()

	link := l.head

	for link != nil {
		if link.next == nil {
			fmt.Println(fmt.Sprintf("%v", link.content))
		} else {
			fmt.Printf("%v -> ", link.content)
		}

		link = link.next
	}
}

func (l *LinkedList[T]) FindFunc(f func(T) bool) T {
	l.lock.Lock()
	defer l.lock.Unlock()

	link := l.head

	// O(n)
	for link.next != nil {
		if f(link.content) {
			return link.content
		}

		link = link.next
	}

	return *new(T)
}

func (l *LinkedList[T]) createLink(item T) *link[T] {
	return &link[T]{
		next:    nil,
		content: item,
	}
}
