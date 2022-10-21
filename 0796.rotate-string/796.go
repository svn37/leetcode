/**
 * We are given two strings, A and B.
 *
 * A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example, if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if A can become B after some number of shifts on A.
 *
 *
 * Example 1:
 * Input: A = 'abcde', B = 'cdeab'
 * Output: true
 *
 * Example 2:
 * Input: A = 'abcde', B = 'abced'
 * Output: false
 *
 *
 * Note:
 *
 *
 * 	A and B will have length at most 100.
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

func KMP(text string, pattern string) int {
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

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	return KMP(A+A, B) != -1
}
