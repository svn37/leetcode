/**
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{
		Next: head,
	} // dummy node

	prev := head
	node := head.Next

	for node != nil {
		if node.Val != val {
			prev.Next = node
			prev = node
		}
		node = node.Next
	}
	prev.Next = node

	return head.Next
}
