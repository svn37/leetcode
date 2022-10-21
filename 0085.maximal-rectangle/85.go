/**
 * Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/14/maximal.jpg" style="width: 402px; height: 322px;" />
 * Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * Output: 6
 * Explanation: The maximal rectangle is shown in the above picture.
 *
 * Example 2:
 *
 * Input: matrix = []
 * Output: 0
 *
 * Example 3:
 *
 * Input: matrix = [["0"]]
 * Output: 0
 *
 * Example 4:
 *
 * Input: matrix = [["1"]]
 * Output: 1
 *
 * Example 5:
 *
 * Input: matrix = [["0","0"]]
 * Output: 0
 *
 *
 * Constraints:
 *
 * 	rows == matrix.length
 * 	cols == matrix.length
 * 	0 <= row, cols <= 200
 * 	matrix[i][j] is '0' or '1'.
 *
 */

package leetcode

// Method 1. Use problem 84 as a subroutine.
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

func calculateArea(stack *Stack, heights []int, right int) int {
	top := stack.Pop()

	var left int
	if !stack.Empty() {
		left = stack.Top() + 1
	}
	return heights[top] * (right - left)
}

func maxHist(heights []int) int {
	stack := &Stack{}
	maxArea := 0

	i := 0
	for i < len(heights) {
		if stack.Empty() || heights[stack.Top()] <= heights[i] {
			stack.Push(i)
			i++
		} else {
			maxArea = max(maxArea, calculateArea(stack, heights, i))
		}
	}

	for !stack.Empty() {
		maxArea = max(maxArea, calculateArea(stack, heights, i))
	}

	return maxArea
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	for i := range heights {
		if matrix[0][i] == '1' {
			heights[i] = 1
		}
	}
	result := maxHist(heights)

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		result = max(result, maxHist(heights))
	}

	return result
}
