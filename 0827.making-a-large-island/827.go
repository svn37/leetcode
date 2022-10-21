/**
 * In a 2D grid of 0s and 1s, we change at most one 0 to a 1.
 *
 * After, what is the size of the largest island? (An island is a 4-directionally connected group of 1s).
 *
 * Example 1:
 *
 *
 * Input: [[1, 0], [0, 1]]
 * Output: 3
 * Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.
 *
 *
 * Example 2:
 *
 *
 * Input: [[1, 1], [1, 0]]
 * Output: 4
 * Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4.
 *
 * Example 3:
 *
 *
 * Input: [[1, 1], [1, 1]]
 * Output: 4
 * Explanation: Can't change any 0 to 1, only one island with area = 4.
 *
 *
 *
 * Notes:
 *
 *
 * 	1 <= grid.length = grid[0].length <= 50.
 * 	0 <= grid[i][j] <= 1.
 *
 *
 *
 *
 */

package leetcode

var dirs []int = []int{0, 1, 0, -1, 0}

func colorIsland(grid [][]int, i, j, m, color int) int {
	grid[i][j] = color

	area := 1
	for d := 0; d < 4; d++ {
		dx := i + dirs[d]
		dy := j + dirs[d+1]

		if dx >= 0 && dy >= 0 && dx < m && dy < m && grid[dx][dy] == 1 {
			area += colorIsland(grid, dx, dy, m, color)
		}
	}
	return area
}

func largestIsland(grid [][]int) int {
	sizes := make([]int, 2) // first two are dummy values

	m := len(grid)

	// color all the islands and save their area to sizes
	color := 2
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				sizes = append(sizes, colorIsland(grid, i, j, m, color))
				color++
			}
		}
	}

	// if the area is filled with '1's, return the area
	if len(sizes) == 3 && sizes[2] == m*m {
		return sizes[2]
	}

	maxArea := 0

	// optimization to create set only once
	// if the value is not cur, it was set in another generation
	set := make(map[int]int)
	cur := 1

	// check all '0's and find which is the largest area that can be combined from surrounding cells
	// as long as they are different islands
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				area := 1 // originally includes the cell
				for d := 0; d < 4; d++ {
					dx := i + dirs[d]
					dy := j + dirs[d+1]

					if dx >= 0 && dy >= 0 && dx < m && dy < m {
						color := grid[dx][dy]
						if color > 0 && set[color] != cur {
							area += sizes[color]
							set[color] = cur
						}
					}
				}

				if maxArea < area {
					maxArea = area
				}
				cur++
			}
		}
	}

	return maxArea
}
