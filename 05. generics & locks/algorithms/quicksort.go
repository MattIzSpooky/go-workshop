package algorithms

func (l *List[T]) quicksort(leftIndex int, rightIndex int) {

	if leftIndex >= rightIndex {
		return
	}

	// Take median of indexes, rounded to lower
	pivotIndex := (leftIndex + rightIndex) / 2

	pivot := l.partition(pivotIndex, leftIndex, rightIndex)

	if pivot > 1 {
		l.quicksort(leftIndex, pivot-1)
	}
	if pivot+1 < rightIndex {
		l.quicksort(pivot+1, rightIndex)
	}
}

/**
 * Compare two values at index i1 and i2 with each-other
 * Returns a negative integer if the value at the first index is smaller than the value at the second index
 * Returns 0 if the value at the first index is equal to the value at the second index
 * Returns a positive integer if the value at the first index is greater that the value at the second index
 */
func (l *List[T]) compare(i1 int, i2 int) int {
	return (*l.items[i1]).ID() - (*l.items[i2]).ID()
}

func (l *List[T]) partition(pivotIndex int, leftIndex int, rightIndex int) int {
	for leftIndex < rightIndex {
		for l.compare(leftIndex, pivotIndex) < 0 {
			leftIndex++
		}

		for l.compare(rightIndex, pivotIndex) > 0 {
			rightIndex--
		}

		if pivotIndex == leftIndex {
			pivotIndex = rightIndex
		} else if pivotIndex == rightIndex {
			pivotIndex = leftIndex
		}
		l.swap(leftIndex, rightIndex)
	}

	return pivotIndex
}

/**
 * Swap the values at index i1 and i2 around
 */
func (l *List[T]) swap(i1 int, i2 int) {
	temp := l.items[i1]
	l.items[i1] = l.items[i2]
	l.items[i2] = temp
}
