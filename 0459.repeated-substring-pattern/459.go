/**
 * Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abab"
 * Output: True
 * Explanation: It's the substring "ab" twice.
 *
 *
 * Example 2:
 *
 *
 * Input: "aba"
 * Output: False
 *
 *
 * Example 3:
 *
 *
 * Input: "abcabcabcabc"
 * Output: True
 * Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
 *
 *
 */

package leetcode

// computes longest proper prefix which is also suffix array
func computeLPSArray(pattern string) []int {
	// first index is 0
	lps := make([]int, len(pattern))

	// l -- length of the previous longest prefix suffix
	i, l := 1, 0
	for i < len(pattern) {
		if pattern[i] == pattern[l] {
			l++
			lps[i] = l
			i++
		} else if l != 0 {
			// tricky case - AAACAAAA and i = 7. s[7] = 3, because s[3-1] = 3
			l = lps[l-1]
		} else {
			lps[i] = 0
			i++
		}
	}

	return lps
}

// knuth-morris-pratt algorithm
func search(text string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	} else if len(text) == 0 {
		return -1
	}

	lps := computeLPSArray(pattern)

	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		} else if j != 0 {
			j = lps[j-1]
		} else {
			i++
		}

		if j == len(pattern) {
			// first index of the occurrence
			return i - j
		}
	}

	return -1
}

func repeatedSubstringPattern(s string) bool {
	ss := (s + s)[1 : 2*len(s)-1]
	return search(ss, s) != -1
}
