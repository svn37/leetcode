/**
 * Given a m * n matrix of distinct numbers, return all lucky numbers in the matrix in any order.
 * A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.
 *
 * Example 1:
 *
 * Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
 * Output: [15]
 * Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column
 *
 * Example 2:
 *
 * Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
 * Output: [12]
 * Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.
 *
 * Example 3:
 *
 * Input: matrix = [[7,8],[1,2]]
 * Output: [7]
 *
 *
 * Constraints:
 *
 * 	m == mat.length
 * 	n == mat[i].length
 * 	1 <= n, m <= 50
 * 	1 <= matrix[i][j] <= 10^5.
 * 	All elements in the matrix are distinct.
 *
 */

package leetcode

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	rows := make([]int, m)
	cols := make([]int, n)
	for i := range cols {
		cols[i] = math.MinInt64
	}

	for i := 0; i < m; i++ {
		minVal := math.MaxInt64
		for j := 0; j < n; j++ {
			minVal = min(minVal, matrix[i][j])
			cols[j] = max(matrix[i][j], cols[j])
		}
		rows[i] = minVal
	}

	luckyNums := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] == cols[j] {
				luckyNums = append(luckyNums, matrix[i][j])
			}
		}
	}
	return luckyNums
}
