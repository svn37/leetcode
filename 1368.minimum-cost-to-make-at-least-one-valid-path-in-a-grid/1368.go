/**
 * Given a m x n grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell. The sign of grid[i][j] can be:
 * 	1 which means go to the cell to the right. (i.e go from grid[i][j] to grid[i][j + 1])
 * 	2 which means go to the cell to the left. (i.e go from grid[i][j] to grid[i][j - 1])
 * 	3 which means go to the lower cell. (i.e go from grid[i][j] to grid[i + 1][j])
 * 	4 which means go to the upper cell. (i.e go from grid[i][j] to grid[i - 1][j])
 *
 * Notice that there could be some invalid signs on the cells of the grid which points outside the grid.
 * You will initially start at the upper left cell (0,0). A valid path in the grid is a path which starts from the upper left cell (0,0) and ends at the bottom-right cell (m - 1, n - 1) following the signs on the grid. The valid path doesn't have to be the shortest.
 * You can modify the sign on a cell with cost = 1. You can modify the sign on a cell one time only.
 * Return the minimum cost to make the grid have at least one valid path.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/02/13/grid1.png" style="width: 542px; height: 528px;" />
 * Input: grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
 * Output: 3
 * Explanation: You will start at point (0, 0).
 * The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)
 * The total cost = 3.
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/02/13/grid2.png" style="width: 419px; height: 408px;" />
 * Input: grid = [[1,1,3],[3,2,2],[1,1,4]]
 * Output: 0
 * Explanation: You can follow the path from (0, 0) to (2, 2).
 *
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/02/13/grid3.png" style="width: 314px; height: 302px;" />
 * Input: grid = [[1,2],[4,3]]
 * Output: 1
 *
 * Example 4:
 *
 * Input: grid = [[2,2,2],[2,2,2]]
 * Output: 3
 *
 * Example 5:
 *
 * Input: grid = [[4]]
 * Output: 0
 *
 *
 * Constraints:
 *
 * 	m == grid.length
 * 	n == grid[i].length
 * 	1 <= m, n <= 100
 *
 */

package leetcode

type Queue struct {
	items []Cell
}

func (q *Queue) Push(val Cell) {
	q.items = append(q.items, val)
}

func (q *Queue) Poll() Cell {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

type Cell = [2]int

var dirs [][]int = [][]int{
	{0, 1},  // right
	{0, -1}, // left
	{1, 0},  // down
	{-1, 0}, // up
}

func dfs(grid [][]int, x, y int, queue *Queue) bool {
	m, n := len(grid), len(grid[0])
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 5 {
		return false
	}
	if x == m-1 && y == n-1 {
		return true
	}
	nextDir := grid[x][y] - 1
	grid[x][y] = 5
	queue.Push([2]int{x, y})

	return dfs(grid, x+dirs[nextDir][0], y+dirs[nextDir][1], queue)
}

// dfs until you find the dead end
// add each visited cell to the bfs queue during the dfs
// change each visited cell's value to 5, which means 'visited'
// do the bfs, and change level (level is the number of changes you have to make to reach the m-1, n-1 cell)
func minCost(grid [][]int) int {
	queue := &Queue{}
	level := 0
	if dfs(grid, 0, 0, queue) {
		return level
	}
	for {
		level++
		for i := queue.Len(); i > 0; i-- {
			cell := queue.Poll()
			x, y := cell[0], cell[1]
			for d := 0; d < 4; d++ {
				dx := x + dirs[d][0]
				dy := y + dirs[d][1]
				found := dfs(grid, dx, dy, queue)
				if found {
					return level
				}
			}
		}
	}
}
