/**
 * Given an N x N grid containing only values 0 and 1, where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized and return the distance.
 * The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.
 * If no land or water exists in the grid, return -1.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex1.JPG" style="width: 185px; height: 87px;" />
 *
 * Input: <span id="example-input-1-1">[[1,0,1],[0,0,0],[1,0,1]]</span>
 * Output: <span id="example-output-1">2</span>
 * Explanation:
 * The cell (1, 1) is as far as possible from all the land with distance 2.
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/05/03/1336_ex2.JPG" style="width: 184px; height: 87px;" />
 *
 * Input: <span id="example-input-2-1">[[1,0,0],[0,0,0],[0,0,0]]</span>
 * Output: <span id="example-output-2">4</span>
 * Explanation:
 * The cell (2, 2) is as far as possible from all the land with distance 4.
 *
 *
 * <span>Note:</span>
 * <ol>
 * 	<span>1 <= grid.length == grid[0].length <= 100</span>
 * 	<span>grid[i][j] is 0 or 1</span>
 * </ol>
 *
 */

package leetcode

type Queue struct {
	items []*Cell
}

func (q *Queue) Push(val *Cell) {
	q.items = append(q.items, val)
}

func (q *Queue) Poll() *Cell {
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

type Cell struct {
	x, y int
}

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	queue := &Queue{}

	// add all the land cells to the queue
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				queue.Push(&Cell{i, j})
			}
		}
	}

	if queue.Len() == m*n || queue.Len() == 0 {
		return -1
	}

	// bfs
	dirs := []int{0, 1, 0, -1, 0}
	size, level := 0, -1

	for !queue.Empty() {
		if size == 0 {
			size = queue.Len()
			level++
		}

		cell := queue.Poll()
		for i := 0; i < 4; i++ {
			dx := cell.x + dirs[i]
			dy := cell.y + dirs[i+1]

			if dx >= 0 && dx < m && dy >= 0 && dy < n && grid[dx][dy] == 0 {
				// instead of visited set
				grid[dx][dy] = 1
				queue.Push(&Cell{dx, dy})
			}
		}

		size--
	}
	return level
}
