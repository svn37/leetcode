/**
 * Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
 * Note that after backspacing an empty text, the text will continue empty.
 * <div>
 * Example 1:
 *
 * Input: S = <span id="example-input-1-1">"ab#c"</span>, T = <span id="example-input-1-2">"ad#c"</span>
 * Output: <span id="example-output-1">true
 * </span><span>Explanation: Both S and T become "ac".</span>
 *
 * <div>
 * Example 2:
 *
 * Input: S = <span id="example-input-2-1">"ab##"</span>, T = <span id="example-input-2-2">"c#d#"</span>
 * Output: <span id="example-output-2">true
 * </span><span>Explanation: Both S and T become "".</span>
 *
 * <div>
 * Example 3:
 *
 * Input: S = <span id="example-input-3-1">"a##c"</span>, T = <span id="example-input-3-2">"#a#c"</span>
 * Output: <span id="example-output-3">true
 * </span><span>Explanation: Both S and T become "c".</span>
 *
 * <div>
 * Example 4:
 *
 * Input: S = <span id="example-input-4-1">"a#c"</span>, T = <span id="example-input-4-2">"b"</span>
 * Output: <span id="example-output-4">false
 * </span><span>Explanation: S becomes "c" while T becomes "b".</span>
 *
 * <span>Note:</span>
 *
 * 	<span>1 <= S.length <= 200</span>
 * 	<span>1 <= T.length <= 200</span>
 * 	<span>S and T only contain lowercase letters and '#' characters.</span>
 *
 * Follow up:
 *
 * 	Can you solve it in O(N) time and O(1) space?
 * </div>
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

func nextIdx(str string, k int) int {
	skip := 0
	for i := k; i >= 0; i-- {
		if str[i] == '#' {
			skip++
		} else {
			if skip > 0 {
				skip--
			} else {
				return i
			}
		}
	}

	return -1
}

func backspaceCompare(S string, T string) bool {
	i := len(S)
	j := len(T)

	for {
		i = nextIdx(S, i-1)
		j = nextIdx(T, j-1)

		if i < 0 && j < 0 {
			return true
		}

		if i < 0 || j < 0 || S[i] != T[j] {
			return false
		}
	}
}
