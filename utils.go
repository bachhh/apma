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
	num, dem := capacity, c // capacity / count
	for i := right; i >= left; i-- {
	}

	return arr
}
