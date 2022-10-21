/**
 * Given a square array of integers A, we want the minimum sum of a falling path through A.
 * A falling path starts at any element in the first row, and chooses one element from each row.  The next row's choice must be in a column that is different from the previous row's column by at most one.
 *
 * Example 1:
 *
 * Input: <span id="example-input-1-1">[[1,2,3],[4,5,6],[7,8,9]]</span>
 * Output: <span id="example-output-1">12</span>
 * Explanation:
 * The possible falling paths are:
 *
 * 	[1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
 * 	[2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
 * 	[3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
 *
 * The falling path with the smallest sum is [1,4,7], so the answer is 12.
 *
 * Constraints:
 *
 * 	1 <= A.length == A[0].length <= 100
 * 	-100 <= A[i][j] <= 100
 *
 */

package leetcode

import "math"

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

func minFallingPathSum(A [][]int) int {
	m := len(A)

	for i := 1; i < m; i++ {
		for j := 0; j < m; j++ {
			topLeft := max(0, j-1)
			topRight := min(m-1, j+1)

			A[i][j] += min(A[i-1][j], min(A[i-1][topLeft], A[i-1][topRight]))
		}
	}

	res := math.MaxInt64
	for i := range A[m-1] {
		res = min(res, A[m-1][i])
	}

	return res
}
