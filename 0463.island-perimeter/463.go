/**
 * You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
 * Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
 * The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
 *
 * Example 1:
 * <img src="https://assets.leetcode.com/uploads/2018/10/12/island.png" style="width: 221px; height: 213px;" />
 * Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
 * Output: 16
 * Explanation: The perimeter is the 16 yellow stripes in the image above.
 *
 * Example 2:
 *
 * Input: grid = [[1]]
 * Output: 4
 *
 * Example 3:
 *
 * Input: grid = [[1,0]]
 * Output: 4
 *
 *
 * Constraints:
 *
 * 	row == grid.length
 * 	col == grid[i].length
 * 	1 <= row, col <= 100
 * 	grid[i][j] is 0 or 1.
 *
 */

package leetcode

func islandPerimeter(grid [][]int) int {
	islands, neighbors := 0, 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				islands++

				// count right neighbors
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					neighbors++
				}

				// count down neighbors
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					neighbors++
				}
			}
		}
	}

	return 4*islands - 2*neighbors
}
