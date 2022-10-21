/**
 *
 * In a 2D grid from (0, 0) to (N-1, N-1), every cell contains a 1, except those cells in the given list mines which are 0.  What is the largest axis-aligned plus sign of 1s contained in the grid?  Return the order of the plus sign.  If there is none, return 0.
 *
 * An "axis-aligned plus sign of 1s of order k" has some center grid[x][y] = 1 along with 4 arms of length k-1 going up, down, left, and right, and made of 1s.  This is demonstrated in the diagrams below.  Note that there could be 0s or 1s beyond the arms of the plus sign, only the relevant area of the plus sign is checked for 1s.
 *
 *
 * Examples of Axis-Aligned Plus Signs of Order k:<br />
 * Order 1:
 * 000
 * 010
 * 000
 *
 * Order 2:
 * 00000
 * 00100
 * 01110
 * 00100
 * 00000
 *
 * Order 3:
 * 0000000
 * 0001000
 * 0001000
 * 0111110
 * 0001000
 * 0001000
 * 0000000
 *
 *
 * Example 1:<br />
 * Input: N = 5, mines = [[4, 2]]
 * Output: 2
 * Explanation:
 * 11111
 * 11111
 * 11111
 * 11111
 * 11011
 * In the above grid, the largest plus sign can only be order 2.  One of them is marked in bold.
 *
 *
 * Example 2:<br />
 * Input: N = 2, mines = []
 * Output: 1
 * Explanation:
 * There is no plus sign of order 2, but there is of order 1.
 *
 *
 * Example 3:<br />
 * Input: N = 1, mines = [[0, 0]]
 * Output: 0
 * Explanation:
 * There is no plus sign, so return 0.
 *
 *
 * Note:<br><ol>
 * N will be an integer in the range [1, 500].
 * mines will have length at most 5000.
 * mines[i] will be length 2 and consist of integers in the range [0, N-1].
 * (Additionally, programs submitted in C, C++, or C# will be judged with a slightly smaller time limit.)
 * </ol>
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

func orderOfLargestPlusSign(N int, mines [][]int) int {
	grid := make([][]int, N)
	for i := range grid {
		grid[i] = make([]int, N)
		for j := range grid[i] {
			grid[i][j] = N
		}
	}

	for _, mine := range mines {
		grid[mine[0]][mine[1]] = 0
	}

	for i := range grid {
		for j, k, l, r, u, d := 0, N-1, 0, 0, 0, 0; j < N; j, k = j+1, k-1 {
			// left
			if grid[i][j] == 0 {
				l = 0
			} else {
				l++
			}
			grid[i][j] = min(grid[i][j], l)

			// right
			if grid[i][k] == 0 {
				r = 0
			} else {
				r++
			}
			grid[i][k] = min(grid[i][k], r)

			// up
			if grid[j][i] == 0 {
				u = 0
			} else {
				u++
			}
			grid[j][i] = min(grid[j][i], u)

			// down
			if grid[k][i] == 0 {
				d = 0
			} else {
				d++
			}
			grid[k][i] = min(grid[k][i], d)
		}
	}

	maxVal := 0
	for i := range grid {
		for j := range grid {
			maxVal = max(maxVal, grid[i][j])
		}
	}
	return maxVal
}
