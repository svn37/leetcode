/**
 * Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 * Constraints:
 *
 * 	1 <= n <= 8
 *
 */

package leetcode

// Recursive, backtracking
func generate(fc, bc int, comb string, res *[]string) {
	if fc == 0 && bc == 0 {
		*res = append(*res, comb)
		return
	}

	if fc > 0 {
		generate(fc-1, bc, comb+"(", res)
	}

	if bc > fc {
		generate(fc, bc-1, comb+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	generate(n, n, "", &res)
	return res
}

// Iterative version.
// Consider the closure number of a valid parentheses sequence S: the least index >= 0 so that S[0], S[1], ..., S[2*index+1] is valid
// Clearly, every parentheses sequence has a unique closure number. We can try to enumerate them individually.
// For each closure number c, we know the starting and ending brackets must be at index 0 and 2*c + 1.
// Then, the 2*c elements between must be a valid sequence, plus the rest of the elements must be a valid sequence.
func generateParenthesis2(n int) []string {
	res := []string{}
	if n == 0 {
		res = append(res, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generateParenthesis(c) {
				for _, right := range generateParenthesis(n - c - 1) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
