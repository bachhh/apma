package apma

import (
	"math"
)

const EMPTY int64 = -1

const LO_TH = 0.5 // lower density threshold
const HI_TH = 2.0 // higher density threshold

/* PMA is a sorted structure for storing integers with support for O(logN) Insert / Search() operation
   and continuous memory arrangement for efficient Scan() operation.
*/
type PMA struct {
	// only accept unsigned int
	// use -1 to mark empty slot
	arr []int64

	sizeDistinct int // number of distinct elements
	// a PMA needs to define 4 density thresholds, 1 upper and lower, and
	level int // keep track of how many levels we have

	segmentSize int
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

// Resize allocate a new array of appropriate size when top level density reaches outside of our density bound.
// At the base, SegmentSize is used to derive all other all parameters of our array,
//  - the numSegment is the Hyperceil (i.e. closest power of 2 rounded up ) of ideal numSeg

func (this *PMA) resize() {
}

/*
 * To insert, find the correct non-empty index of x and it's coresponding "segment" node in the segment binary tree
 * If the current segment has enough free space (threshold is below upper bound), Insert x and rebalance the whole segment
 * If not, recursively "travel up" the tree to find a "parent segment" that has enough free space. Then rebalance all entries inside this segment
 * If the even root does not have enough free space, grow the entire array
 */
func (this *PMA) Insert(x uint) {
	i := this.find(x)

	curLevel := 0 // leaf
	segLo, segHi := this.getSegmentLevel(i, curLevel)
	countDistinct := 0

	left, right := i, i+1
	for curLevel <= this.level {
		for ; left >= segLo; left-- {
			if this.arr[left] > EMPTY {
				countDistinct++
			}
		}

		for ; right <= segHi; right++ {
			if this.arr[right] > EMPTY {
				countDistinct++
			}
		}
		density := float64(countDistinct) / float64(segHi-segLo)
		p_l, t_l := this.getDensity(curLevel)
		if p_l <= density && density <= t_l {
			break
		}

		// current segment density not suitable, goes up
		curLevel++
		segLo, segHi = this.getSegmentLevel(i, curLevel)
	}

	// we reach here either cause curLevel > top level, or curLevel's density is suitable
	// left = seglow -1, right = seghigh +1
	left, right = left+1, right-1
	if curLevel > this.level { // reached top level without suitable segment
		this.upsize()
	}

	// TODO: insert
	this.spreadInsert(left, right, x)
}

func (this *PMA) upsize() {
}

// spread all elements in the range array[left, right] evenly spaced
// insert the x element at the appropriate place also
func (this *PMA) spreadInsert(left, right int, x uint, count int) {
	// s is the "spacing" between each non-zero element

	s := (right - left + 1) / count

	r, w := left, left
	for w < right {
		if this.arr[r] == EMPTY {
			r++
		} else if this.arr[w] == EMPTY {
			this.arr[w] = this.arr[r]
			w += s
		}
	}
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

// getSegmentLevel return the [low, high] of the segment that index belongs to
// Explain: at level 0, new segment is located every <segmentSize> index
// At level 1, 2 segments are merged, so a new segment is located every <2*segmentSize>
// At level 2, 4 segments are merged, so a new segment is located every <4*segmentSize>
// At level l, a new segment is located at every <2^l*segmentSize>
//
// If @param: level are are higher than max level of array, just return
// the whole array i.e.: [0, len(array)-1]
func (this *PMA) getSegmentLevel(index int, level int) (low, high int) {
	segmentSize := int(math.Pow(2, float64(level))) * this.segmentSize
	if segmentSize > len(this.arr) {
		segmentSize = len(this.arr)
	}
	low = (index / segmentSize) * segmentSize // floor off to closest segment size
	return low, low + segmentSize - 1
}
