package apma

func placement(capacity, count int) (place func(arr []int64, index int)) {
	// if value is less than zero, then we place the zeros in array
	// if zero is more than value, then we place the values in array
	placingZero := false
	if count > (capacity+1)/2 {
		placingZero = true
		count = capacity - count
	}
	spacing := capacity / count

	return func(arr []int64, index int) {
	}
}

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

	isWriteIndex := checkWriteIndex(right-left+1, count)
	for i := left; i <= right; i++ {
	}
}
