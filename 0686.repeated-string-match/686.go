/**
 * Given two strings a and b, return the minimum number of times you should repeat string a so that string b is a substring of it. If it is impossible for b​​​​​​ to be a substring of a after repeating it, return -1.
 * Notice: string "abc" repeated 0 times is "",  repeated 1 time is "abc" and repeated 2 times is "abcabc".
 *
 * Example 1:
 *
 * Input: a = "abcd", b = "cdabcdab"
 * Output: 3
 * Explanation: We return 3 because by repeating a three times "abcdabcdabcd", b is a substring of it.
 *
 * Example 2:
 *
 * Input: a = "a", b = "aa"
 * Output: 2
 *
 * Example 3:
 *
 * Input: a = "a", b = "a"
 * Output: 1
 *
 * Example 4:
 *
 * Input: a = "abc", b = "wxyz"
 * Output: -1
 *
 *
 * Constraints:
 *
 * 	1 <= a.length <= 10^4
 * 	1 <= b.length <= 10^4
 * 	a and b consist of lower-case English letters.
 *
 */

package leetcode

import "bytes"

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

// The reason why we only need to check x and x+1 is the following.
// A can be written as s1+s2. The only way B becomes a substring when A is repeated
// is if B itself is a repetition of substrings in A in some order. If B = [s1+s2]*k
// then we only need to check the lower bound x=len(B)\len(A). If B=[s2+s1]*x then
// repeating A x times is not enough since we cannot fully recover the last repetition;
// e.g. we will be missing s1. Therefore, we need to do it x+1 times. For example, A=s1s2
// and B=s2s1s2s1. In this case, repeating A 2 times (lower bound) will become 2*A=s1s2s1s2.
// As you see, for us to recover B we need another A so we can append s1 to the last s2 above.
// Therefore, we only need to check x,x+1. If we still dont find B as a substring, then they dont
// follow the above form and hence the answer would be -1
func repeatedStringMatch(A string, B string) int {
	var str bytes.Buffer
	var count int
	for str.Len() < len(B) {
		str.WriteString(A)
		count++
	}

	idx := search(str.String(), B)
	if idx >= 0 {
		return count
	}

	str.WriteString(A)
	idx = search(str.String(), B)
	if idx >= 0 {
		return count + 1
	}
	return -1
}
