/**
 * A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Now suppose you are given the locations and height of all the buildings as shown on a cityscape photo (Figure A), write a program to output the skyline formed by these buildings collectively (Figure B).
 * <a href="/static/images/problemset/skyline1.jpg" target="_blank"><img alt="Buildings" src="https://assets.leetcode.com/uploads/2018/10/22/skyline1.png" style="max-width: 45%; border-width: 0px; border-style: solid;" /> </a> <a href="/static/images/problemset/skyline2.jpg" target="_blank"> <img alt="Skyline Contour" src="https://assets.leetcode.com/uploads/2018/10/22/skyline2.png" style="max-width: 45%; border-width: 0px; border-style: solid;" /> </a>
 *
 * The geometric information of each building is represented by a triplet of integers [Li, Ri, Hi], where Li and Ri are the x coordinates of the left and right edge of the ith building, respectively, and Hi is its height. It is guaranteed that 0 &le; Li, Ri &le; INT_MAX, 0 < Hi &le; INT_MAX, and Ri - Li > 0. You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
 *
 * For instance, the dimensions of all buildings in Figure A are recorded as: [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] .
 *
 * The output is a list of "key points" (red dots in Figure B) in the format of [ [x1,y1], [x2, y2], [x3, y3], ... ] that uniquely defines a skyline. A key point is the left endpoint of a horizontal line segment. Note that the last key point, where the rightmost building ends, is merely used to mark the termination of the skyline, and always has zero height. Also, the ground in between any two adjacent buildings should be considered part of the skyline contour.
 *
 * For instance, the skyline in Figure B should be represented as:[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ].
 *
 * Notes:
 *
 *
 * 	The number of buildings in any input list is guaranteed to be in the range [0, 10000].
 * 	The input list is already sorted in ascending order by the left x position Li.
 * 	The output list must be sorted by the x position.
 * 	There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...[2 3], [4 5], [7 5], [11 5], [12 7]...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...[2 3], [4 5], [12 7], ...]
 *
 *
 */

package leetcode

import (
	"math/rand"
)

// Method 1. Divide and conquer
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// check for redundant skylines, a skyline is
// redundant if it has same height or left as previous
func add(res [][]int, skyline []int) [][]int {
	n := len(res)
	if n > 0 && res[n-1][1] == skyline[1] {
		return res
	}

	if n > 0 && res[n-1][0] == skyline[0] {
		res[n-1][1] = max(res[n-1][1], skyline[1])
		return res
	}

	res = append(res, skyline)
	return res
}

// similar to merge in merge sort
func merge(left, right [][]int) [][]int {
	res := make([][]int, 0, len(left)+len(right))

	// maintain heights of left and right skylines
	var h1, h2 int
	var x, maxh int

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		// compare x coordinates and put the smaller one in result
		if left[i][0] < right[j][0] {
			h1 = left[i][1]

			x = left[i][0]
			i++
		} else if left[i][0] > right[j][0] {
			h2 = right[j][1]

			x = right[j][0]
			j++
		} else {
			h1 = left[i][1]
			h2 = right[j][1]

			x = right[j][0]
			i++
			j++
		}
		maxh = max(h1, h2)
		res = add(res, []int{x, maxh})
	}

	for i < len(left) {
		res = add(res, left[i])
		i++
	}

	for j < len(right) {
		res = add(res, right[j])
		j++
	}

	return res
}

// this function is similar to mergeSort
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	if len(buildings) == 1 {
		skylines := make([][]int, 0, 2)
		building := buildings[0]

		skylines = append(skylines, []int{building[0], building[2]})
		skylines = append(skylines, []int{building[1], 0})
		return skylines
	}

	mid := len(buildings) / 2

	left := getSkyline(buildings[:mid])
	right := getSkyline(buildings[mid:])
	return merge(left, right)
}

// Method 2. Heap
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

func (h *Heap) Max() int {
	return h.storage[0]
}

func (h *Heap) ExtractMax() int {
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
	for parent != -1 && h.storage[i] > h.storage[parent] {
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
	largest := i

	if l < len(h.storage) && h.storage[l] > h.storage[largest] {
		largest = l
	}

	if r < len(h.storage) && h.storage[r] > h.storage[largest] {
		largest = r
	}

	if largest != i {
		h.storage[largest], h.storage[i] = h.storage[i], h.storage[largest]
		h.heapify(largest)
	}
}

func qsort(a [][]int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a))

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i][0] < a[right][0] || (a[i][0] == a[right][0] && a[i][1] < a[right][1]) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
}

func getSkyline2(buildings [][]int) [][]int {
	result := [][]int{}
	// coordinates of points (both start and end) sorted by x-axis
	points := make([][]int, 0, len(buildings))

	for _, building := range buildings {
		points = append(points, []int{building[0], -building[2]})
		points = append(points, []int{building[1], building[2]})
	}

	qsort(points)

	// if value is more than zero, the item has been deleted
	m := make(map[int]int)

	heap := &Heap{}
	heap.Push(0)

	prev := 0
	for _, point := range points {
		if point[1] < 0 {
			heap.Push(-point[1])
		} else {
			// instead of O(n) heap deletion
			m[point[1]]++
		}

		cur := heap.Max()
		for m[cur] > 0 {
			heap.ExtractMax()
			m[cur]--
			cur = heap.Max()
		}
		if cur != prev {
			result = append(result, []int{point[0], cur})
			prev = cur
		}
	}

	return result
}
