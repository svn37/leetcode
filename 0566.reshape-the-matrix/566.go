/**
 * In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.
 *
 *
 *
 * You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.
 *
 *  The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.
 *
 *
 *
 * If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
 *
 *
 * Example 1:<br />
 *
 * Input:
 * nums =
 * [[1,2],
 *  [3,4]]
 * r = 1, c = 4
 * Output:
 * [[1,2,3,4]]
 * Explanation:<br>The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
 *
 *
 *
 * Example 2:<br />
 *
 * Input:
 * nums =
 * [[1,2],
 *  [3,4]]
 * r = 2, c = 4
 * Output:
 * [[1,2],
 *  [3,4]]
 * Explanation:<br>There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
 *
 *
 *
 * Note:<br>
 * <ol>
 * The height and width of the given matrix is in range [1, 100].
 * The given r and c are all positive.
 * </ol>
 *
 */

package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	m := len(nums)
	n := len(nums[0])
	if m*n != r*c {
		return nums
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, 0, c)
	}

	var rowNum, colNum int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if colNum == c {
				rowNum++
				colNum = 0
			}
			result[rowNum] = append(result[rowNum], nums[i][j])
			colNum++
		}
	}
	return result
}
