/**
 * Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.
 *
 * Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:
 *
 *
 * 	If S[i] == "I", then A[i] < A[i+1]
 * 	If S[i] == "D", then A[i] > A[i+1]
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"IDID"</span>
 * Output: <span id="example-output-1">[0,4,1,3,2]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"III"</span>
 * Output: <span id="example-output-2">[0,1,2,3]</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"DDI"</span>
 * Output: <span id="example-output-3">[3,2,0,1]</span>
 * </div>
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= S.length <= 10000
 * 	S only contains characters "I" or "D".
 * </ol>
 */

package leetcode

func diStringMatch(S string) []int {
	minNum := 0
	maxNum := len(S)
	res := make([]int, 0, len(S)+1)

	for i := range S {
		if S[i] == 'I' {
			res = append(res, minNum)
			minNum++
		} else if S[i] == 'D' {
			res = append(res, maxNum)
			maxNum--
		}
	}
	return append(res, maxNum)
}
