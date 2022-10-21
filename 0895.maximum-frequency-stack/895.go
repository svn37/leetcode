/**
 * Implement FreqStack, a class which simulates the operation of a stack-like data structure.
 *
 * FreqStack has two functions:
 *
 *
 * 	push(int x), which pushes an integer x onto the stack.
 * 	pop(), which removes and returns the most frequent element in the stack.
 *
 * 		If there is a tie for most frequent element, the element closest to the top of the stack is removed and returned.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * <span id="example-input-1-1">["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"]</span>,
 * <span id="example-input-1-2">[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]</span>
 * Output: <span id="example-output-1">[null,null,null,null,null,null,null,5,7,5,4]</span>
 * Explanation:
 * After making six .push operations, the stack is [5,7,5,7,4,5] from bottom to top.  Then:
 *
 * pop() -> returns 5, as 5 is the most frequent.
 * The stack becomes [5,7,5,7,4].
 *
 * pop() -> returns 7, as 5 and 7 is the most frequent, but 7 is closest to the top.
 * The stack becomes [5,7,5,4].
 *
 * pop() -> returns 5.
 * The stack becomes [5,7,4].
 *
 * pop() -> returns 4.
 * The stack becomes [5,7].
 *
 *
 *
 *
 * Note:
 *
 *
 * 	Calls to FreqStack.push(int x) will be such that 0 <= x <= 10^9.
 * 	It is guaranteed that FreqStack.pop() won't be called if the stack has zero elements.
 * 	The total number of FreqStack.push calls will not exceed 10000 in a single test case.
 * 	The total number of FreqStack.pop calls will not exceed 10000 in a single test case.
 * 	The total number of FreqStack.push and FreqStack.pop calls will not exceed 150000 across all test cases.
 *
 *
 * <div>
 *
 * </div>
 *
 */

package leetcode

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() int {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type FreqStack struct {
	// map from integer to frequency
	freq map[int]int
	// map from frequency to stack
	m map[int]*Stack
	// maxFreq for popping an element
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq: make(map[int]int),
		m:    make(map[int]*Stack),
	}
}

func (this *FreqStack) Push(x int) {
	this.freq[x]++
	freq := this.freq[x]
	this.maxFreq = max(this.maxFreq, freq)

	stack, ok := this.m[freq]
	if !ok {
		stack = &Stack{}
		this.m[freq] = stack
	}
	stack.Push(x)
}

func (this *FreqStack) Pop() int {
	stack := this.m[this.maxFreq]
	pop := stack.Pop()

	this.freq[pop]--
	if stack.Empty() {
		this.maxFreq--
	}

	return pop
}
