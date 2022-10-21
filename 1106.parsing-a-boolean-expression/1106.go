/**
 * Return the result of evaluating a given boolean expression, represented as a string.
 * An expression can either be:
 *
 * 	"t", evaluating to True;
 * 	"f", evaluating to False;
 * 	"!(expr)", evaluating to the logical NOT of the inner expression expr;
 * 	"&amp;(expr1,expr2,...)", evaluating to the logical AND of 2 or more inner expressions expr1, expr2, ...;
 * 	"|(expr1,expr2,...)", evaluating to the logical OR of 2 or more inner expressions expr1, expr2, ...
 *
 *
 * Example 1:
 *
 * Input: expression = "!(f)"
 * Output: true
 *
 * Example 2:
 *
 * Input: expression = "|(f,t)"
 * Output: true
 *
 * Example 3:
 *
 * Input: expression = "&amp;(t,f)"
 * Output: false
 *
 * Example 4:
 *
 * Input: expression = "|(&amp;(t,f,t),!(t))"
 * Output: false
 *
 *
 * Constraints:
 *
 * 	1 <= expression.length <= 20000
 * 	expression[i] consists of characters in {'(', ')', '&amp;', '|', '!', 't', 'f', ','}.
 * 	expression is a valid expression representing a boolean, as given in the description.
 *
 */

package leetcode

func parse(expr []byte, lo, hi int) bool {
	char := rune(expr[lo])
	if hi-lo == 1 {
		// base case
		return char == 't'
	}
	res := char == '&'
	level := 0
	start := lo + 2

	for i := lo + 2; i < hi; i++ {
		c := expr[i]
		if level == 0 && (c == ',' || c == ')') {
			cur := parse(expr, start, i)
			start = i + 1
			if char == '&' {
				res = res && cur
			} else if char == '|' {
				res = res || cur
			} else {
				res = !cur
			}
		}
		if c == '(' {
			level++
		} else if c == ')' {
			level--
		}
	}
	return res
}

func parseBoolExpr(expr string) bool {
	return parse([]byte(expr), 0, len(expr))
}
