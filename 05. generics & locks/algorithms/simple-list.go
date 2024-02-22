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
	// I could implement a binary-search like way to add an item, but I am being lazy here.
	l.items = append(l.items, item)
}

func (l *List[T]) Sort() {
	leftIndex := 0
	rightIndex := len(l.items)

	if rightIndex < 1 {
		return // No need to sort.
	}

	l.quicksort(leftIndex, rightIndex-1)
}

func (l *List[T]) Search(id int) *T {
	return l.binarySearch(id)
}

func (l *List[T]) Print() {
	for i, item := range l.items {
		fmt.Println(fmt.Sprintf("Position: [%d] -> %v", i, *item))
	}
}
func (l *List[T]) Randomize() {
	rand.Shuffle(len(l.items), func(i, j int) { l.items[i], l.items[j] = l.items[j], l.items[i] })
}
