/**
 * Implement a basic calculator to evaluate a simple expression string.
 *
 * The expression string contains only non-negative integers, +, -, *, / operators and empty spaces  . The integer division should truncate toward zero.
 *
 * Example 1:
 *
 *
 * Input: "3+2*2"
 * Output: 7
 *
 *
 * Example 2:
 *
 *
 * Input: " 3/2 "
 * Output: 1
 *
 * Example 3:
 *
 *
 * Input: " 3+5 / 2 "
 * Output: 5
 *
 *
 * Note:
 *
 *
 * 	You may assume that the given expression is always valid.
 * 	Do not use the eval built-in library function.
 *
 *
 */

package leetcode

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

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func calculate(s string) int {
	if len(s) == 0 {
		return -1
	}

	stack := &Stack{}
	sign := byte('+')
	val := 0

	for i := 0; i < len(s); i++ {
		// skip spaces
		// unless it is the last character,
		// then we need to push the last digit to the stack
		if s[i] == ' ' && i != len(s)-1 {
			continue
		}

		if s[i] >= 48 && s[i] <= 57 {
			val = (val * 10) + int(s[i]-'0')
			if i != len(s)-1 {
				continue
			}
		}

		if sign == '+' {
			stack.Push(val)
		} else if sign == '-' {
			stack.Push(-val)
		} else if sign == '*' {
			stack.Push(stack.Pop() * val)
		} else if sign == '/' {
			stack.Push(stack.Pop() / val)
		}

		sign = s[i]
		val = 0
	}

	for !stack.Empty() {
		val += stack.Pop()
	}
	return val
}
