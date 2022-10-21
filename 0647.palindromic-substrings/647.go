/**
 * Given a string, your task is to count how many palindromic substrings in this string.
 *
 * The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
 *
 * Example 1:
 *
 *
 * Input: "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The input string length won't exceed 1000.
 * </ol>
 *
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Manacher's algorithm
func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// positions, including before first character
	// and after last character
	posLen := 2*n + 1
	lps := make([]int, posLen)

	lps[0], lps[1] = 0, 1

	C := 1 // centerPosition
	R := 2 // centerRightPosition

	// palindromic substrings count
	count := 1

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

		// update palindromic substrings count
		count += (lps[i] + 1) / 2

		// adjust centerPosition if palindrome expansion happened
		if i+lps[i] > R {
			C = i
			R = i + lps[i]
		}
	}

	return count
}

// easy obvious method
func expandAroundCenter(s string, i, j int) int {
	count := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		count++
	}
	return count
}

func countSubstrings2(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += expandAroundCenter(s, i, i)
		count += expandAroundCenter(s, i, i+1)
	}
	return count
}
