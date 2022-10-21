/**
 * Given a 2D grid of 0s and 1s, return the number of elements in the largest square subgrid that has all 1s on its border, or 0 if such a subgrid doesn't exist in the grid.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: 9
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[1,1,0,0]]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= grid.length <= 100
 * 	1 <= grid[0].length <= 100
 * 	grid[i][j] is 0 or 1
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// horizontal and vertical auxiliary arrays
	left := make([][]int, m)
	top := make([][]int, m)

	// fill the arrays
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		top[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if j > 0 {
					left[i][j] = left[i][j-1] + 1
				} else {
					left[i][j] = 1
				}

				if i > 0 {
					top[i][j] = top[i-1][j] + 1
				} else {
					top[i][j] = 1
				}
			}
		}
	}

	maxSide := 0
	// start from the bottom rightmost corner
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			side := min(left[i][j], top[i][j])

			for side > maxSide {
				validTopSide := left[i-side+1][j] >= side
				validLeftSide := top[i][j-side+1] >= side

				if validTopSide && validLeftSide {
					maxSide = side
				}
				side--
			}
		}
	}

	return maxSide * maxSide
}
