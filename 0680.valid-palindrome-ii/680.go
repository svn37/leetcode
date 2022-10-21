/**
 *
 * Given a non-empty string s, you may delete at most one character.  Judge whether you can make it a palindrome.
 *
 *
 * Example 1:<br />
 *
 * Input: "aba"
 * Output: True
 *
 *
 *
 * Example 2:<br />
 *
 * Input: "abca"
 * Output: True
 * Explanation: You could delete the character 'c'.
 *
 *
 *
 * Note:<br>
 * <ol>
 * The string will only contain lowercase characters a-z.
 * The maximum length of the string is 50000.
 * </ol>
 *
 */

package leetcode

func valid(s string, misspell bool) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			if misspell {
				return false
			}
			return valid(s[i+1:j+1], true) || valid(s[i:j], true)
		}
	}
	return true
}

func validPalindrome(s string) bool {
	return valid(s, false)
}
