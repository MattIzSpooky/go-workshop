package algorithms

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
