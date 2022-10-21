/**
 * Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
 *
 * Example 1:
 * <img src="https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png" style="width: 412px; height: 161px;" />
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
 *
 * Example 2:
 *
 * Input: height = [4,2,0,3,2,5]
 * Output: 9
 *
 *
 * Constraints:
 *
 * 	n == height.length
 * 	0 <= n <= 3 * 10^4
 * 	0 <= height[i] <= 10^5
 *
 */

package leetcode

type Heap struct {
	storage []*Height
}

func NewHeap(a []*Height) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMax() *Height {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *Height) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].height > h.storage[parent].height {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) Len() int {
	return len(h.storage)
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

	if l < len(h.storage) && h.storage[l].height > h.storage[largest].height {
		largest = l
	}

	if r < len(h.storage) && h.storage[r].height > h.storage[largest].height {
		largest = r
	}

	if largest != i {
		h.storage[largest], h.storage[i] = h.storage[i], h.storage[largest]
		h.heapify(largest)
	}
}

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() int {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// Method 1. Heap
type Height struct {
	height, idx int
}

func trap(height []int) int {
	heap := &Heap{}
	for i, h := range height {
		heap.Push(&Height{
			height: h,
			idx:    i,
		})
	}

	if heap.Len() < 2 {
		return 0
	}
	start := heap.ExtractMax()
	end := heap.ExtractMax()

	area := end.height * (abs(start.idx-end.idx) - 1)

	if start.idx > end.idx {
		start, end = end, start
	}

	for !heap.Empty() {
		max := heap.ExtractMax()

		if max.idx > start.idx && max.idx < end.idx {
			area -= max.height
		} else if max.idx < start.idx {
			area += max.height * (start.idx - max.idx - 1)
			start = max
		} else {
			area += max.height * (max.idx - end.idx - 1)
			end = max
		}
	}

	return area
}

// Method 2. Stack
func trap2(height []int) int {
	stack := &Stack{}
	area := 0

	for i := range height {
		for !stack.Empty() && height[stack.Peek()] < height[i] {
			prev := stack.Pop()
			if stack.Empty() {
				break
			}
			h := min(height[i], height[stack.Peek()]) - height[prev]
			w := i - stack.Peek() - 1
			area += h * w
		}

		stack.Push(i)
	}

	return area
}

// Method 3. Two pointers.
func trap3(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	area := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				area += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				area += rightMax - height[right]
			}
			right--
		}
	}

	return area
}

// Method 4. Dynamic programming.
func trap4(height []int) int {
	if len(height) == 0 {
		return 0
	}

	// max to the left of height[i]
	left := make([]int, len(height))
	// max to the right of height[i]
	right := make([]int, len(height))

	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[len(right)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	area := 0

	for i := range height {
		area += min(left[i], right[i]) - height[i]
	}

	if area < 0 {
		return 0
	}

	return area
}
