/**
 * Given two strings s and t, return the minimum window in s which will contain all the characters in t. If there is no such window in s that covers all characters in t, return the empty string "".
 * Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.
 *
 * Example 1:
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Example 2:
 * Input: s = "a", t = "a"
 * Output: "a"
 *
 * Constraints:
 *
 * 	1 <= s.length, t.length <= 10^5
 * 	s and t consist of English letters.
 *
 *
 * Follow up: Could you find an algorithm that runs in O(n) time?
 */

package leetcode

func minWindow(s string, t string) string {
	chars := make(map[byte]int)
	for i := range t {
		chars[t[i]]++
	}

	// len(s)+1 instead of math.MaxInt64
	start, end := 0, len(s)+1
	counter := len(t)

	left, right := 0, 0

	for right < len(s) {
		if chars[s[right]] > 0 {
			counter--
		}
		chars[s[right]]--
		right++

		for counter == 0 {
			if right-left < end-start {
				start, end = left, right
			}

			chars[s[left]]++
			if chars[s[left]] > 0 {
				counter++
			}
			left++
		}
	}

	if end == len(s)+1 {
		return ""
	}
	return s[start:end]
}
