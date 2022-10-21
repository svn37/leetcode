/**
 *
 * Find the length of the longest substring T of a given string (consists of lowercase letters only) such that every character in T appears no less than k times.
 *
 *
 * Example 1:
 *
 * Input:
 * s = "aaabb", k = 3
 *
 * Output:
 * 3
 *
 * The longest substring is "aaa", as 'a' is repeated 3 times.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s = "ababbc", k = 2
 *
 * Output:
 * 5
 *
 * The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
 *
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// for c=1:26, we are going to use sliding window (left, right)
// to find the longest window which contains exactly c unique characters
// and for each character, there are at least K repeating ones.
func longestSubstring(s string, k int) int {
	if k < 2 {
		return len(s)
	}
	maxVal := 0

	for c := 1; c <= 26; c++ {
		chars := make([]int, 26)
		left, right := 0, 0
		unique, noLessThanK := 0, 0

		for right < len(s) {
			if unique <= c {
				idx := s[right] - 'a'
				if chars[idx] == 0 {
					unique++
				}
				chars[idx]++
				if chars[idx] == k {
					noLessThanK++
				}
				right++
			} else {
				idx := s[left] - 'a'
				if chars[idx] == k {
					noLessThanK--
				}
				chars[idx]--
				if chars[idx] == 0 {
					unique--
				}
				left++
			}

			if unique == c && unique == noLessThanK {
				maxVal = max(maxVal, right-left)
			}
		}
	}
	return maxVal
}
