/**
 * Given an input string s, reverse the order of the words.
 * A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
 * Return a string of the words in reverse order concatenated by a single space.
 * Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.
 *
 * Example 1:
 *
 * Input: s = "the sky is blue"
 * Output: "blue is sky the"
 *
 * Example 2:
 *
 * Input: s = "  hello world  "
 * Output: "world hello"
 * Explanation: Your reversed string should not contain leading or trailing spaces.
 *
 * Example 3:
 *
 * Input: s = "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
 *
 * Example 4:
 *
 * Input: s = "  Bob    Loves  Alice   "
 * Output: "Alice Loves Bob"
 *
 * Example 5:
 *
 * Input: s = "Alice does not even like bob"
 * Output: "bob like even not does Alice"
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^4
 * 	s contains English letters (upper-case and lower-case), digits, and spaces ' '.
 * 	There is at least one word in s.
 *
 *
 * Follow up:
 *
 * 	Could you solve it in-place with O(1) extra space?
 *
 *
 *
 */

package leetcode

import "bytes"

// []byte(s) is the only extra space.
func reduceSpaces(S []byte) []byte {
	S = bytes.TrimSpace(S)
	j := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ' ' && S[i+1] == ' ' {
			continue
		}
		S[j] = S[i]
		j++
	}
	return S[:j]
}

func reverse(S []byte) {
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		S[i], S[j] = S[j], S[i]
	}
}

func reverseWords(s string) string {
	S := []byte(s)
	S = reduceSpaces(S)
	reverse(S)

	// reverse each word
	prev := 0
	for i := 0; i <= len(S); i++ {
		if i == len(S) || S[i] == ' ' {
			reverse(S[prev:i])
			prev = i + 1
		}
	}
	return string(S)
}

// split s into words
func reverseWords2(s string) string {
	words := []string{}
	prev := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if i-prev > 0 {
				words = append(words, s[prev:i])
			}
			prev = i + 1
		}
	}
	if len(words) == 0 {
		return ""
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	var str bytes.Buffer
	for i := 0; i < len(words)-1; i++ {
		str.WriteString(words[i])
		str.WriteRune(' ')
	}
	str.WriteString(words[len(words)-1])
	return str.String()
}
