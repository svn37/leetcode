/**
 * An N x N board contains only 0s and 1s. In each move, you can swap any 2 rows with each other, or any 2 columns with each other.
 *
 * What is the minimum number of moves to transform the board into a "chessboard" - a board where no 0s and no 1s are 4-directionally adjacent? If the task is impossible, return -1.
 *
 *
 * Examples:
 * Input: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
 * Output: 2
 * Explanation:
 * One potential sequence of moves is shown below, from left to right:
 *
 * 0110     1010     1010
 * 0110 --> 1010 --> 0101
 * 1001     0101     1010
 * 1001     0101     0101
 *
 * The first move swaps the first and second column.
 * The second move swaps the second and third row.
 *
 *
 * Input: board = [[0, 1], [1, 0]]
 * Output: 0
 * Explanation:
 * Also note that the board with 0 in the top left corner,
 * 01
 * 10
 *
 * is also a valid chessboard.
 *
 * Input: board = [[1, 0], [1, 0]]
 * Output: -1
 * Explanation:
 * No matter what sequence of moves you make, you cannot end with a valid chessboard.
 *
 *
 * Note:
 *
 *
 * 	board will have the same number of rows and columns, a number in the range [2, 30].
 * 	board[i][j] will be only 0s or 1s.
 *
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func movesToChessboard(board [][]int) int {
	var rowSum, colSum, rowSwap, colSwap int

	for i := range board {
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}

		for j := range board {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] != 0 {
				return -1
			}
		}
	}

	n := len(board)
	if rowSum != n/2 && rowSum != (n+1)/2 {
		return -1
	}
	if colSum != n/2 && colSum != (n+1)/2 {
		return -1
	}

	if n%2 == 1 {
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
	} else {
		colSwap = min(n-colSwap, colSwap)
		rowSwap = min(n-rowSwap, rowSwap)
	}
	return (colSwap + rowSwap) / 2
}
