/**
 * Given a m * n matrix grid which is sorted in non-increasing order both row-wise and column-wise.
 *
 * Return the number of negative numbers in grid.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
 * Output: 8
 * Explanation: There are 8 negatives number in the matrix.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[3,2],[1,0]]
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[1,-1],[-1,-1]]
 * Output: 3
 *
 *
 * Example 4:
 *
 *
 * Input: grid = [[-1]]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 100
 * 	-100 <= grid[i][j] <= 100
 *
 */

package leetcode

// This solution uses the fact that the negative regions of the matrix will form a "staircase" shape
func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	r := m - 1
	c := 0
	count := 0

	for r >= 0 && c < n {
		if grid[r][c] < 0 {
			r--
			count += n - c
		} else {
			c++
		}
	}
	return count
}
