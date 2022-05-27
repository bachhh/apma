package apma

// Given an array of randomly placed "value", and "space" like this {1,0,0,0,2,3,4,0,0,6,7...}
//
// - "value" are elements are > 0
// - "space" are the zero elements
// - Assume the density (value_counts / capacity) is always
//    a rational factor ( 1/2, 1/3, 1/4... )
// - All values elements are sorted, ascending order.
//
// Our goal is to redistribute array so that all the values are evenly space:
//      - {1 0 0 2 0 0 3 0 0 4 0 0 6 0 0 7 0 0 } like this
//
// O(n) complexity, O(1) mem ?

// spread all elements in the range array[left, right] evenly spaced
// insert the x element at the appropriate place also
func diluteInsert(arr []int64, left, right int, x uint, count int) {
	// s is the "spacing" between each non-zero element

	capacity := right - left + 1
	spacing := capacity / count
	y := 0
	valueCounter := 0 // keep track of how many value before x

	isCorrectPlace := func(index int) bool {
		return index*count == spacing
	}

	x, y := count, len(arr)
	for x >= 0 {
		if y == spacing*x {
			swap(arr[x], arr[y])
			x--
		}
		y--
	}

}
