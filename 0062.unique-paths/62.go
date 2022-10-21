/**
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 * The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
 * How many possible unique paths are there?
 *
 * Example 1:
 * <img src="https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png" style="width: 400px; height: 183px;" />
 * Input: m = 3, n = 7
 * Output: 28
 *
 * Example 2:
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
 * 1. Right -> Down -> Down
 * 2. Down -> Down -> Right
 * 3. Down -> Right -> Down
 *
 * Example 3:
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 * Example 4:
 *
 * Input: m = 3, n = 3
 * Output: 6
 *
 *
 * Constraints:
 *
 * 	1 <= m, n <= 100
 * 	It's guaranteed that the answer will be less than or equal to 2 * 10^9.
 *
 */

package leetcode

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	for i := 1; i < len(memo); i++ {
		memo[i][0] = 1
	}

	for i := 1; i < len(memo[0]); i++ {
		memo[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i-1][j] + memo[i][j-1]
		}
	}

	return memo[m-1][n-1]
}
