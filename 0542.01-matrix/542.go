/**
 * Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.
 *
 * The distance between two adjacent cells is 1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * [[0,0,0],
 *  [0,1,0],
 *  [0,0,0]]
 *
 * Output:
 * [[0,0,0],
 *  [0,1,0],
 *  [0,0,0]]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [[0,0,0],
 *  [0,1,0],
 *  [1,1,1]]
 *
 * Output:
 * [[0,0,0],
 *  [0,1,0],
 *  [1,2,1]]
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The number of elements of the given matrix will not exceed 10,000.
 * 	There are at least one 0 in the given matrix.
 * 	The cells are adjacent in only four directions: up, down, left and right.
 * </ol>
 *
 */

package leetcode

import "math"

type Queue struct {
	storage [][2]int
}

func (q *Queue) Push(cell [2]int) {
	q.storage = append(q.storage, cell)
}

func (q *Queue) Poll() [2]int {
	cell := q.storage[0]
	q.storage = q.storage[1:]
	return cell
}

func (q *Queue) Empty() bool {
	return len(q.storage) == 0
}

func updateMatrix(matrix [][]int) [][]int {
	queue := &Queue{}
	dirs := []int{0, 1, 0, -1, 0}

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				queue.Push([2]int{i, j})
			} else {
				matrix[i][j] = math.MaxInt64
			}
		}
	}

	for !queue.Empty() {
		cell := queue.Poll()
		for d := 0; d < 4; d++ {
			x := cell[0] + dirs[d]
			y := cell[1] + dirs[d+1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[cell[0]][cell[1]] {
				queue.Push([2]int{x, y})
				matrix[x][y] = matrix[cell[0]][cell[1]] + 1
			}
		}
	}
	return matrix
}
