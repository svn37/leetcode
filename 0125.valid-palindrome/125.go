/**
 * Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 * Constraints:
 *
 * 	s consists only of printable ASCII characters.
 *
 */

package leetcode

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9')
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		leftChar := toLower(s[left])
		if !isAlphanumeric(leftChar) {
			left++
			continue
		}

		rightChar := toLower(s[right])
		if !isAlphanumeric(rightChar) {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}
