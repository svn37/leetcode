/**
 * Given an m x n board of characters and a list of strings words, return all words on the board.
 * Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/07/search1.jpg" style="width: 322px; height: 322px;" />
 * Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
 * Output: ["eat","oath"]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/07/search2.jpg" style="width: 162px; height: 162px;" />
 * Input: board = [["a","b"],["c","d"]], words = ["abcb"]
 * Output: []
 *
 *
 * Constraints:
 *
 * 	m == board.length
 * 	n == board[i].length
 * 	1 <= m, n <= 12
 * 	board[i][j] is a lowercase English letter.
 * 	1 <= words.length <= 3 * 10^4
 * 	1 <= words[i].length <= 10
 * 	words[i] consists of lowercase English letters.
 * 	All the strings of words are unique.
 *
 */

package leetcode

type TrieNode struct {
	children []*TrieNode
	word     string
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
	}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, c := range word {
		nextNode := node.children[int(c-'a')]
		if nextNode == nil {
			nextNode = NewTrie()
			node.children[int(c-'a')] = nextNode
		}
		node = nextNode
	}
	node.word = word
}

var dirs []int = []int{0, 1, 0, -1, 0}

func dfs(board [][]byte, i, j int, m, n int, trie *TrieNode, res *[]string) {
	char := board[i][j]
	if char == '#' || trie.children[char-'a'] == nil {
		return
	}

	trie = trie.children[char-'a']
	if len(trie.word) > 0 {
		*res = append(*res, trie.word)
		trie.word = "" // get rid of duplicates which can be found later
	}

	// instead of visited[m][n]
	board[i][j] = '#'
	for k := 0; k < 4; k++ {
		x := i + dirs[k]
		y := j + dirs[k+1]
		if x >= 0 && x < m && y >= 0 && y < n {
			dfs(board, x, y, m, n, trie, res)
		}
	}
	board[i][j] = char
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}

	trie := NewTrie()
	for i := range words {
		trie.Insert(words[i])
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, m, n, trie, &res)
		}
	}
	return res
}
