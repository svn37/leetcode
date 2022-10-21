/**
 * Given a string text, we are allowed to swap two of the characters in the string. Find the length of the longest substring with repeated characters.
 *
 * Example 1:
 *
 * Input: text = "ababa"
 * Output: 3
 * Explanation: We can swap the first 'b' with the last 'a', or the last 'b' with the first 'a'. Then, the longest repeated character substring is "aaa", which its length is 3.
 *
 * Example 2:
 *
 * Input: text = "aaabaaa"
 * Output: 6
 * Explanation: Swap 'b' with the last 'a' (or the first 'a'), and we get longest repeated character substring "aaaaaa", which its length is 6.
 *
 * Example 3:
 *
 * Input: text = "aaabbaaa"
 * Output: 4
 *
 * Example 4:
 *
 * Input: text = "aaaaa"
 * Output: 5
 * Explanation: No need to swap, longest repeated character substring is "aaaaa", length is 5.
 *
 * Example 5:
 *
 * Input: text = "abcdef"
 * Output: 1
 *
 *
 * Constraints:
 *
 * 	1 <= text.length <= 20000
 * 	text consist of lowercase English characters only.
 *
 */

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Simple sliding window.
// Note that we could group characters such that
// [[1,3],[2,4]...] - each element is index of character
// as a performance optimization
func maxRepOpt1(text string) int {
	count := make([]int, 26)
	for _, char := range text {
		count[char-'a']++
	}

	res := 0
	// detect the longest substring (with up to one different character)
	for char := 'a'; char <= 'z'; char++ {
		i, j, gap := 0, 0, 0
		for ; i < len(text); i++ {
			if text[i] != byte(char) {
				gap++
			}
			if gap > 1 {
				if text[j] != byte(char) {
					gap--
				}
				j++
			}
		}
		res = max(res, min(i-j, count[char-'a']))
	}
	return res
}
