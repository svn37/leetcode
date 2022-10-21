/**
 * Given a string s consisting only of characters a, b and c.
 * Return the number of substrings containing at least one occurrence of all these characters a, b and c.
 *
 * Example 1:
 *
 * Input: s = "abcabc"
 * Output: 10
 * Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).
 *
 * Example 2:
 *
 * Input: s = "aaacb"
 * Output: 3
 * Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb".
 *
 * Example 3:
 *
 * Input: s = "abc"
 * Output: 1
 *
 *
 * Constraints:
 *
 * 	3 <= s.length <= 5 x 10^4
 * 	s only consists of a, b or c characters.
 *
 */

package leetcode

func numberOfSubstrings(s string) int {
	count := 0
	left, right := 0, 0

	one, two, three := 0, 0, 0
	for right < len(s) {
		if s[right] == 'a' {
			one++
		} else if s[right] == 'b' {
			two++
		} else if s[right] == 'c' {
			three++
		}

		for one > 0 && two > 0 && three > 0 && left < right {
			count += len(s) - right

			if s[left] == 'a' {
				one--
			} else if s[left] == 'b' {
				two--
			} else if s[left] == 'c' {
				three--
			}
			left++
		}
		right++
	}
	return count
}
