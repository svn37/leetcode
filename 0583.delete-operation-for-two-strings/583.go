/**
 *
 * Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.
 *
 *
 * Example 1:<br />
 *
 * Input: "sea", "eat"
 * Output: 2
 * Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
 *
 *
 *
 * Note:<br>
 * <ol>
 * The length of given words won't exceed 500.
 * Characters in given words can only be lower-case letters.
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

// 2D Dynamic Programming solution
// To make them identical, just find the longest common subsequence.
// The rest of the characters have to be deleted from the both the strings,
// which does not belong to longest common subsequence.
func minDistance(word1 string, word2 string) int {
	W1 := []byte(word1)
	W2 := []byte(word2)
	LW1 := len(W1)
	LW2 := len(W2)

	memo := make([][]int, LW1+1)
	for i := range memo {
		memo[i] = make([]int, LW2+1)
	}

	for i := 1; i <= LW1; i++ {
		for j := 1; j <= LW2; j++ {
			if W1[i-1] == W2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}
	return LW1 + LW2 - 2*memo[LW1][LW2]
}
