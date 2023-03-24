/**
 * Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
 *
 * Example 1:
 *
 * Input: matrix =
 * [
 *   [0,1,1,1],
 *   [1,1,1,1],
 *   [0,1,1,1]
 * ]
 * Output: 15
 * Explanation:
 * There are 10 squares of side 1.
 * There are 4 squares of side 2.
 * There is  1 square of side 3.
 * Total number of squares = 10 + 4 + 1 = 15.
 *
 * Example 2:
 *
 * Input: matrix =
 * [
 *   [1,0,1],
 *   [1,1,0],
 *   [1,1,0]
 * ]
 * Output: 7
 * Explanation:
 * There are 6 squares of side 1.
 * There is 1 square of side 2.
 * Total number of squares = 6 + 1 = 7.
 *
 *
 * Constraints:
 *
 * 	1 <= arr.length <= 300
 * 	1 <= arr[0].length <= 300
 * 	0 <= arr[i][j] <= 1
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// no need for preprocessing
func countSquares(matrix [][]int) int {
	res := 0
	for i := range matrix {
		for j := range matrix[i] {
			// if i == 0 || j == 0, simply add 1 square
			if matrix[i][j] > 0 && i > 0 && j > 0 {
				// clever trick - to count 3*3 square, change matrix[i][j] to 2
				matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
			}
			res += matrix[i][j]
		}
	}
	return res
}
