/**
 * According to the <a href="https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life" target="_blank">Wikipedia's article</a>: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."
 *
 * Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its <a href="https://en.wikipedia.org/wiki/Moore_neighborhood" target="_blank">eight neighbors</a> (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):
 *
 * <ol>
 * 	Any live cell with fewer than two live neighbors dies, as if caused by under-population.
 * 	Any live cell with two or three live neighbors lives on to the next generation.
 * 	Any live cell with more than three live neighbors dies, as if by over-population..
 * 	Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
 * </ol>
 *
 * Write a function to compute the next state (after one update) of the board given its current state. <span>The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously.</span>
 *
 * Example:
 *
 *
 * Input:
 * <span id="example-input-1-1">[
 *   [0,1,0],
 *   [0,0,1],
 *   [1,1,1],
 *   [0,0,0]
 * ]</span>
 * Output:
 * <span id="example-output-1">[
 *   [0,0,0],
 *   [1,0,1],
 *   [0,1,1],
 *   [0,1,0]
 * ]</span>
 *
 *
 * Follow up:
 *
 * <ol>
 * 	Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
 * 	In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?
 * </ol>
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func liveNeighbors(board [][]int, m, n, i, j int) int {
	lives := 0
	for k := max(i-1, 0); k < min(i+2, m); k++ {
		for l := max(j-1, 0); l < min(j+2, n); l++ {
			isLive := board[k][l] & 1
			lives += isLive
		}
	}
	lives -= board[i][j] // subtract oneself
	return lives
}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives := liveNeighbors(board, m, n, i, j)

			if board[i][j] == 1 && (lives == 2 || lives == 3) {
				board[i][j] = 3 // will be live in the next generation
			}

			if board[i][j] == 0 && lives == 3 {
				board[i][j] = 2 // will be live but was dead
			}
		}
	}

	// start new generation
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1 // clear least significant bit
		}
	}
}
