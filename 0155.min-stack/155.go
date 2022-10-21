/**
 * Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
 *
 * 	push(x) -- Push element x onto stack.
 * 	pop() -- Removes the element on top of the stack.
 * 	top() -- Get the top element.
 * 	getMin() -- Retrieve the minimum element in the stack.
 *
 *
 * Example 1:
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 * Output
 * [null,null,null,null,-3,null,0,-2]
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 * Constraints:
 *
 * 	Methods pop, top and getMin operations will always be called on non-empty stacks.
 *
 */

package leetcode

import "math"

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

func (s *Stack) Top() int {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

// Core Idea:
// 1.Minimum value is always followed by the second minimum value
// (duplicate value of the second minimum value, to ensure that when
// pop function removes the 2nd min, it does not disrupt the stack order).
// 2.And while popping you pop min and 2nd min so that, you get the correct
// min value for the remaining stack and the remaining stack top also
// points to the right place.
type MinStack struct {
	min   int
	stack *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   math.MaxInt64,
		stack: &Stack{},
	}
}

func (this *MinStack) Push(x int) {
	// only push the old minimum value when the current
	// minimum value changes after pushing the new value x
	if x <= this.min {
		this.stack.Push(this.min)
		this.min = x
	}
	this.stack.Push(x)
}

func (this *MinStack) Pop() {
	// if pop operation could result in the changing of the current minimum value,
	// pop twice and change the current minimum value to the last minimum value.
	if this.stack.Pop() == this.min {
		this.min = this.stack.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.stack.Top()
}

func (this *MinStack) GetMin() int {
	return this.min
}
