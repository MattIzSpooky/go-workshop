package algorithms

import "slices"

func (l *List[T]) binarySearch(id int) *T {
	minId := 0
	maxId := len(l.items) - 1

	if maxId < 1 {
		return l.items[0]
	}

	for minId <= maxId {
		mid := (minId + maxId) / 2

		if id == (*l.items[mid]).ID() {
			return l.items[mid]
		}

		if (*l.items[mid]).ID() > id {
			maxId = mid - 1
		} else {
			minId = mid + 1
		}
	}

	return nil
}

func (l *List[T]) binaryAdd(item *T) {
	minIndex := 0
	maxIndex := len(l.items) - 1

	if maxIndex < 0 {
		l.items = append(l.items, item)
		return
	}

	for minIndex <= maxIndex {
		mid := (minIndex + maxIndex) / 2

		listItemId := (*l.items[mid]).ID()
		newItemId := (*item).ID()

		if (listItemId - newItemId) > 0 {
			maxIndex = mid - 1
		} else if (listItemId - newItemId) < 0 {
			minIndex = mid + 1
		}
	}

	var i int
	if maxIndex < minIndex {
		i = minIndex
	} else {
		i = (maxIndex + minIndex) / 2
	}

	if i == -1 {
		i = 0
	}

	l.items = slices.Insert(l.items, i, item)
}
