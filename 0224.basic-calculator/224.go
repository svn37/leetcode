/**
 * Implement a basic calculator to evaluate a simple expression string.
 *
 * The expression string may contain open ( and closing parentheses ), the plus + or minus sign -, non-negative integers and empty spaces  .
 *
 * Example 1:
 *
 *
 * Input: "1 + 1"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: " 2-1 + 2 "
 * Output: 3
 *
 * Example 3:
 *
 *
 * Input: "(1+(4+5+2)-3)+(6+8)"
 * Output: 23
 * Note:
 *
 *
 * 	You may assume that the given expression is always valid.
 * 	Do not use the eval built-in library function.
 *
 *
 */

package leetcode

func calculate(s string) int {
	answer, num, foundDigit := 0, 0, false
	signs := []int{1, 1}

	s += " " // add one more iteration
	for i := 0; i < len(s); i++ {
		for s[i] >= '0' && s[i] <= '9' {
			foundDigit = true
			num = num*10 + int(s[i]-'0')
			i++
		}

		last := len(signs) - 1
		if foundDigit {
			answer += signs[last] * num
			signs = signs[:last]
			num, last, foundDigit = 0, last-1, false
		}

		if s[i] == '+' || s[i] == '(' {
			signs = append(signs, signs[last])
		} else if s[i] == '-' {
			signs = append(signs, signs[last]*(-1))
		} else if s[i] == ')' {
			signs = signs[:last]
		}
	}
	return answer
}
