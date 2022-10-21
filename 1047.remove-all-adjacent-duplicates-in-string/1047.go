/**
 * Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
 *
 * We repeatedly make duplicate removals on S until we no longer can.
 *
 * Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">"abbaca"</span>
 * Output: <span id="example-output-1">"ca"</span>
 * Explanation:
 * For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= S.length <= 20000
 * 	S consists only of English lowercase letters.
 * </ol>
 */

package leetcode

type Stack struct {
	storage []byte
}

func (s *Stack) Push(i byte) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() byte {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Top() byte {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func (s *Stack) String() string {
	return string(s.storage)
}

func removeDuplicates(S string) string {
	stack := &Stack{}
	for i := range S {
		if i != 0 && !stack.Empty() && S[i] == stack.Top() {
			stack.Pop()
		} else {
			stack.Push(S[i])
		}
	}
	return stack.String()
}
