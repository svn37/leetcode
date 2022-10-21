/**
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
 *
 * Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)
 *
 * Example 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,1,1,0,1,0,0,0,0,0,0,0,0],
 *  [0,1,0,0,1,1,0,0,1,0,1,0,0],
 *  [0,1,0,0,1,1,0,0,1,1,1,0,0],
 *  [0,0,0,0,0,0,0,0,0,0,1,0,0],
 *  [0,0,0,0,0,0,0,1,1,1,0,0,0],
 *  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 * Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
 *
 * Example 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 *
 * Note: The length of each dimension in the given grid does not exceed 50.
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addIsland(grid [][]int, i, j, m, n int) int {
	count := 1
	grid[i][j] = 2

	if i > 0 && grid[i-1][j] == 1 {
		count += addIsland(grid, i-1, j, m, n)
	}

	if i+1 < m && grid[i+1][j] == 1 {
		count += addIsland(grid, i+1, j, m, n)
	}

	if j > 0 && grid[i][j-1] == 1 {
		count += addIsland(grid, i, j-1, m, n)
	}

	if j+1 < n && grid[i][j+1] == 1 {
		count += addIsland(grid, i, j+1, m, n)
	}

	return count
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, addIsland(grid, i, j, m, n))
			}
		}
	}

	return maxArea
}
