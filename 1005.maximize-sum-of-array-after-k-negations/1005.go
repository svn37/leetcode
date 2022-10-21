/**
 * Given an array A of integers, we must modify the array in the following way: we choose an i and replace A[i] with -A[i], and we repeat this process K times in total.  (We may choose the same index i multiple times.)
 *
 * Return the largest possible sum of the array after modifying it in this way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = <span id="example-input-1-1">[4,2,3]</span>, K = <span id="example-input-1-2">1</span>
 * Output: <span id="example-output-1">5
 * Explanation: Choose indices (1,) and A becomes [4,-2,3].</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: A = <span id="example-input-2-1">[3,-1,0,2]</span>, K = <span id="example-input-2-2">3</span>
 * Output: 6
 * <span id="example-output-1">Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: A = <span id="example-input-3-1">[2,-3,-1,5,-4]</span>, K = <span id="example-input-3-2">2</span>
 * Output: <span id="example-output-3">13
 * </span><span id="example-output-1">Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].</span>
 *
 * </div>
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= A.length <= 10000
 * 	1 <= K <= 10000
 * 	-100 <= A[i] <= 100
 * </ol>
 *
 */

package leetcode

import (
	"math"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Method 1. Sorting.
// the logic is like this:
// before K is 0, use the following rules, after K is 0 - just sum the remaining elements
// all negative numbers change sign and added to the sum, BUT we keep track of the minimum value
// the minimum one is used when we first encounter the positive number (or 0)
// positive number:
// if K is divisible by 2, nothing changes, just sum
// if not:
// if positive number is smaller than value, it doesn't count, subtract it
// however if it is greater, add it to sum BUT the minimum value which added previously is subtracted (twice)
func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	sum := 0
	minVal := math.MaxInt64
	for i := 0; i < len(A); i, K = i+1, K-1 {
		if K > 0 {
			if A[i] < 0 {
				sum += -A[i]
				minVal = min(minVal, -A[i])
			} else {
				if K%2 == 0 {
					sum += A[i]
				} else {
					if minVal < A[i] {
						sum += A[i]
						sum -= 2 * minVal
					} else {
						sum -= A[i]
					}
				}
				K = 0
			}
		} else {
			sum += A[i]
		}
	}
	return sum
}

// Method 2. Use priority queue.
type Heap struct {
	storage []int
}

func NewHeap(a []int) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMin() int {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val int) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i] < h.storage[parent] {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) Empty() bool {
	return len(h.storage) == 0
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 != 0 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func (h *Heap) leftchild(i int) int {
	return 2*i + 1
}

func (h *Heap) rightchild(i int) int {
	return 2*i + 2
}

func (h *Heap) heapify(i int) {
	l := h.leftchild(i)
	r := h.rightchild(i)
	smallest := i

	if l < len(h.storage) && h.storage[l] < h.storage[smallest] {
		smallest = l
	}

	if r < len(h.storage) && h.storage[r] < h.storage[smallest] {
		smallest = r
	}

	if smallest != i {
		h.storage[smallest], h.storage[i] = h.storage[i], h.storage[smallest]
		h.heapify(smallest)
	}
}

func largestSumAfterKNegations2(A []int, K int) int {
	heap := NewHeap(A)

	for K > 0 {
		heap.Push(-heap.ExtractMin())
		K--
	}

	sum := 0
	for !heap.Empty() {
		sum += heap.ExtractMin()
	}
	return sum
}
