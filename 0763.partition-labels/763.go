/**
 * A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.
 *
 * Example 1:
 *
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
 *
 *
 * Note:
 *
 * 	S will have length in range [1, 500].
 * 	S will consist of lowercase English letters ('a' to 'z') only.
 *
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

func partitionLabels(S string) []int {
	// record last index of each character
	chars := make([]int, 26)
	for i, char := range S {
		chars[char-'a'] = i
	}

	result := make([]int, 0)
	left, right := 0, 0

	for i, char := range S {
		right = max(right, chars[char-'a'])
		if i == right {
			result = append(result, right-left+1)
			left = i + 1
		}
	}
	return result
}
