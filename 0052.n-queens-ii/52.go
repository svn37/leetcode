/**
 * The n-queens puzzle is the problem of placing n queens on an n&times;n chessboard such that no two queens attack each other.
 *
 * <img src="https://assets.leetcode.com/uploads/2018/10/12/8-queens.png" style="width: 258px; height: 276px;" />
 *
 * Given an integer n, return the number of distinct solutions to the n-queens puzzle.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
 * [
 *  [".Q..",  // Solution 1
 *   "...Q",
 *   "Q...",
 *   "..Q."],
 *
 *  ["..Q.",  // Solution 2
 *   "Q...",
 *   "...Q",
 *   ".Q.."]
 * ]
 *
 *
 */

package leetcode

func countSolutions(row, n int, flag []bool) int {
	if row == n {
		return 1
	}

	count := 0
	for col := 0; col < n; col++ {
		if !flag[col] && !flag[n+row+col] && !flag[4*n-2+col-row] {
			flag[col] = true
			flag[n+row+col] = true
			flag[4*n-2+col-row] = true

			// don't need to set queens on the board at all,
			// because there is no solution to reconstruct
			count += countSolutions(row+1, n, flag)

			flag[col] = false
			flag[n+row+col] = false
			flag[4*n-2+col-row] = false
		}
	}
	return count
}

func totalNQueens(n int) int {
	// the number of columns is n,
	// the number of 45° diagonals is 2 * n - 1,
	// the number of 135° diagonals is also 2 * n - 1

	// flag[0] to flag[n - 1] indicate if the column had a queen before
	// flag[n] to flag[3 * n - 2] indicate if the 45° diagonal had a queen before
	// flag[3 * n - 1] to flag[5 * n - 3] indicate if the 135° diagonal had a queen before
	flag := make([]bool, 5*n-2)

	return countSolutions(0, n, flag)
}
