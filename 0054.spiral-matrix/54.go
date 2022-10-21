/**
 * Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 *  [ 1, 2, 3 ],
 *  [ 4, 5, 6 ],
 *  [ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [
 *   [1, 2, 3, 4],
 *   [5, 6, 7, 8],
 *   [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */

package leetcode

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	if len(matrix) == 0 {
		return result
	}

	rowBegin := 0
	rowEnd := len(matrix) - 1

	colBegin := 0
	colEnd := len(matrix[0]) - 1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for j := colBegin; j <= colEnd; j++ {
			result = append(result, matrix[rowBegin][j])
		}
		rowBegin++

		for j := rowBegin; j <= rowEnd; j++ {
			result = append(result, matrix[j][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for j := colEnd; j >= colBegin; j-- {
				result = append(result, matrix[rowEnd][j])
			}
			rowEnd--
		}

		if colBegin <= colEnd {
			for j := rowEnd; j >= rowBegin; j-- {
				result = append(result, matrix[j][colBegin])
			}
			colBegin++
		}
	}

	return result
}
