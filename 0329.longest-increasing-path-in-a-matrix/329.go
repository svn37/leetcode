/**
 * Given an integer matrix, find the length of the longest increasing path.
 *
 * From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).
 *
 * Example 1:
 *
 *
 * Input: nums =
 * [
 *   [<font color="red">9</font>,9,4],
 *   [<font color="red">6</font>,6,8],
 *   [<font color="red">2</font>,<font color="red">1</font>,1]
 * ]
 * Output: 4
 * Explanation: The longest increasing path is [1, 2, 6, 9].
 *
 *
 * Example 2:
 *
 *
 * Input: nums =
 * [
 *   [<font color="red">3</font>,<font color="red">4</font>,<font color="red">5</font>],
 *   [3,2,<font color="red">6</font>],
 *   [2,2,1]
 * ]
 * Output: 4
 * Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
 *
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dirs []int = []int{0, 1, 0, -1, 0}

func dfs(matrix [][]int, i, j int, m, n int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	maxVal := 1
	for k := 0; k < 4; k++ {
		x := i + dirs[k]
		y := j + dirs[k+1]
		// clever trick - in dfs we continue searching for increasing path only if
		// matrix[x][y] > matrix[i][j], so we don't need a visited[m][n] array
		if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j] {
			maxVal = max(maxVal, 1+dfs(matrix, x, y, m, n, cache))
		}
	}
	cache[i][j] = maxVal
	return maxVal
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	// at least one number, matrix is not empty
	maxVal := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxVal = max(maxVal, dfs(matrix, i, j, m, n, cache))
		}
	}
	return maxVal
}
