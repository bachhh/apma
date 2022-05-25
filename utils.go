package apma

func isWriteIndex(index int, capacity, count int) bool {
	// if value is less than zero, then we place the zeros in array
	// if zero is more than value, then we place the values in array
	return true
}

// spread all elements in the range array[left, right] evenly spaced
// insert the x element at the appropriate place also
func diluteInsert(arr []int64, left, right int, x uint, count int) {
	// s is the "spacing" between each non-zero element

	r, w := left, left
	capacity := right - left + 1
	for w <= right {
		if isWriteIndex(r, capacity, count) { // CASE01
			r++
			continue
		}

		if arr[r] != EMPTY {
			arr[w], arr[r] = arr[r], arr[w]
		}

		// CASE02
		if arr[r] == EMPTY {
			r++
		}

		for { // do while
			w++
		}
	}

}
