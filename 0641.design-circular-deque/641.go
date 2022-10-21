/**
 * Design your implementation of the circular double-ended queue (deque).
 *
 * Your implementation should support following operations:
 *
 *
 * 	MyCircularDeque(k): Constructor, set the size of the deque to be k.
 * 	insertFront(): Adds an item at the front of Deque. Return true if the operation is successful.
 * 	insertLast(): Adds an item at the rear of Deque. Return true if the operation is successful.
 * 	deleteFront(): Deletes an item from the front of Deque. Return true if the operation is successful.
 * 	deleteLast(): Deletes an item from the rear of Deque. Return true if the operation is successful.
 * 	getFront(): Gets the front item from the Deque. If the deque is empty, return -1.
 * 	getRear(): Gets the last item from Deque. If the deque is empty, return -1.
 * 	isEmpty(): Checks whether Deque is empty or not.
 * 	isFull(): Checks whether Deque is full or not.
 *
 *
 *
 *
 * Example:
 *
 *
 * MyCircularDeque circularDeque = new MycircularDeque(3); // set the size to be 3
 * circularDeque.insertLast(1);			// return true
 * circularDeque.insertLast(2);			// return true
 * circularDeque.insertFront(3);			// return true
 * circularDeque.insertFront(4);			// return false, the queue is full
 * circularDeque.getRear();  			// return 2
 * circularDeque.isFull();				// return true
 * circularDeque.deleteLast();			// return true
 * circularDeque.insertFront(4);			// return true
 * circularDeque.getFront();			// return 4
 *
 *
 *
 *
 * Note:
 *
 *
 * 	All values will be in the range of [0, 1000].
 * 	The number of operations will be in the range of [1, 1000].
 * 	Please do not use the built-in Deque library.
 *
 *
 */

package leetcode

type MyCircularDeque struct {
	items       []int
	front, rear int
	count, k    int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		items: make([]int, k),
		front: k - 1,
		k:     k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.items[this.front] = value
	this.front--

	// wraparound
	this.front += this.k
	this.front %= this.k

	this.count++

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.items[this.rear] = value
	this.rear++

	// wraparound
	this.rear %= this.k

	this.count++

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.front++

	// wraparound
	this.front %= this.k

	this.count--

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.rear--

	// wraparound
	this.rear += this.k
	this.rear %= this.k

	this.count--

	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[(this.front+1)%this.k]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[(this.rear-1+this.k)%this.k]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.k
}
