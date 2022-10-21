/**
 * Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
 * <ol>
 * 	Each row must contain the digits 1-9 without repetition.
 * 	Each column must contain the digits 1-9 without repetition.
 * 	Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
 * </ol>
 * Note:
 *
 * 	A Sudoku board (partially filled) could be valid but is not necessarily solvable.
 * 	Only the filled cells need to be validated according to the mentioned rules.
 *
 *
 * Example 1:
 * <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" style="height:250px; width:250px" />
 * Input: board =
 * [["5","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: true
 *
 * Example 2:
 *
 * Input: board =
 * [["8","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
 *
 *
 * Constraints:
 *
 * 	board.length == 9
 * 	board[i].length == 9
 * 	board[i][j] is a digit or '.'.
 *
 */

package leetcode

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// we compare i values (epochs)
	// instead we could create sets inside the first for loop
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// check rows
			if board[i][j] != '.' {
				digit := board[i][j] - '1'

				if rows[digit] == i+1 {
					return false
				}
				rows[digit] = i + 1
			}

			// check columns
			if board[j][i] != '.' {
				digit := board[j][i] - '1'

				if cols[digit] == i+1 {
					return false
				}
				cols[digit] = i + 1
			}

			// check squares
			rowIndex := 3 * (i / 3)
			colIndex := 3 * (i % 3)
			cell := board[rowIndex+j/3][colIndex+j%3]

			if cell != '.' {
				digit := cell - '1'

				if squares[digit] == i+1 {
					return false
				}
				squares[digit] = i + 1
			}
		}
	}

	return true
}

// Performance is worse, because there are lots of string concatenations.
// come up with some cell encoding and store it in one set.
func isValidSudoku2(board [][]byte) bool {
	set := make(map[string]struct{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell != '.' {
				number := cell - '0'

				row := fmt.Sprintf("%d in row %d", number, i)
				col := fmt.Sprintf("%d in col %d", number, j)
				block := fmt.Sprintf("%d in block %d", number, 3*(i/3)+j/3)

				_, seenRow := set[row]
				_, seenCol := set[col]
				_, seenBlock := set[block]
				if seenRow || seenCol || seenBlock {
					return false
				}

				set[row] = struct{}{}
				set[col] = struct{}{}
				set[block] = struct{}{}
			}
		}
	}

	return true
}
