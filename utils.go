package apma

func diluteInsert(arr []int64, left, right int, x int64, count int) []int64 {
	c := 0
	// pack
	for i := left; i <= right; i++ {
		if arr[i] != EMPTY {
			if arr[i] > x { // sneak in x
				arr[c] = x
			} else {
				arr[c], arr[i] = arr[i], arr[c]
			}
			c++
		}
	}
	// there should be at least one last slot to insert x
	if c <= right {
		arr[c] = x
	} else {
		panic("trying to insert into a full array")
	}

	capacity := right - left + 1
	spacing := capacity / c
	spaceZero := c+1 >= capacity*2
	c--

	// spread
	i := right
	for c >= 0 {
		// CASE01: each index i = spacing*k is a zero, so we writes
		// value in indices between them.
		// or
		// CASE02: each index i = spacing*k is a value, so we swap
		if (spaceZero && i%spacing != 0) ||
			(!spaceZero && i%spacing == 0) {
			arr[i], arr[c] = arr[c], EMPTY
			c--
		}
		i--
	}
	return arr
}
