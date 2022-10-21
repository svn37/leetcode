/**
 * Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
 *
 * Follow up:<br />
 * What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?
 *
 *
 * Example:
 *
 * // Init a singly linked list [1,2,3].
 * ListNode head = new ListNode(1);
 * head.next = new ListNode(2);
 * head.next.next = new ListNode(3);
 * Solution solution = new Solution(head);
 *
 * // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
 * solution.getRandom();
 *
 *
 */

package leetcode

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-random-node/discuss/85690/using-%22Reservoir-sampling%22-O(1)-space-O(n)-time-complexityuff0cc%2B%2B
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())

	return Solution{head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	node := this.head
	val := node.Val
	for i := 1; node.Next != nil; i++ {
		node = node.Next
		// important: inclusive range
		if rand.Intn(i+1) == i {
			val = node.Val
		}
	}

	return val
}
