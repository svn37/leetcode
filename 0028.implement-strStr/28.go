/**
 * Implement <a href="http://www.cplusplus.com/reference/cstring/strstr/" target="_blank">strStr()</a>.
 * Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
 * Clarification:
 * What should we return when needle is an empty string? This is a great question to ask during an interview.
 * For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's <a href="http://www.cplusplus.com/reference/cstring/strstr/" target="_blank">strstr()</a> and Java's <a href="https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)" target="_blank">indexOf()</a>.
 *
 * Example 1:
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * Example 2:
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 * Example 3:
 * Input: haystack = "", needle = ""
 * Output: 0
 *
 * Constraints:
 *
 * 	0 <= haystack.length, needle.length <= 5 * 10^4
 * 	haystack and needle consist of only lower-case English characters.
 *
 */

package leetcode

// Method 1. KMP algorithm.
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

func strStr(text string, pattern string) int {
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

// Method 2. Rabin-Karp algorithm.
const base = 16777619

func hash(s string) uint32 {
	var h uint32
	for i := range s {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func rabinKarp(text, pattern string) int {
	n, m := len(text), len(pattern)

	if m == 0 {
		return 0
	} else if n == 0 || n < m {
		return -1
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hash(pattern)
	h := hash(text[:m])

	for i := 0; i < n-m+1; i++ {
		if i > 0 {
			h = h - mult*uint32(text[i-1])
			h = h*base + uint32(text[i+m-1])
		}

		if h == hp && text[i:i+m] == pattern {
			return i
		}
	}
	return -1
}

func strStr2(text string, pattern string) int {
	return rabinKarp(text, pattern)
}

// Method 3. Z algorithm.
// (here the generic version, or we can stop building Z when the first match is found)
func computeZ(s string) []int {
	Z := make([]int, len(s))

	// [L,R] -- Z-box, matches with prefix
	L, R := 0, 0
	for i := 1; i < len(s); i++ {
		// nothing matches, calculate manually
		if i > R {
			L = i
			R = i

			// R-L is prefix position, R is Z-box boundary
			for R < len(s) && s[R-L] == s[R] {
				R++
			}
			Z[i] = R - L
			R--
		} else {
			// k is prefix position
			k := i - L

			// if Z[k] is within the current Z-box, Z[i] = Z[k]
			if Z[k] < R-i+1 {
				Z[i] = Z[k]
				// otherwise calculate manually
			} else {
				L = i

				for R < len(s) && s[R-L] == s[R] {
					R++
				}
				Z[i] = R - L
				R--
			}
		}
	}
	return Z
}

func strStr3(text string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	} else if len(text) == 0 || len(text) < len(pattern) {
		return -1
	}

	concat := pattern + "$" + text

	Z := computeZ(concat)

	for i := range concat {
		if Z[i] == len(pattern) {
			// get rid of pattern + sentinel prefix
			return i - len(pattern) - 1
		}
	}
	return -1
}
