/**
 * Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
 *
 *
 *
 * <img src="https://assets.leetcode.com/uploads/2018/10/12/histogram.png" style="width: 188px; height: 204px;" /><br />
 * <small>Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].</small>
 *
 *
 *
 * <img src="https://assets.leetcode.com/uploads/2018/10/12/histogram_area.png" style="width: 188px; height: 204px;" /><br />
 * <small>The largest rectangle is shown in the shaded area, which has area = 10 unit.</small>
 *
 *
 *
 * Example:
 *
 *
 * Input: [2,1,5,6,2,3]
 * Output: 10
 *
 *
 */

package leetcode

// Method 1. Stack
type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() int {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Top() int {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// example input [2,1,4,5,6,3]
// stack: [0], index: 1
// stack: [], index: 1, area: 2
// stack: [1], index: 2
// stack: [1 2], index: 3
// stack: [1 2 3], index: 4
// stack: [1 2 3 4], index: 5
// stack: [1 2 3], index: 5, area: 6
// stack: [1 2], index: 5, area: 10
// stack: [1], index: 5, area: 12
// stack: [1 5], index: 6
// stack: [1], index: 6, area: 12
// stack: [], index: 6, area: 6

// when a bar is popped, we calculate the area with the popped bar as smallest bar
func calculateArea(stack *Stack, heights []int, right int) int {
	top := stack.Pop()

	var left int
	if !stack.Empty() {
		left = stack.Top() + 1
	}
	return heights[top] * (right - left)
}

// traverse all bars from left to right, maintain an increasing stack of bars
// O(n) -- every bar is pushed to stack once
// a bar is popped from stack when a bar of smaller height is seen
func largestRectangleArea(heights []int) int {
	stack := &Stack{}
	maxArea := 0

	i := 0
	for i < len(heights) {
		// if bar is higher than top stack, push it to stack
		if stack.Empty() || heights[stack.Top()] <= heights[i] {
			stack.Push(i)
			i++
		} else {
			// otherwise, calculate rectangle area
			// i is right width
			// pay attention, i is not increased here
			// this step is repeated in the next iteration, if heights[i] < heights[top]
			maxArea = max(maxArea, calculateArea(stack, heights, i))
		}
	}

	for !stack.Empty() {
		maxArea = max(maxArea, calculateArea(stack, heights, i))
	}

	return maxArea
}

// Method 2. Dynamic programming.
func largestRectangleArea2(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	// idx of the first bar to the left that is lower than current
	lessFromLeft := make([]int, len(heights))
	lessFromLeft[0] = -1

	// idx of the first bar to the right that is lower than current
	lessFromRight := make([]int, len(heights))
	lessFromRight[len(heights)-1] = len(heights)

	// fill lessFromLeft
	for i := 1; i < len(heights); i++ {
		p := i - 1

		// the trick is we don't need to recalculate linearly
		// we jump through several indices
		for p >= 0 && heights[p] >= heights[i] {
			p = lessFromLeft[p]
		}
		lessFromLeft[i] = p
	}

	// fill lessFromRight
	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1

		for p < len(heights) && heights[p] >= heights[i] {
			p = lessFromRight[p]
		}
		lessFromRight[i] = p
	}

	maxArea := 0
	for i := range heights {
		maxArea = max(maxArea, heights[i]*(lessFromRight[i]-lessFromLeft[i]-1))
	}

	return maxArea
}
