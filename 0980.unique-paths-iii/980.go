/**
 * On a 2-dimensional grid, there are 4 types of squares:
 *
 *
 * 	1 represents the starting square.  There is exactly one starting square.
 * 	2 represents the ending square.  There is exactly one ending square.
 * 	0 represents empty squares we can walk over.
 * 	-1 represents obstacles that we cannot walk over.
 *
 *
 * Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]</span>
 * Output: <span id="example-output-1">2</span>
 * Explanation: We have the following two paths:
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
 * 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[[1,0,0,0],[0,0,0,0],[0,0,0,2]]</span>
 * Output: <span id="example-output-2">4</span>
 * Explanation: We have the following four paths:
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
 * 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
 * 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
 * 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">[[0,1],[2,0]]</span>
 * Output: <span id="example-output-3">0</span>
 * Explanation:
 * There is no path that walks over every empty square exactly once.
 * Note that the starting and ending square can be anywhere in the grid.
 *
 * </div>
 * </div>
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= grid.length * grid[0].length <= 20
 * </ol>
 */

package leetcode

var dirs []int = []int{0, 1, 0, -1, 0}

// theoretically what we could do is
// store state the same way we did in "shortest path to get all keys"
// and use it as composite key in the memo hashmap
func dfs(grid [][]int, x, y, cur, target int) int {
	// check if ending square
	if grid[x][y] == 2 {
		if cur == target {
			return 1
		}
		return 0
	}

	m, n := len(grid), len(grid[0])

	grid[x][y] = -1
	cur++

	result := 0

	for i := 0; i < 4; i++ {
		dx := x + dirs[i]
		dy := y + dirs[i+1]

		// continue if square contains 0 or 2
		if dx >= 0 && dx < m && dy >= 0 && dy < n && grid[dx][dy] >= 0 {
			result += dfs(grid, dx, dy, cur, target)
		}
	}

	// restore for other dfs searches
	grid[x][y] = 0

	return result
}

func uniquePathsIII(grid [][]int) int {
	// starting square coordinates
	var sx, sy int
	// count of empty squares we must walk over
	// including starting square
	empty := 1

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				// empty square
				empty++
			} else if grid[i][j] == 1 {
				// starting square coordinates
				sx, sy = i, j
			}
		}
	}

	// number of walks possible
	return dfs(grid, sx, sy, 0, empty)
}
