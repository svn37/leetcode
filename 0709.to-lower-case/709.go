/**
 * Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"Hello"</span>
 * Output: <span id="example-output-1">"hello"</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"here"</span>
 * Output: <span id="example-output-2">"here"</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"LOVELY"</span>
 * Output: <span id="example-output-3">"lovely"</span>
 *
 * </div>
 * </div>
 * </div>
 */

package leetcode

func toLowerCase(str string) string {
	S := []byte(str)
	for i := range S {
		if S[i] >= byte('A') && S[i] <= byte('Z') {
			S[i] += 32
		}
	}
	return string(S)
}
