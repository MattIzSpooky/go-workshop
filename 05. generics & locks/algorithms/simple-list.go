package algorithms

import (
	"fmt"
	"math/rand"
)

type IDer interface {
	ID() int
}

type List[T IDer] struct {
	items []*T
}

func (l *List[T]) Add(item *T) {
	l.binaryAdd(item)
}

func (l *List[T]) Sort() {
	leftIndex := 0
	rightIndex := len(l.items) - 1

	if rightIndex < 1 {
		return // No need to sort.
	}

	l.quicksort(leftIndex, rightIndex)
}

func (l *List[T]) Search(id int) *T {
	return l.binarySearch(id)
}

func (l *List[T]) Print() {
	for i, item := range l.items {
		// Fun fact: Golang supports emojis as it its fully UTF-8.
		fmt.Println(fmt.Sprintf("Position: [%d] ➔️ %v", i, *item))
	}
}
func (l *List[T]) Randomize() {
	rand.Shuffle(len(l.items), func(i, j int) { l.items[i], l.items[j] = l.items[j], l.items[i] })
}
