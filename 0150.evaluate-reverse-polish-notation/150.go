/**
 * Evaluate the value of an arithmetic expression in <a href="http://en.wikipedia.org/wiki/Reverse_Polish_notation" target="_blank">Reverse Polish Notation</a>.
 *
 * Valid operators are +, -, *, /. Each operand may be an integer or another expression.
 *
 * Note:
 *
 *
 * 	Division between two integers should truncate toward zero.
 * 	The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
 *
 *
 * Example 1:
 *
 *
 * Input: ["2", "1", "+", "3", "*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 *
 * Example 2:
 *
 *
 * Input: ["4", "13", "5", "/", "+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 *
 * Example 3:
 *
 *
 * Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * Output: 22
 * Explanation:
 *   ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 */

package leetcode

import "strconv"

type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() int {
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func evalRPN(tokens []string) int {
	stack := &Stack{}
	for i := range tokens {
		if tokens[i] == "+" {
			stack.Push(stack.Pop() + stack.Pop())
		} else if tokens[i] == "-" {
			second := stack.Pop()
			first := stack.Pop()

			stack.Push(first - second)
		} else if tokens[i] == "*" {
			stack.Push(stack.Pop() * stack.Pop())
		} else if tokens[i] == "/" {
			second := stack.Pop()
			first := stack.Pop()

			stack.Push(first / second)
		} else {
			val, _ := strconv.Atoi(tokens[i])
			stack.Push(val)
		}
	}
	return stack.Pop()
}
