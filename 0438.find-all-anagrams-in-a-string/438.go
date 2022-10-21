/**
 * Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
 *
 * Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
 *
 * The order of output does not matter.
 *
 * Example 1:
 *
 * Input:
 * s: "cbaebabacd" p: "abc"
 *
 * Output:
 * [0, 6]
 *
 * Explanation:
 * The substring with start index = 0 is "cba", which is an anagram of "abc".
 * The substring with start index = 6 is "bac", which is an anagram of "abc".
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s: "abab" p: "ab"
 *
 * Output:
 * [0, 1, 2]
 *
 * Explanation:
 * The substring with start index = 0 is "ab", which is an anagram of "ab".
 * The substring with start index = 1 is "ba", which is an anagram of "ab".
 * The substring with start index = 2 is "ab", which is an anagram of "ab".
 *
 *
 */

package leetcode

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}

	m := make(map[byte]int)
	for i := range p {
		char := p[i]
		m[char]++
	}

	// num of characters left to form an anagram
	left := len(m)

	for start, end := -len(p), 0; end < len(s); start, end = start+1, end+1 {
		if start >= 0 {
			startChar := s[start]
			if count, ok := m[startChar]; ok {
				if count == 0 {
					left++
				}
				m[startChar]++
			}
		}

		endChar := s[end]
		if count, ok := m[endChar]; ok {
			if count <= 1 {
				if count == 1 {
					left--
				}
				if left == 0 {
					res = append(res, start+1)
				}
			}
			m[endChar]--
		}
	}
	return res
}
