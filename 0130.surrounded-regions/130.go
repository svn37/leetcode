/**
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn&rsquo;t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
 *
 */

package leetcode

func mark(board [][]byte, i, j int) {
	board[i][j] = '*'
	dirs := []int{0, 1, 0, -1, 0}

	for d := 0; d < 4; d++ {
		dx := i + dirs[d]
		dy := j + dirs[d+1]
		if dx >= 0 && dy >= 0 && dx < len(board) && dy < len(board[dx]) && board[dx][dy] == 'O' {
			mark(board, dx, dy)
		}
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			mark(board, 0, j)
		}

		if board[m-1][j] == 'O' {
			mark(board, m-1, j)
		}
	}

	for j := 1; j < m-1; j++ {
		if board[j][0] == 'O' {
			mark(board, j, 0)
		}

		if board[j][n-1] == 'O' {
			mark(board, j, n-1)
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}
