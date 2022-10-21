/**
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 * A sudoku solution must satisfy all of the following rules:
 * <ol>
 * 	Each of the digits 1-9 must occur exactly once in each row.
 * 	Each of the digits 1-9 must occur exactly once in each column.
 * 	Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
 * </ol>
 * The '.' character indicates empty cells.
 *
 * Example 1:
 * <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" style="height:250px; width:250px" />
 * Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
 * Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
 * Explanation: The input board is shown above and the only valid solution is shown below:
 * <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png" style="height:250px; width:250px" />
 *
 *
 * Constraints:
 *
 * 	board.length == 9
 * 	board[i].length == 9
 * 	board[i][j] is a digit or '.'.
 * 	It is guaranteed that the input board has only one solution.
 *
 */

package leetcode

// Simple solver, without many optimizations

func isValid(board [][]byte, row, col int, char byte) bool {
	blockRow := (row / 3) * 3
	blockCol := (col / 3) * 3
	for i := 0; i < 9; i++ {
		// check if the column had the same character before
		// check if the row had the same character before
		// check if the square had the same character before
		if board[i][col] == char ||
			board[row][i] == char ||
			board[blockRow+i/3][blockCol+i%3] == char {
			return false
		}
	}
	return true
}

// row and col are starting positions for searching through the board
func solve(board [][]byte, row, col int) bool {
	for i := row; i < 9; i++ {
		for j := col; j < 9; j++ {
			if board[i][j] == '.' {
				for char := byte('1'); char <= byte('9'); char++ {
					if isValid(board, i, j, char) {
						// assign the number and see if this choice
						// will satisfy all the constraints in the future
						board[i][j] = char

						if solve(board, i, j+1) {
							// we need to find at least one solution
							return true
						}
						board[i][j] = '.'
					}
				}
				// none of the characters could fill the board properly
				return false
			}
		}
		// reset col here, because it didn't start from 0,
		// but need to in the next iteration
		col = 0
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}
