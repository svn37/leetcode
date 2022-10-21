/**
 * Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
 *
 * 	'?' Matches any single character.
 * 	'*' Matches any sequence of characters (including the empty sequence).
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
 * Input: s = "aa", p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 *
 * Example 3:
 *
 * Input: s = "cb", p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
 *
 * Example 4:
 *
 * Input: s = "adceb", p = "*a*b"
 * Output: true
 * Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
 *
 * Example 5:
 *
 * Input: s = "acdcb", p = "a*c?b"
 * Output: false
 *
 *
 * Constraints:
 *
 * 	0 <= s.length, p.length <= 2000
 * 	s contains only lowercase English letters.
 * 	p contains only lowercase English letters, '?' or '*'.
 *
 */

package leetcode

// Simple recursive solution, time limit exceeded
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if p[0] == '*' {
		// optimization - skip several consecutive '*'
		for len(p) > 1 && p[1] == '*' {
			p = p[1:]
		}
		return isMatch(s, p[1:]) ||
			(len(s) > 0 && isMatch(s[1:], p))
	} else {
		return len(s) > 0 && (s[0] == p[0] || p[0] == '?') && isMatch(s[1:], p[1:])
	}
}

// Greedy algorithm, restarts with the latest '*' if pattern didn't match
// The right solution, passes all tests.
func isMatch2(s string, p string) bool {
	i, j, match, starIdx := 0, 0, 0, -1

	for i < len(s) {
		if j < len(p) && (p[j] == '?' || s[i] == p[j]) {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			starIdx = j
			match = i
			j++
		} else if starIdx != -1 {
			// try again to match starting with match++ character
			j = starIdx + 1
			match++
			i = match
		} else {
			return false
		}
	}

	for j < len(p) && p[j] == '*' {
		j++
	}

	return j == len(p)
}

// 2D Dynamic programming solution, similar to regular expression matching.
// dp[i][j] = dp[i - 1][j - 1] && (s[i - 1] == p[j - 1] || p[j - 1] == '?'), if p[j - 1] != '*';
// dp[i][j] = dp[i][j - 1] || dp[i - 1][j], if p[j - 1] == '*'
func isMatch3(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || (i > 0 && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
		}
	}

	return dp[m][n]
}

// 1D Dynamic programming solution.
func isMatch4(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([]bool, n+1)

	for i := 0; i <= m; i++ {
		prev := dp[0]
		dp[0] = i == 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if p[j-1] == '*' {
				dp[j] = dp[j-1] || (i > 0 && dp[j])
			} else {
				dp[j] = i > 0 && prev && (s[i-1] == p[j-1] || p[j-1] == '?')
			}
			prev = temp
		}
	}

	return dp[n]
}
