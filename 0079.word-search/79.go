/**
 * Given a 2D board and a word, find if the word exists in the grid.
 * The word can be constructed from letters of sequentially adjacent cells, where "adjacent" cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/04/word2.jpg" style="width: 322px; height: 242px;" />
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
 * Output: true
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg" style="width: 322px; height: 242px;" />
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
 * Output: true
 *
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/15/word3.jpg" style="width: 322px; height: 242px;" />
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
 * Output: false
 *
 *
 * Constraints:
 *
 * 	board and word consists only of lowercase and uppercase English letters.
 * 	1 <= board.length <= 200
 * 	1 <= board[i].length <= 200
 * 	1 <= word.length <= 10^3
 *
 */

package leetcode

func search(board [][]byte, i, j, pos int, word string) bool {
	if board[i][j] == word[pos] {
		if pos+1 == len(word) {
			return true
		}

		board[i][j] = '0' // dummy value

		if i > 0 {
			up := search(board, i-1, j, pos+1, word)
			if up {
				return true
			}
		}

		if i < len(board)-1 {
			down := search(board, i+1, j, pos+1, word)
			if down {
				return true
			}
		}

		if j > 0 {
			left := search(board, i, j-1, pos+1, word)
			if left {
				return true
			}
		}

		if j < len(board[0])-1 {
			right := search(board, i, j+1, pos+1, word)
			if right {
				return true
			}
		}

		board[i][j] = word[pos] // restore values
	}

	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	for i := range board {
		for j := range board[i] {
			ok := search(board, i, j, 0, word)
			if ok {
				return true
			}
		}
	}
	return false
}
