/**
 * You are given a string containing only 4 kinds of characters 'Q', 'W', 'E' and 'R'.
 * A string is said to be balanced if each of its characters appears n/4 times where n is the length of the string.
 * Return the minimum length of the substring that can be replaced with any other string of the same length to make the original string s balanced.
 * Return 0 if the string is already balanced.
 *
 * Example 1:
 *
 * Input: s = "QWER"
 * Output: 0
 * Explanation: s is already balanced.
 * Example 2:
 *
 * Input: s = "QQWE"
 * Output: 1
 * Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.
 *
 * Example 3:
 *
 * Input: s = "QQQW"
 * Output: 2
 * Explanation: We can replace the first "QQ" to "ER".
 *
 * Example 4:
 *
 * Input: s = "QQQQ"
 * Output: 3
 * Explanation: We can replace the last 3 'Q' to make s = "QWER".
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^5
 * 	s.length is a multiple of 4
 * 	s contains only 'Q', 'W', 'E' and 'R'.
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func balancedString(s string) int {
	count := make([]int, 26)
	for _, char := range s {
		count[char-'A']++
	}

	res := len(s)
	k := len(s) / 4
	// sliding window
	// if we need to find substring, it should be the first guess
	for left, right := 0, 0; right < len(s); right++ {
		count[s[right]-'A']--

		// under these 'count' conditions we have the possible substring
		// however, it is not necessarily minimal
		// <= is because we decrease count in the line above, but by definition
		// it is guaranteed that such string will be possible
		// s.length is a multiple of 4 && each of its characters appears n/4 times
		for left < len(s) && count['Q'-'A'] <= k && count['W'-'A'] <= k && count['E'-'A'] <= k && count['R'-'A'] <= k {
			res = min(res, right-left+1)
			count[s[left]-'A']++
			left++
		}
	}
	return res
}
