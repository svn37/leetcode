/**
 * You are given a string s that consists of lower case English letters and brackets.
 * Reverse the strings in each pair of matching parentheses, starting from the innermost one.
 * Your result should not contain any brackets.
 *
 * Example 1:
 *
 * Input: s = "(abcd)"
 * Output: "dcba"
 *
 * Example 2:
 *
 * Input: s = "(u(love)i)"
 * Output: "iloveu"
 * Explanation: The substring "love" is reversed first, then the whole string is reversed.
 *
 * Example 3:
 *
 * Input: s = "(ed(et(oc))el)"
 * Output: "leetcode"
 * Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.
 *
 * Example 4:
 *
 * Input: s = "a(bcdefghijkl(mno)p)q"
 * Output: "apmnolkjihgfedcbq"
 *
 *
 * Constraints:
 *
 * 	0 <= s.length <= 2000
 * 	s only contains lower case English characters and parentheses.
 * 	It's guaranteed that all parentheses are balanced.
 *
 */

package leetcode

import "bytes"

type Stack struct {
	storage []int
}

func (s *Stack) Push(val int) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Pop() int {
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func reverseParentheses(s string) string {
	stack := &Stack{}
	m := make(map[int]int)

	for i, char := range s {
		if char == '(' {
			stack.Push(i)
		} else if char == ')' {
			j := stack.Pop()
			m[i] = j
			m[j] = i
		}
	}

	var b bytes.Buffer
	for i, d := 0, 1; i < len(s); i += d {
		if s[i] == '(' || s[i] == ')' {
			i = m[i]
			d *= -1
		} else {
			b.WriteByte(s[i])
		}
	}

	return b.String()
}
