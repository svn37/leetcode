/**
 * Given a pattern and a string s, find if s follows the same pattern.
 * Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
 *
 * Example 1:
 *
 * Input: pattern = "abba", s = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 * Input: pattern = "abba", s = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 * Input: pattern = "aaaa", s = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 * Input: pattern = "abba", s = "dog dog dog dog"
 * Output: false
 *
 *
 * Constraints:
 *
 * 	1 <= pattern.length <= 300
 * 	pattern contains only lower-case English letters.
 * 	1 <= s.length <= 3000
 * 	s contains only lower-case English letters and spaces ' '.
 * 	s does not contain any leading or trailing spaces.
 * 	All the words in s are separated by a single space.
 *
 */

package leetcode

func wordPattern(pattern string, str string) bool {
	chars := make(map[byte]int)
	words := make(map[string]int)

	i, j, k := 0, 0, 0
	for i < len(pattern) && j < len(str) {
		start := j
		for j < len(str) && str[j] != ' ' {
			j++
		}
		word := str[start:j]
		char := pattern[i]

		if chars[char] != words[word] {
			return false
		}
		chars[char] = i + 1
		words[word] = k + 1

		i++
		j++
		k++
	}

	// j-1 because we usually move pointer forward  to skip the space
	return i == len(pattern) && j-1 == len(str)
}
