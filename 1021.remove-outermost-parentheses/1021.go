/**
 * A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.
 *
 * A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.
 *
 * Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.
 *
 * Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"(()())(())"</span>
 * Output: <span id="example-output-1">"()()()"</span>
 * Explanation:
 * The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
 * After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"(()())(())(()(()))"</span>
 * Output: <span id="example-output-2">"()()()()(())"</span>
 * Explanation:
 * The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
 * After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"()()"</span>
 * Output: <span id="example-output-3">""</span>
 * Explanation:
 * The input string is "()()", with primitive decomposition "()" + "()".
 * After removing outer parentheses of each part, this is "" + "" = "".
 *
 *
 *
 * </div>
 * </div>
 *
 * Note:
 *
 * <ol>
 * 	S.length <= 10000
 * 	S[i] is "(" or ")"
 * 	S is a valid parentheses string
 * </ol>
 *
 * <div>
 * <div>
 * <div> </div>
 * </div>
 * </div>
 */

package leetcode

import "bytes"

func removeOuterParentheses(S string) string {
	var str bytes.Buffer

	par := 0
	start := 0
	for i, char := range S {
		if char == '(' {
			par++
		} else {
			par--
		}

		if par == 0 {
			str.WriteString(S[start+1 : i])
			start = i + 1
		}
	}

	return str.String()
}
