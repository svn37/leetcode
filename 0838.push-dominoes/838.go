/**
 * There are<font face="monospace"> N</font> dominoes in a line, and we place each domino vertically upright.
 *
 * In the beginning, we simultaneously push some of the dominoes either to the left or to the right.
 *
 * <img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png" style="height: 160px;" />
 *
 * After each second, each domino that is falling to the left pushes the adjacent domino on the left.
 *
 * Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.
 *
 * When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.
 *
 * For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.
 *
 * Given a string "S" representing the initial state. S[i] = 'L', if the i-th domino has been pushed to the left; S[i] = 'R', if the i-th domino has been pushed to the right; S[i] = '.', if the i-th domino has not been pushed.
 *
 * Return a string representing the final state.
 *
 * Example 1:
 *
 *
 * Input: ".L.R...LR..L.."
 * Output: "LL.RR.LLRRLL.."
 *
 *
 * Example 2:
 *
 *
 * Input: "RR.L"
 * Output: "RR.L"
 * Explanation: The first domino expends no additional force on the second domino.
 *
 *
 * Note:
 *
 * <ol>
 * 	0 <= N <= 10^5
 * 	String dominoes contains only 'L', 'R' and '.'
 * </ol>
 *
 */

package leetcode

func pushDominoes(dominoes string) string {
	left, right, border := 0, -1, -1

	dom := []byte(dominoes)
	for i := 0; i <= len(dom); i++ {
		if i == len(dom) || dom[i] == 'R' {
			if right > border {
				for j := right; j < i; j++ {
					dom[j] = 'R'
				}
			}
			right = i + 1
		} else if dom[i] == 'L' {
			left = i - 1
			for right < left {
				if right > border {
					dom[right] = 'R'
					right++
				}
				dom[left] = 'L'
				left--
			}
			border = i
			right = i
		}
	}

	return string(dom)
}
