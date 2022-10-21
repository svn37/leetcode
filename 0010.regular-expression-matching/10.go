/**
 * Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:
 *
 * 	'.' Matches any single character.​​​​
 * 	'*' Matches zero or more of the preceding element.
 *
 * The matching should cover the entire input string (not partial).
 *
 * Example 1:
 *
 * Input: s = "aa", p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 * Example 2:
 *
 * Input: s = "aa", p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
 *
 * Example 3:
 *
 * Input: s = "ab", p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 * Example 4:
 *
 * Input: s = "aab", p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
 *
 * Example 5:
 *
 * Input: s = "mississippi", p = "mis*is*p*."
 * Output: false
 *
 *
 * Constraints:
 *
 * 	0 <= s.length <= 20
 * 	0 <= p.length <= 30
 * 	s contains only lowercase English letters.
 * 	p contains only lowercase English letters, '.', and '*'.
 * 	It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
 *
 */

package leetcode

// Simple recursive solution
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) ||
			(len(s) > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p))
	} else {
		return len(s) > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}
}

// 2D Dynamic programming solution
// We define dp[i][j] to be true if s[0..i) matches p[0..j) and false otherwise. The state equations will be:
//
// dp[i][j] = dp[i - 1][j - 1], if p[j - 1] != '*' && (s[i - 1] == p[j - 1] || p[j - 1] == '.');
// dp[i][j] = dp[i][j - 2], if p[j - 1] == '*' and the pattern repeats for 0 time;
// dp[i][j] = dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'), if p[j - 1] == '*' and the pattern repeats for at least 1 time.
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// empty string matches empty pattern
	dp[0][0] = true

	// i starts with 0, but j starts with 1
	// first row, because empty string might be captured by a pattern
	// and it is used in later checks
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// if repeats 0 times || repeats at least once
				dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				// if previous chars matched && current char matches
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}

	return dp[m][n]
}

// 1D Dynamic programming solution
func isMatch22(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([]bool, n+1)

	for i := 0; i <= m; i++ {
		// dp[0] from the previous iteration
		prev := dp[0]
		// need to set true for i == 0 here
		dp[0] = i == 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if p[j-1] == '*' {
				// if repeats 0 times || repeats at least once
				dp[j] = dp[j-2] || (i > 0 && dp[j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				// if previous chars matched && current char matches
				dp[j] = i > 0 && prev && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
			prev = temp
		}
	}

	return dp[n]
}
