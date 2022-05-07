package apma

import (
	"fmt"
)

const EMPTY int64 = -1

const LO_TH = 0.5 // lower density threshold
const HI_TH = 2.0 // higher density threshold

type PMA struct {
	// only accept unsigned int
	// use -1 to mark empty slot
	arr          []int64
	sizeDistinct int // number of distinct elements
	// a PMA needs to define 4 density thresholds, 1 upper and lower, and
	level int // keep track of how many levels we have
	// for both level H (root) and level 0 (leaf) respectively
	// all levels inbetween will be extrapolated using an arithematic sequence.
	// upper density
	t0 float64
	th float64
	// lower density
	p0 float64
	ph float64

	// all other properties can be devised from these fields
}

/*
 * To insert, find the correct non-empty index of x and it's coresponding "segment" node in the segment binary tree
 * If the current segment has enough free space (threshold is below upper bound), Insert x and rebalance the whole segment
 * If not, recursively "travel up" the tree to find a "parent segment" that has enough free space. Then rebalance all entries inside this segment
 * If the even root does not have enough free space, grow the entire array
 */
func (this *PMA) Insert(x uint) {
	// find the correct "leaf" segment where x belongs
	i := this.find(x)

	segLo, segHi := this.getSegment(i)
	countDistinct := 0
	for i := segLo; i <= segHi; i++ {
		if this.arr[i] > EMPTY {
			countDistinct++
		}
	}

	fmt.Println(segLo, segHi)
}

/* From leaf to root, find the lowest node that has correct density
 * (between [lower,upper]) then return the boundary.
 * If the whole array violate density, return [-1, len(array)]
 * Calculation:
 *  h       segment start at floor(index / h) * h
 *  h-1     segment start at floor(index / (h-1)) * (h-1)
 *  etc ...
 */
func (this *PMA) findValidAncestor(segLo, segHi int) {
}

func (this *PMA) getDensity(level int) (lower, upper float64) {
	return 0.5, 1.0
}

// insertRebalance take all element in [lo, hi], together with x, reinsert
// into the array with equal distance empty space.
func (this *PMA) insertRebalance(x uint, segLo, segHi int) {
}

/* modified binary search with gap

* case 1
    l         m         h
   [0 0 0 0 0 0 0 0 0 0 a]
  all entry <= mid is empty, low = mid +1

* case 2 , find(3)
    l     i   m         h
   [0 0 0 2 0 0 0 0 4 0 9]
  mid is empty, but mid - 2 is non empty,
  mid = mid - 2, then do a normal binsearch comparison

* case 3 , find(3)
    l         m         h
   [0 0 0 0 0 2 0 0 4 0 9]
  mid is non-empty -> do a normal binsearch comparison

*/
func (this *PMA) find(x uint) (index int) {
	x64 := int64(x)

	// return index of the element >= x in the array using binary search
	low, high := 0, len(this.arr)-1
	for low < high {
		mid := low + (high-low)/2
		check := mid

		// instead of checking directly, we seek left until we see a non-null entry
		for check >= low && this.arr[check] == EMPTY {
			check--
		}

		// seek past the end of low without seeing a non-empty
		if check < low {
			low = mid + 1
			continue
		}

		if this.arr[check] < x64 {
			low = check + 1
		} else if this.arr[check] >= x64 {
			high = check
		}
	}
	return low
}

// getSegment return the [low, high] of leaf segment that index belongs to
func (this *PMA) getSegment(index int) (l, h int) {
	return 0, 0 // TODO
}
