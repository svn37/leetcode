/**
 * Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.
 *
 * You may return the answer in any order.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">["bella","label","roller"]</span>
 * Output: <span id="example-output-1">["e","l","l"]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">["cool","lock","cook"]</span>
 * Output: <span id="example-output-2">["c","o"]</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	1 <= A.length <= 100
 * 	1 <= A[i].length <= 100
 * 	A[i][j] is a lowercase letter
 * </ol>
 * </div>
 * </div>
 */

package leetcode

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonChars(A []string) []string {
	res := []string{}
	chars := make([]int, 26)
	for i := range chars {
		chars[i] = math.MaxInt64
	}

	for _, str := range A {
		count := make([]int, 26)
		for j := range str {
			count[str[j]-'a']++
		}
		for j := range count {
			chars[j] = min(chars[j], count[j])
		}
	}
	for i := range chars {
		for j := 0; j < chars[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}
