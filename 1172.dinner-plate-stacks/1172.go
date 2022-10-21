/**
 * You have an infinite number of stacks arranged in a row and numbered (left to right) from 0, each of the stacks has the same maximum capacity.
 * Implement the DinnerPlates class:
 *
 * 	DinnerPlates(int capacity) Initializes the object with the maximum capacity of the stacks.
 * 	void push(int val) Pushes the given positive integer val into the leftmost stack with size less than capacity.
 * 	int pop() Returns the value at the top of the rightmost non-empty stack and removes it from that stack, and returns -1 if all stacks are empty.
 * 	int popAtStack(int index) Returns the value at the top of the stack with the given index and removes it from that stack, and returns -1 if the stack with that given index is empty.
 *
 * Example:
 *
 * Input:
 * ["DinnerPlates","push","push","push","push","push","popAtStack","push","push","popAtStack","popAtStack","pop","pop","pop","pop","pop"]
 * [[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]
 * Output:
 * [null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]
 * Explanation:
 * DinnerPlates D = DinnerPlates(2);  // Initialize with capacity = 2
 * D.push(1);
 * D.push(2);
 * D.push(3);
 * D.push(4);
 * D.push(5);         // The stacks are now:  2  4
 *                                            1  3  5
 *                                            ﹈ ﹈ ﹈
 * D.popAtStack(0);   // Returns 2.  The stacks are now:     4
 *                                                        1  3  5
 *                                                        ﹈ ﹈ ﹈
 * D.push(20);        // The stacks are now: 20  4
 *                                            1  3  5
 *                                            ﹈ ﹈ ﹈
 * D.push(21);        // The stacks are now: 20  4 21
 *                                            1  3  5
 *                                            ﹈ ﹈ ﹈
 * D.popAtStack(0);   // Returns 20.  The stacks are now:     4 21
 *                                                         1  3  5
 *                                                         ﹈ ﹈ ﹈
 * D.popAtStack(2);   // Returns 21.  The stacks are now:     4
 *                                                         1  3  5
 *                                                         ﹈ ﹈ ﹈
 * D.pop()            // Returns 5.  The stacks are now:      4
 *                                                         1  3
 *                                                         ﹈ ﹈
 * D.pop()            // Returns 4.  The stacks are now:   1  3
 *                                                         ﹈ ﹈
 * D.pop()            // Returns 3.  The stacks are now:   1
 *                                                         ﹈
 * D.pop()            // Returns 1.  There are no stacks.
 * D.pop()            // Returns -1.  There are still no stacks.
 *
 *
 * Constraints:
 *
 * 	1 <= capacity <= 20000
 * 	1 <= val <= 20000
 * 	0 <= index <= 100000
 * 	At most 200000 calls will be made to push, pop, and popAtStack.
 *
 */

package leetcode

type Heap struct {
	storage []*Stack
}

func NewHeap(a []*Stack) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) Min() *Stack {
	return h.storage[0]
}

func (h *Heap) ExtractMin() *Stack {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *Stack) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].num < h.storage[parent].num {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) Empty() bool {
	return len(h.storage) == 0
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 != 0 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func (h *Heap) leftchild(i int) int {
	return 2*i + 1
}

func (h *Heap) rightchild(i int) int {
	return 2*i + 2
}

func (h *Heap) heapify(i int) {
	l := h.leftchild(i)
	r := h.rightchild(i)
	smallest := i

	if l < len(h.storage) && h.storage[l].num < h.storage[smallest].num {
		smallest = l
	}

	if r < len(h.storage) && h.storage[r].num < h.storage[smallest].num {
		smallest = r
	}

	if smallest != i {
		h.storage[smallest], h.storage[i] = h.storage[i], h.storage[smallest]
		h.heapify(smallest)
	}
}

type Stack struct {
	storage []int
	num     int // number in the list for extracting min from the heap
}

func NewStack(num, capacity int) *Stack {
	return &Stack{
		storage: make([]int, 0, capacity),
		num:     num,
	}
}

func (s *Stack) Push(val int) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Pop() int {
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Size() int {
	return len(s.storage)
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

type DinnerPlates struct {
	list     []*Stack
	heap     *Heap
	capacity int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		list:     make([]*Stack, 0),
		heap:     &Heap{},
		capacity: capacity,
	}
}

func (this *DinnerPlates) pushStack() *Stack {
	stack := NewStack(len(this.list), this.capacity)
	this.list = append(this.list, stack)
	return stack
}

func (this *DinnerPlates) popStack() {
	this.list = this.list[:len(this.list)-1]
}

func (this *DinnerPlates) rightmostStack() *Stack {
	return this.list[len(this.list)-1]
}

func (this *DinnerPlates) getStack(index int) *Stack {
	return this.list[index]
}

func (this *DinnerPlates) removeEmptyStacks() {
	for len(this.list) > 0 && this.rightmostStack().Size() == 0 {
		this.popStack()
	}
}

func (this *DinnerPlates) Push(val int) {
	if this.capacity == 0 {
		return
	}

	// after pop some empty stacks at the end were deleted.
	// they remain in the heap, need to get rid of them
	for !this.heap.Empty() && this.heap.Min().num >= len(this.list) {
		this.heap.ExtractMin()
	}

	var leftmostStack *Stack
	if this.heap.Empty() {
		leftmostStack = this.pushStack()
		if leftmostStack.Size() < this.capacity-1 {
			this.heap.Push(leftmostStack)
		}
	} else {
		leftmostStack = this.heap.Min()
		if leftmostStack.Size() == this.capacity-1 {
			this.heap.ExtractMin()
		}
	}
	leftmostStack.Push(val)
}

func (this *DinnerPlates) Pop() int {
	if len(this.list) == 0 {
		return -1
	}

	// after popAtStack any stacks might be empty.
	// need to clean up space
	this.removeEmptyStacks()

	defer func() {
		this.removeEmptyStacks()
	}()

	rightmostStack := this.rightmostStack()

	if rightmostStack.Size() == this.capacity {
		this.heap.Push(rightmostStack)
	}

	return rightmostStack.Pop()
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(this.list) {
		return -1
	}

	stack := this.getStack(index)
	if stack.Size() == this.capacity {
		this.heap.Push(stack)
	}

	if stack.Size() == 0 {
		return -1
	}

	return stack.Pop()
}
