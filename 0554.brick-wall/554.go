/**
 * There is a brick wall in front of you. The wall is rectangular and has several rows of bricks. The bricks have the same height but different width. You want to draw a vertical line from the top to the bottom and cross the least bricks.
 *
 * The brick wall is represented by a list of rows. Each row is a list of integers representing the width of each brick in this row from left to right.
 *
 * If your line go through the edge of a brick, then the brick is not considered as crossed. You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.
 *
 * You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.
 *
 *
 *
 * Example:
 *
 *
 * Input: [[1,2,2,1],
 *         [3,1,2],
 *         [1,3,2],
 *         [2,4],
 *         [3,1,2],
 *         [1,3,1,1]]
 *
 * Output: 2
 *
 * Explanation:
 * <img src="https://assets.leetcode.com/uploads/2018/10/12/brick_wall.png" style="width: 100%; max-width: 350px" />
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The width sum of bricks in different rows are the same and won't exceed INT_MAX.
 * 	The number of bricks in each row is in range [1,10,000]. The height of wall is in range [1,10,000]. Total number of bricks of the wall won't exceed 20,000.
 * </ol>
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	minVal := len(wall)

	for i := range wall {
		// skip the rightmost brick
		for j := 0; j < len(wall[i])-1; j++ {
			if j > 0 {
				wall[i][j] += wall[i][j-1]
			}

			minVal = min(minVal, len(wall)-m[wall[i][j]]-1)
			m[wall[i][j]]++
		}
	}

	return minVal
}
