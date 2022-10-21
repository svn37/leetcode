/**
 * Given a string that contains only digits 0-9 and a target value, return all possibilities to add binary operators (not unary) +, -, or * between the digits so they evaluate to the target value.
 *
 * Example 1:
 *
 *
 * Input: num = "123", target = 6
 * Output: ["1+2+3", "1*2*3"]
 *
 *
 * Example 2:
 *
 *
 * Input: num = "232", target = 8
 * Output: ["2*3+2", "2+3*2"]
 *
 * Example 3:
 *
 *
 * Input: num = "105", target = 5
 * Output: ["1*0+5","10-5"]
 *
 * Example 4:
 *
 *
 * Input: num = "00", target = 0
 * Output: ["0+0", "0-0", "0*0"]
 *
 *
 * Example 5:
 *
 *
 * Input: num = "3456237490", target = 9191
 * Output: []
 *
 *
 * Constraints:
 *
 * 	0 <= num.length <= 10
 * 	num only contain digits.
 *
 */

package leetcode

import "fmt"

func backtrack(num string, target, idx, eval, prev int, expr string, res *[]string) {
	if idx == len(num) {
		if eval == target {
			*res = append(*res, expr)
		}
		return
	}

	val := 0
	for i := idx; i < len(num); i++ {
		if val == 0 && i != idx {
			break
		}
		val *= 10
		val += int(num[i] - '0')

		if idx == 0 {
			backtrack(num, target, i+1, val, val, num[:i+1], res)
		} else {
			backtrack(num, target, i+1, eval+val, val, fmt.Sprintf("%s+%s", expr, num[idx:i+1]), res)

			backtrack(num, target, i+1, eval-val, -val, fmt.Sprintf("%s-%s", expr, num[idx:i+1]), res)

			// for example, if you have a sequence of 12345 and you have proceeded to 1 + 2 + 3, now your eval is 6 right?
			// If you want to add a * between 3 and 4, you would take 3 as the digit to be multiplied, so you want to take
			// it out from the existing eval. You have 1 + 2 + 3 * 4 and the eval now is (1 + 2 + 3) - 3 + (3 * 4)
			backtrack(num, target, i+1, eval-prev+prev*val, prev*val, fmt.Sprintf("%s*%s", expr, num[idx:i+1]), res)
		}
	}
}

func addOperators(num string, target int) []string {
	res := []string{}
	if len(num) == 0 {
		return res
	}

	backtrack(num, target, 0, 0, 0, "", &res)
	return res
}
