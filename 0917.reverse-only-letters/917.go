/**
 * Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.
 *
 *
 *
 * <div>
 * <div>
 * <div>
 * <ol>
 * </ol>
 * </div>
 * </div>
 * </div>
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"ab-cd"</span>
 * Output: <span id="example-output-1">"dc-ba"</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"a-bC-dEf-ghIj"</span>
 * Output: <span id="example-output-2">"j-Ih-gfE-dCba"</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"Test1ng-Leet=code-Q!"</span>
 * Output: <span id="example-output-3">"Qedo1ct-eeLg=ntse-T!"</span>
 *
 *
 *
 *
 * <div>
 * <span>Note:</span>
 *
 * <ol>
 * 	S.length <= 100
 * 	33 <= S[i].ASCIIcode <= 122
 * 	S doesn't contain \ or "
 * </ol>
 * </div>
 * </div>
 * </div>
 * </div>
 */

package leetcode

func notLetter(b byte) bool {
	return b < 65 || (b > 90 && b < 97)
}

func reverseOnlyLetters(str string) string {
	S := []byte(str)
	left, right := 0, len(S)-1
	for left < right {
		if notLetter(S[left]) {
			left++
		} else if notLetter(S[right]) {
			right--
		} else {
			S[left], S[right] = S[right], S[left]
			left++
			right--
		}
	}
	return string(S)
}
