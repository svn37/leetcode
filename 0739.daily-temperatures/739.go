/**
 *
 * Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature.  If there is no future day for which this is possible, put 0 instead.
 *
 * For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
 *
 *
 * Note:
 * The length of temperatures will be in the range [1, 30000].
 * Each temperature will be an integer in the range [30, 100].
 *
 */

package leetcode

type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Peek() int {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Pop() int {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))

	stack := &Stack{}

	for i := range T {
		for !stack.Empty() && T[i] > T[stack.Peek()] {
			idx := stack.Pop()
			res[idx] = i - idx
		}
		stack.Push(i)
	}
	return res
}
