/**
 * Given a string s, return the longest palindromic substring in s.
 *
 * Example 1:
 *
 * Input: s = "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 * Example 2:
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 * Example 3:
 *
 * Input: s = "a"
 * Output: "a"
 *
 * Example 4:
 *
 * Input: s = "ac"
 * Output: "a"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 1000
 * 	s consist of only digits and English letters (lower-case and/or upper-case),
 *
 */

package leetcode

// Method 1. Check all 2N - 1 centers.
// O(n^2)
func expandAroundCenter(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j
}

func longestPalindrome(s string) string {
	var begin, end int
	for i := range s {
		begin1, end1 := expandAroundCenter(s, i, i)
		begin2, end2 := expandAroundCenter(s, i, i+1)
		if (end1 - begin1) > (end - begin) {
			begin, end = begin1, end1
		}

		if (end2 - begin2) > (end - begin) {
			begin, end = begin2, end2
		}
	}

	return s[begin:end]
}

// Method 2. Manacher's algorithm.
// 0(n)
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome2(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	// positions, including before first character
	// and after last character
	posLen := 2*n + 1
	lps := make([]int, posLen)

	lps[0], lps[1] = 0, 1

	C := 1 // centerPosition
	R := 2 // centerRightPosition

	// needed to recover the palindrome
	maxLen := 1 // longest palindrome length
	maxCenterPosition := 0

	// currentRightPosition
	for i := 2; i < posLen; i++ {
		mirror := 2*C - i // currentLeftPosition
		diff := R - i

		// if currentRightPosition i is within centerRightPosition R
		if diff > 0 {
			// palindrome is as long as the mirror palindrome contained
			// within center palindrome (not leaving the boundary)
			lps[i] = min(lps[mirror], diff)
		} else {
			lps[i] = 0
		}

		// odd positions correspond to the real letters in s, so need to compare characters
		// for even positions just increment
		for i+lps[i]+1 < posLen && i-lps[i] > 0 &&
			((i+lps[i]+1)%2 == 0 || s[(i+lps[i]+1)/2] == s[(i-lps[i]-1)/2]) {
			lps[i]++
		}

		// update longest palindromic substring
		if lps[i] > maxLen {
			maxLen = lps[i]
			maxCenterPosition = i
		}

		// adjust centerPosition if palindrome expansion happened
		if i+lps[i] > R {
			C = i
			R = i + lps[i]
		}
	}

	start := (maxCenterPosition - maxLen) / 2
	end := start + maxLen

	return s[start:end]
}

// Method 3. Longest common substring with 1D DP.
// first reverse string
func longestPalindrome3(s string) string {
	S := []byte(s)
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		S[i], S[j] = S[j], S[i]
	}
	rs := string(S)

	// need only one row and prev variable
	memo := make([]int, len(s)+1)

	maxLen := 0
	maxLenIdx := 0

	for i := 0; i <= len(s); i++ {
		prev := 0
		for j := 0; j <= len(rs); j++ {
			if i > 0 && j > 0 && s[i-1] == rs[j-1] {
				memo[j], prev = prev+1, memo[j]

				// check lengths [0:startI] == [endJ:len(s)]
				// against cases like "[abacd]fg[dcaba]"
				if maxLen < memo[j] && len(s)-j == i-memo[j] {
					maxLen = memo[j]
					maxLenIdx = j
				}
			} else {
				memo[j], prev = 0, memo[j]
			}
		}
	}

	return rs[maxLenIdx-maxLen : maxLenIdx]
}

// do not reverse string
func longestPalindrome33(s string) string {
	if len(s) == 0 {
		return s
	}

	memo := make([]int, len(s)+1)

	maxLen := 0
	maxLenIdx := 0

	for i := len(s) - 1; i >= -1; i-- {
		prev := 0
		for j := 0; j <= len(s); j++ {
			if i+1 < len(s) && j > 0 && s[i+1] == s[j-1] {
				memo[j], prev = prev+1, memo[j]

				if maxLen < memo[j] && j == i+memo[j] {
					maxLen = memo[j]
					maxLenIdx = j
				}
			} else {
				memo[j], prev = 0, memo[j]
			}
		}
	}

	return s[maxLenIdx-maxLen : maxLenIdx+1]
}
