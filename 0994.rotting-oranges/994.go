/**
 * In a given grid, each cell can have one of three values:
 *
 *
 * 	the value 0 representing an empty cell;
 * 	the value 1 representing a fresh orange;
 * 	the value 2 representing a rotten orange.
 *
 *
 * Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
 *
 * Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.
 *
 *
 *
 * <div>
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/02/16/oranges.png" style="width: 712px; height: 150px;" />
 *
 *
 * Input: <span id="example-input-1-1">[[2,1,1],[1,1,0],[0,1,1]]</span>
 * Output: <span id="example-output-1">4</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[[2,1,1],[0,1,1],[1,0,1]]</span>
 * Output: <span id="example-output-2">-1</span>
 * Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">[[0,2]]</span>
 * Output: <span id="example-output-3">0</span>
 * Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= grid.length <= 10
 * 	1 <= grid[0].length <= 10
 * 	grid[i][j] is only 0, 1, or 2.
 * </ol>
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

type Queue struct {
	arr []*Cell
}

func (q *Queue) Push(val *Cell) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Poll() *Cell {
	val := q.arr[0]
	q.arr = q.arr[1:]
	return val
}

func (q *Queue) Len() int {
	return len(q.arr)
}

func (q *Queue) Empty() bool {
	return len(q.arr) == 0
}

type Cell struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	queue := &Queue{}
	countFresh := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue.Push(&Cell{i, j})
			} else if grid[i][j] == 1 {
				countFresh++
			}
		}
	}
	if countFresh == 0 {
		return 0
	}

	level := 0
	dirs := []int{0, 1, 0, -1, 0}

	for !queue.Empty() {
		level++
		for size := queue.Len(); size > 0; size-- {
			cell := queue.Poll()
			for i := 0; i < 4; i++ {
				x := cell.x + dirs[i]
				y := cell.y + dirs[i+1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					queue.Push(&Cell{x, y})
					grid[x][y] = 2
					countFresh--
				}
			}
		}
	}
	if countFresh > 0 {
		return -1
	}
	return level - 1
}
