/**
 * Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
 * Note: You can only move either down or right at any point in time.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg" style="width: 242px; height: 242px;" />
 * Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
 * Output: 7
 * Explanation: Because the path 1 &rarr; 3 &rarr; 1 &rarr; 1 &rarr; 1 minimizes the sum.
 *
 * Example 2:
 *
 * Input: grid = [[1,2,3],[4,5,6]]
 * Output: 12
 *
 *
 * Constraints:
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 200
 * 	0 <= grid[i][j] <= 100
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	memo := make([][]int, len(grid))
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				memo[i][j] = grid[i][j] + min(memo[i-1][j], memo[i][j-1])
			} else {
				if i > 0 {
					memo[i][j] = grid[i][j] + memo[i-1][j]
				} else if j > 0 {
					memo[i][j] = grid[i][j] + memo[i][j-1]
				} else {
					memo[i][j] = grid[i][j]
				}
			}
		}
	}

	return memo[len(grid)-1][len(grid[0])-1]
}
