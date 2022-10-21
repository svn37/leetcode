/**
 * Remove the minimum number of invalid parentheses in order to make the input string valid. Return all possible results.
 *
 * Note: The input string may contain letters other than the parentheses ( and ).
 *
 * Example 1:
 *
 *
 * Input: "()())()"
 * Output: ["()()()", "(())()"]
 *
 *
 * Example 2:
 *
 *
 * Input: "(a)())()"
 * Output: ["(a)()()", "(a())()"]
 *
 *
 * Example 3:
 *
 *
 * Input: ")("
 * Output: [""]
 *
 */

package leetcode

func remove(str string, li, lj int, par string, res *[]string) {
	stack := 0
	// i is always greater or equal than j
	// j is the previous deletion position
	// before li string is guaranteed to be valid
	for i := li; i < len(str); i++ {
		if str[i] == par[0] {
			stack++
		}
		if str[i] == par[1] {
			stack--
		}

		// as soon as string becomes invalid
		if stack < 0 {
			for j := lj; j <= i; j++ {
				// in a series of consecutive ), or (, delete the first
				if str[j] == par[1] && (j == lj || str[j-1] != par[1]) {
					remove(str[:j]+str[j+1:], i, j, par, res)
				}
			}
			return
		}
	}
	S := []byte(str)
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		S[i], S[j] = S[j], S[i]
	}
	reversed := string(S)
	// (())() -> reverse )())((
	// can reuse the code, but par is )(

	// what happens here is:
	// if string starts with (, we reverse it for the first time (it must be valid)
	// if string starts with ), it is the second reversal, and one of the results is ready
	if par[0] == '(' {
		remove(reversed, 0, 0, ")(", res)
	} else {
		*res = append(*res, reversed)
	}
}

func removeInvalidParentheses(str string) []string {
	res := []string{}
	remove(str, 0, 0, "()", &res)

	return res
}
