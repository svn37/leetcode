/**
 * Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
 * You have the following three operations permitted on a word:
 *
 * 	Insert a character
 * 	Delete a character
 * 	Replace a character
 *
 *
 * Example 1:
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 * Example 2:
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 * Constraints:
 *
 * 	0 <= word1.length, word2.length <= 500
 * 	word1 and word2 consist of lowercase English letters.
 *
 */

package leetcode

/*
 * If word1[i - 1] != word2[j - 1], we need to consider three cases.
 *
 *  Replace word1[i - 1] by word2[j - 1] (dp[i][j] = dp[i - 1][j - 1] + 1);
 *  If word1[0..i - 1) = word2[0..j) then delete word1[i - 1] (dp[i][j] = dp[i - 1][j] + 1);
 *  If word1[0..i) + word2[j - 1] = word2[0..j) then insert word2[j - 1] to word1[0..i) (dp[i][j] = dp[i][j - 1] + 1).
 */

// - - 1 2 3 4 5
// - - h o r s e
// 1 r 1 2 2 3 4
// 2 o 1 1 2 3 4
// 3 s 3 2 2 2 3

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := range memo {
		memo[i] = make([]int, len(word2)+1)
		memo[i][0] = i
		if i == 0 {
			for j := 1; j <= len(word2); j++ {
				memo[i][j] = j
			}
		}
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = min(memo[i-1][j-1], min(memo[i-1][j], memo[i][j-1])) + 1
			}
		}
	}

	return memo[len(word1)][len(word2)]
}
