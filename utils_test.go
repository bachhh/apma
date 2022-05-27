package apma

import (
	"fmt"
	"testing"
)

func TestValidateArray(t *testing.T) {
	testCase := map[string][]int64{
		"0.3 head concentrate": []int64{1, 2, 3, 4, 5, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		"0.5 head concentrate": []int64{1, 2, 3, 4, 5, 6, 7, 8, -1, -1, -1, -1, -1, -1, -1, -1},

		// heavy head skew that requires backtrackinng
		// [1 2 3 4 5 6 7 0 0 0 0 0 0 0 ]
		//
		// BOTH head and tail skew, requires backtracking
		// [1 2 3 0 0 0 0 4]
	}

	countDistinct := func(arr []int64) int {
		count := 0
		for i := range arr {
			if arr[i] > EMPTY {
				count++
			}
		}
		return count
	}

	for _, arr := range testCase {
		diluteInsert(arr, 0, len(arr)-1, 1, countDistinct(arr))
	}
}

func TestWriteIndex(t *testing.T) {
	max := 10

	for capacity := 1; capacity < max; capacity++ {
		for count := 1; count <= capacity; count++ {
			for i := 0; i < capacity; i++ {
				if isWriteIndex(i, capacity, count) {
					fmt.Printf("%d ", i)
				} else {
					fmt.Printf("_ ")
				}
			}
			fmt.Println()
		}
	}
}
