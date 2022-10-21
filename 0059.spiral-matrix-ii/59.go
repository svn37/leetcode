/**
 * Given a positive integer n, generate a square matrix filled with elements from 1 to n^2 in spiral order.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 *  [ 1, 2, 3 ],
 *  [ 8, 9, 4 ],
 *  [ 7, 6, 5 ]
 * ]
 *
 *
 */

package leetcode

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	rowStart := 0
	colStart := 0
	rowEnd := n - 1
	colEnd := n - 1

	num := 0
	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			num++
			m[rowStart][i] = num
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			num++
			m[i][colEnd] = num
		}
		colEnd--

		if rowStart < rowEnd && colStart < colEnd {
			for i := colEnd; i >= colStart; i-- {
				num++
				m[rowEnd][i] = num
			}
			rowEnd--

			for i := rowEnd; i >= rowStart; i-- {
				num++
				m[i][colStart] = num
			}
			colStart++
		}
	}

	return m
}
