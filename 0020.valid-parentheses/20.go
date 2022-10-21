/**
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 * An input string is valid if:
 * <ol>
 * 	Open brackets must be closed by the same type of brackets.
 * 	Open brackets must be closed in the correct order.
 * </ol>
 *
 * Example 1:
 *
 * Input: s = "()"
 * Output: true
 *
 * Example 2:
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 * Example 3:
 *
 * Input: s = "(]"
 * Output: false
 *
 * Example 4:
 *
 * Input: s = "([)]"
 * Output: false
 *
 * Example 5:
 *
 * Input: s = "{[]}"
 * Output: true
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 10^4
 * 	s consists of parentheses only '()[]{}'.
 *
 */

package leetcode

// stack implemented as a singly linked list
type Stack struct {
	top *Node
}

type Node struct {
	val  rune
	prev *Node
}

func (s *Stack) Push(val rune) {
	s.top = &Node{
		val:  val,
		prev: s.top,
	}
}

func (s *Stack) Pop() rune {
	pop := s.top.val
	s.top = s.top.prev
	return pop
}

func (s *Stack) Empty() bool {
	return s.top == nil
}

func isValid(s string) bool {
	stack := &Stack{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.Push(c)
		case ')':
			if stack.Empty() || stack.Pop() != '(' {
				return false
			}
		case ']':
			if stack.Empty() || stack.Pop() != '[' {
				return false
			}
		case '}':
			if stack.Empty() || stack.Pop() != '{' {
				return false
			}
		}
	}

	return stack.Empty()
}
