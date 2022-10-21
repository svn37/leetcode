/**
 * Given a 2D grid consists of 0s (land) and 1s (water).  An island is a maximal 4-directionally connected group of <font face="monospace">0</font>s and a closed island is an island totally (all left, top, right, bottom) surrounded by 1s.
 * Return the number of closed islands.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/10/31/sample_3_1610.png" style="width: 240px; height: 120px;" />
 *
 * Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
 * Output: 2
 * Explanation:
 * Islands in gray are closed because they are completely surrounded by water (group of 1s).
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/10/31/sample_4_1610.png" style="width: 160px; height: 80px;" />
 *
 * Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
 * Output: 1
 *
 * Example 3:
 *
 * Input: grid = [[1,1,1,1,1,1,1],
 *                [1,0,0,0,0,0,1],
 *                [1,0,1,1,1,0,1],
 *                [1,0,1,0,1,0,1],
 *                [1,0,1,1,1,0,1],
 *                [1,0,0,0,0,0,1],
 *                [1,1,1,1,1,1,1]]
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	1 <= grid.length, grid[0].length <= 100
 * 	0 <= grid[i][j] <=1
 *
 */

package leetcode

var dirs = []int{0, 1, 0, -1, 0}

func dfs(grid [][]int, i, j int, m, n int) {
	grid[i][j] = 1
	for k := 0; k < 4; k++ {
		dx := i + dirs[k]
		dy := j + dirs[k+1]
		if dx >= 0 && dx < m && dy >= 0 && dy < n && grid[dx][dy] == 0 {
			dfs(grid, dx, dy, m, n)
		}
	}
}

func closedIsland(grid [][]int) int {
	m := len(grid)
	if m < 2 {
		return 0
	}
	n := len(grid[0])
	if n < 2 {
		return 0
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			dfs(grid, 0, j, m, n)
		}

		if grid[m-1][j] == 0 {
			dfs(grid, m-1, j, m, n)
		}
	}

	for i := 1; i < m-1; i++ {
		if grid[i][0] == 0 {
			dfs(grid, i, 0, m, n)
		}

		if grid[i][n-1] == 0 {
			dfs(grid, i, n-1, m, n)
		}
	}

	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				dfs(grid, i, j, m-1, n-1)
				count++
			}
		}
	}
	return count
}
