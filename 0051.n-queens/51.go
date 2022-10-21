/**
 * The n-queens puzzle is the problem of placing n queens on an n&times;n chessboard such that no two queens attack each other.
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/10/12/8-queens.png" style="width: 258px; height: 276px;" />
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
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
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
 *
 *
 */

package leetcode

// Simple solver
func isValid(nQueens [][]byte, row, col, n int) bool {
	// check if the column had a queen before
	for i := 0; i < row; i++ {
		if nQueens[i][col] == 'Q' {
			return false
		}
	}

	// check if the 45° diagonal had a queen before
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if nQueens[i][j] == 'Q' {
			return false
		}
	}

	// check if the 135° diagonal had a queen before
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if nQueens[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func solve(nQueens [][]byte, row, n int, res *[][]string) {
	if row == n {
		// convert to string
		answer := make([]string, len(nQueens))
		for i := range nQueens {
			answer[i] = string(nQueens[i])
		}
		// add to final array
		*res = append(*res, answer)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(nQueens, row, col, n) {
			nQueens[row][col] = 'Q'
			solve(nQueens, row+1, n, res)
			nQueens[row][col] = '.'
		}
	}
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	// board representation with queens
	nQueens := make([][]byte, n)
	for i := range nQueens {
		nQueens[i] = make([]byte, n)
		for j := range nQueens[i] {
			nQueens[i][j] = '.'
		}
	}

	solve(nQueens, 0, n, &res)
	return res
}

// Memoize whether the same column, row, or diagonal
// already have the queen in flag array
func solve2(nQueens [][]byte, row, n int, flag []bool, res *[][]string) {
	if row == n {
		// convert to string
		answer := make([]string, len(nQueens))
		for i := range nQueens {
			answer[i] = string(nQueens[i])
		}
		// add to final array
		*res = append(*res, answer)
		return
	}

	for col := 0; col < n; col++ {
		if !flag[col] && !flag[n+row+col] && !flag[4*n-2+col-row] {
			flag[col] = true
			flag[n+row+col] = true
			flag[4*n-2+col-row] = true

			nQueens[row][col] = 'Q'
			solve2(nQueens, row+1, n, flag, res)
			nQueens[row][col] = '.'

			flag[col] = false
			flag[n+row+col] = false
			flag[4*n-2+col-row] = false
		}
	}
}

func solveNQueens2(n int) [][]string {
	res := [][]string{}
	// board representation with queens
	nQueens := make([][]byte, n)
	for i := range nQueens {
		nQueens[i] = make([]byte, n)
		for j := range nQueens[i] {
			nQueens[i][j] = '.'
		}
	}

	// the number of columns is n,
	// the number of 45° diagonals is 2 * n - 1,
	// the number of 135° diagonals is also 2 * n - 1

	// flag[0] to flag[n - 1] indicate if the column had a queen before
	// flag[n] to flag[3 * n - 2] indicate if the 45° diagonal had a queen before
	// flag[3 * n - 1] to flag[5 * n - 3] indicate if the 135° diagonal had a queen before
	flag := make([]bool, 5*n-2)

	solve2(nQueens, 0, n, flag, &res)
	return res
}
