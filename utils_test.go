package apma

import (
	"fmt"
	"testing"
)

func TestWriteIndex(t *testing.T) {
	E := EMPTY
	cases := map[string]struct {
		count int
		arr   []int64
	}{
		"a": {7, []int64{1, 2, 4, 5, E, E, E, E, E, E, 6, 7, 8}},
		"b": {7, []int64{1, 2, 4, 5, E, 6, 7, 8}},
		"c": {6, []int64{E, 1, 2, 4, 5, E, 7, 8}},
		"d": {5, []int64{1, 2, 4, 5, 6, E, E}},
		"e": {3, []int64{1, 2, 4, E, E, E, E}},
		"f": {3, []int64{E, E, E, E, 3, 4, 5}},
		"g": {7, []int64{E, 1, 2, 4, 5, 6, 7, 8}},
	}
	for key, v := range cases {
		fmt.Println(key)
		_ = diluteInsert(v.arr, 0, len(v.arr)-1, 9, v.count)
		// fmt.Printf("%+v\n%+v\n", v.arr, output)
	}
}
