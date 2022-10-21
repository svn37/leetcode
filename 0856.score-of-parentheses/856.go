/**
 * Given a balanced parentheses string S, compute the score of the string based on the following rule:
 *
 *
 * 	() has score 1
 * 	AB has score A + B, where A and B are balanced parentheses strings.
 * 	(A) has score 2 * A, where A is a balanced parentheses string.
 *
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"()"</span>
 * Output: <span id="example-output-1">1</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">"(())"</span>
 * Output: <span id="example-output-2">2</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">"()()"</span>
 * Output: <span id="example-output-3">2</span>
 *
 *
 * <div>
 * Example 4:
 *
 *
 * Input: <span id="example-input-4-1">"(()(()))"</span>
 * Output: <span id="example-output-4">6</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	S is a balanced parentheses string, containing only ( and ).
 * 	2 <= S.length <= 50
 * </ol>
 * </div>
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

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

func scoreOfParentheses(S string) int {
	stack := &Stack{}
	stack.Push(0)

	for i := 1; i < len(S); i++ {
		if S[i] == '(' && S[i-1] == '(' {
			stack.Push(0)
		} else if S[i] == ')' && S[i-1] == '(' {
			stack.Push(stack.Pop() + 1)
		} else if S[i] == ')' && S[i-1] == ')' {
			stack.Push(stack.Pop()*2 + stack.Pop())
		}
	}
	return stack.Pop()
}
