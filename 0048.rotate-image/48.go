/**
 * You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
 * You have to rotate the image <a href="https://en.wikipedia.org/wiki/In-place_algorithm" target="_blank">in-place</a>, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg" style="width: 642px; height: 242px;" />
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[7,4,1],[8,5,2],[9,6,3]]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg" style="width: 800px; height: 321px;" />
 * Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 * Example 3:
 *
 * Input: matrix = [[1]]
 * Output: [[1]]
 *
 * Example 4:
 *
 * Input: matrix = [[1,2],[3,4]]
 * Output: [[3,1],[4,2]]
 *
 *
 * Constraints:
 *
 * 	matrix.length == n
 * 	matrix[i].length == n
 * 	1 <= n <= 20
 * 	-1000 <= matrix[i][j] <= 1000
 *
 */

package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	for j := 0; j < n-1; j++ {
		for i := j; i < n-1-j; i++ {
			matrix[j][i], matrix[i][n-j-1], matrix[n-j-1][n-i-1], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[j][i], matrix[i][n-j-1], matrix[n-j-1][n-i-1]
		}
	}
}
