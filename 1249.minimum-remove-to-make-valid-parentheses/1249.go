/**
 * Given a string <font face="monospace">s</font> of '(' , ')' and lowercase English characters.
 *
 * Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
 *
 * Formally, a parentheses string is valid if and only if:
 *
 *
 * 	It is the empty string, contains only lowercase characters, or
 * 	It can be written as AB (A concatenated with B), where A and B are valid strings, or
 * 	It can be written as (A), where A is a valid string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "lee(t(c)o)de)"
 * Output: "lee(t(c)o)de"
 * Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "a)b(c)d"
 * Output: "ab(c)d"
 *
 *
 * Example 3:
 *
 *
 * Input: s = "))(("
 * Output: ""
 * Explanation: An empty string is also valid.
 *
 *
 * Example 4:
 *
 *
 * Input: s = "(a(b(c)d)"
 * Output: "a(b(c)d)"
 *
 *
 *
 * Constraints:
 *
 *
 * 	1 <= s.length <= 10^5
 * 	s[i] is one of  '(' , ')' and lowercase English letters.
 *
 */

package leetcode

import "bytes"

type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() int {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func minRemoveToMakeValid(s string) string {
	stack := &Stack{}

	S := []byte(s)
	for i := range S {
		if S[i] == '(' {
			stack.Push(i)
		} else if S[i] == ')' {
			if stack.Empty() {
				S[i] = '*'
			} else {
				stack.Pop()
			}
		}
	}

	for !stack.Empty() {
		S[stack.Pop()] = '*'
	}

	var str bytes.Buffer
	for i := range S {
		if S[i] != '*' {
			str.WriteByte(S[i])
		}
	}
	return str.String()
}
