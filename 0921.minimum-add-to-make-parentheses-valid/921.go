/**
 * Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.
 *
 * Formally, a parentheses string is valid if and only if:
 *
 *
 * 	It is the empty string, or
 * 	It can be written as AB (A concatenated with B), where A and B are valid strings, or
 * 	It can be written as (A), where A is a valid string.
 *
 *
 * Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"())"</span>
 * Output: <span id="example-output-1">1</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"((("</span>
 * Output: <span id="example-output-2">3</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"()"</span>
 * Output: <span id="example-output-3">0</span>
 *
 *
 * <div>
 * Example 4:
 *
 *
 * Input: <span id="example-input-4-1">"()))(("</span>
 * Output: <span id="example-output-4">4</span>
 *
 *
 * </div>
 * </div>
 * </div>
 *
 * Note:
 *
 * <ol>
 * 	S.length <= 1000
 * 	S only consists of '(' and ')' characters.
 * </ol>
 *
 * <div>
 * <div>
 * <div>
 * <div> </div>
 * </div>
 * </div>
 * </div>
 */

package leetcode

func minAddToMakeValid(S string) int {
	forward := 0
	backward := 0
	for _, c := range S {
		if c == '(' {
			forward++
		} else {
			if forward == 0 {
				backward++
			} else {
				forward--
			}
		}
	}

	return forward + backward
}
