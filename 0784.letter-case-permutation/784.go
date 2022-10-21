/**
 * Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.
 * Return a list of all possible strings we could create. You can return the output in any order.
 *
 * Example 1:
 *
 * Input: S = "a1b2"
 * Output: ["a1b2","a1B2","A1b2","A1B2"]
 *
 * Example 2:
 *
 * Input: S = "3z4"
 * Output: ["3z4","3Z4"]
 *
 * Example 3:
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 * Example 4:
 *
 * Input: S = "0"
 * Output: ["0"]
 *
 *
 * Constraints:
 *
 * 	S will be a string with length between 1 and 12.
 * 	S will consist only of letters or digits.
 *
 */

package leetcode

func backtrack(S []byte, i int, res *[]string) {
	if i == len(S) {
		*res = append(*res, string(S))
		return
	}
	backtrack(S, i+1, res)
	if S[i] >= 65 && S[i] <= 90 {
		S[i] += ' '
		backtrack(S, i+1, res)
	} else if S[i] >= 97 {
		S[i] -= ' '
		backtrack(S, i+1, res)
	}
}

func letterCasePermutation(str string) []string {
	res := []string{}
	S := []byte(str)
	backtrack(S, 0, &res)

	return res
}
