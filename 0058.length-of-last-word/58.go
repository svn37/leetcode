/**
 * Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word (last word means the last appearing word if we loop from left to right) in the string.
 * If the last word does not exist, return 0.
 * Note: A word is defined as a maximal substring consisting of non-space characters only.
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 */

package leetcode

func lengthOfLastWord(s string) int {
	// trim the string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s = s[:i]
		} else {
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s[i+1:])
		}
	}
	return len(s)
}
