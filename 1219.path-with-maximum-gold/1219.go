/**
 * In a gold mine grid of size m * n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.
 * Return the maximum amount of gold you can collect under the conditions:
 *
 * 	Every time you are located in a cell you will collect all the gold in that cell.
 * 	From your position you can walk one step to the left, right, up or down.
 * 	You can't visit the same cell more than once.
 * 	Never visit a cell with 0 gold.
 * 	You can start and stop collecting gold from any position in the grid that has some gold.
 *
 *
 * Example 1:
 *
 * Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
 * Output: 24
 * Explanation:
 * [[0,6,0],
 *  [5,8,7],
 *  [0,9,0]]
 * Path to get the maximum gold, 9 -> 8 -> 7.
 *
 * Example 2:
 *
 * Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
 * Output: 28
 * Explanation:
 * [[1,0,7],
 *  [2,0,6],
 *  [3,4,5],
 *  [0,3,0],
 *  [9,0,20]]
 * Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
 *
 *
 * Constraints:
 *
 * 	1 <= grid.length, grid[i].length <= 15
 * 	0 <= grid[i][j] <= 100
 * 	There are at most 25 cells containing gold.
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution if the number of cells is small (25)
func dfs(grid [][]int, i, j, m, n int) int {
	dirs := []int{0, 1, 0, -1, 0}

	result := 0
	prev := grid[i][j]
	grid[i][j] = 0

	for d := 0; d < 4; d++ {
		dx := i + dirs[d]
		dy := j + dirs[d+1]
		if dx >= 0 && dy >= 0 && dx < m && dy < n && grid[dx][dy] > 0 {
			result = max(result, grid[dx][dy]+dfs(grid, dx, dy, m, n))
		}
	}
	grid[i][j] = prev

	return result
}

func getMaximumGold(grid [][]int) int {
	maxVal := 0

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				maxVal = max(maxVal, grid[i][j]+dfs(grid, i, j, m, n))
			}
		}
	}

	return maxVal
}
