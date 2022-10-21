/**
 *
 * Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:
 * <ol>
 * Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 * Any right parenthesis ')' must have a corresponding left parenthesis '('.
 * Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 * '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
 * An empty string is also valid.
 * </ol>
 *
 *
 * Example 1:<br />
 *
 * Input: "()"
 * Output: True
 *
 *
 *
 * Example 2:<br />
 *
 * Input: "(*)"
 * Output: True
 *
 *
 *
 * Example 3:<br />
 *
 * Input: "(*))"
 * Output: True
 *
 *
 *
 * Note:<br>
 * <ol>
 * The string size will be in the range [1, 100].
 * </ol>
 *
 */

package leetcode

// Method 1. Slow, stupid recursive.
func check(S string, start, counter int) bool {
	alts := false
	for i := start; i < len(S); i++ {
		if S[i] == '(' {
			counter++
		} else if S[i] == ')' {
			counter--
		} else if S[i] == '*' {
			alts = alts || check(S, i+1, counter+1) || check(S, i+1, counter-1)
		}

		if counter < 0 || alts {
			break
		}
	}
	return counter == 0 || alts
}

func checkValidString(S string) bool {
	return check(S, 0, 0)
}

// Method 2. Count bounds (explained below)
func checkValidString2(S string) bool {
	// keep the balance
	// each * opens 3 paths
	// but we should only keep track of thelower and upper bounds
	// if the upper bound is negative, that means all routes have more ‘)’ than ‘(’ --> all routes are invalid
	// if the lower bound is 0 at the very end --> at least path gave that zero
	low := 0
	high := 0

	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			low++
			high++
		} else if S[i] == ')' {
			if low > 0 {
				low--
			}
			high--
		} else {
			if low > 0 {
				low--
			}
			high++
		}
		if high < 0 {
			return false
		}
	}
	return low == 0
}
