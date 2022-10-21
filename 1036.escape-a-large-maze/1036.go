/**
 * In a 1 million by 1 million grid, the coordinates of each grid square are (x, y).
 * We start at the source square and want to reach the target square.  Each move, we can walk to a 4-directionally adjacent square in the grid that isn't in the given list of blocked squares.
 * Return true if and only if it is possible to reach the target square through a sequence of moves.
 *
 * Example 1:
 *
 * Input: blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
 * Output: false
 * Explanation: The target square is inaccessible starting from the source square, because we can't walk outside the grid.
 *
 * Example 2:
 *
 * Input: blocked = [], source = [0,0], target = [999999,999999]
 * Output: true
 * Explanation: Because there are no blocked cells, it's possible to reach the target square.
 *
 *
 * Constraints:
 *
 * 	0 <= blocked.length <= 200
 * 	blocked[i].length == 2
 * 	0 <= blocked[i][j] < 10^6
 * 	source.length == target.length == 2
 * 	0 <= source[i][j], target[i][j] < 10^6
 * 	source != target
 *
 */

package leetcode

type Queue struct {
	items []Cell
}

func (q *Queue) Push(val Cell) {
	q.items = append(q.items, val)
}

func (q *Queue) PopFront() Cell {
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

var dirs []int = []int{0, 1, 0, -1, 0}

func bfs(source, target []int, blocked [][]int) int {
	queue := &Queue{}
	visited := make(map[Cell]struct{})
	for i := range blocked {
		cell := [2]int{blocked[i][0], blocked[i][1]}
		visited[cell] = struct{}{}
	}

	cell := [2]int{source[0], source[1]}
	queue.Push(cell)
	visited[cell] = struct{}{}

	size := 1

	for !queue.Empty() {
		cell := queue.PopFront()

		for i := 0; i < 4; i++ {
			dx := cell[0] + dirs[i]
			dy := cell[1] + dirs[i+1]

			if dx >= 0 && dx < 1000000 && dy >= 0 && dy < 1000000 {
				cell := [2]int{dx, dy}

				if _, ok := visited[cell]; !ok {
					if cell[0] == target[0] && cell[1] == target[1] {
						return 2
					}

					if size+1 == 20000 {
						return 1
					}

					queue.Push(cell)
					visited[cell] = struct{}{}
					size++
				}
			}
		}
	}
	// bfs finished before 20000 area was reached, there is no way to reach target
	return 0
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	count := bfs(source, target, blocked)
	if count == 2 {
		return true
	}
	count += bfs(target, source, blocked)
	return count == 2
}
