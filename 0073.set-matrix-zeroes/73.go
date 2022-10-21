/**
 * Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank">in-place</a>.
 * Follow up:
 *
 * 	A straight forward solution using O(mn) space is probably a bad idea.
 * 	A simple improvement uses O(m + n) space, but still not the best solution.
 * 	Could you devise a constant space solution?
 *
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg" style="width: 450px; height: 169px;" />
 * Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: [[1,0,1],[0,0,0],[1,0,1]]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg" style="width: 450px; height: 137px;" />
 * Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 * Constraints:
 *
 * 	m == matrix.length
 * 	n == matrix[0].length
 * 	1 <= m, n <= 200
 * 	-2^31 <= matrix[i][j] <= 2^31 - 1
 *
 */

package leetcode

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	// mark row and column start as zero
	var firstRow bool
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 && matrix[i][j] == 0 {
				firstRow = true
				continue
			}

			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// set all the values properly
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	// matrix[0][0] is the corner case, because requires extra firstRow variable
	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if firstRow {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
}
