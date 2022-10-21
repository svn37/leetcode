/**
 * Given two strings s and t, return the number of distinct subsequences of s which equals t.
 * A string's subsequence is a new string formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).
 * It's guaranteed the answer fits on a 32-bit signed integer.
 *
 * Example 1:
 *
 * Input: s = "rabbbit", t = "rabbit"
 * Output: 3
 * Explanation:
 * As shown below, there are 3 ways you can generate "rabbit" from S.
 * <u>rabb</u>b<u>it</u>
 * <u>ra</u>b<u>bbit</u>
 * <u>rab</u>b<u>bit</u>
 *
 * Example 2:
 *
 * Input: s = "babgbag", t = "bag"
 * Output: 5
 * Explanation:
 * As shown below, there are 5 ways you can generate "bag" from S.
 * <u>ba</u>b<u>g</u>bag
 * <u>ba</u>bgba<u>g</u>
 * <u>b</u>abgb<u>ag</u>
 * ba<u>b</u>gb<u>ag</u>
 * babg<u>bag</u>
 *
 * Constraints:
 *
 * 	0 <= s.length, t.length <= 1000
 * 	s and t consist of English letters.
 *
 */

package leetcode

// 2D dynamic programming.
func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}

	// the empty string is a subsequence of any string but only 1 time.
	for j := 0; j <= len(s); j++ {
		dp[0][j] = 1
	}

	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}

// 1D dynamic programming.
func numDistinct2(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1

	for j := 1; j <= len(s); j++ {
		for i := len(t); i > 0; i-- {
			if t[i-1] == s[j-1] {
				dp[i] += dp[i-1]
			}
		}
	}
	return dp[len(t)]
}
