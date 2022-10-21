/**
 * Reverse a singly linked list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 *
 *
 * Follow up:
 *
 * A linked list can be reversed either iteratively or recursively. Could you implement both?
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func reverse(node, prevNode *ListNode) *ListNode {
	if node.Next == nil {
		node.Next = prevNode
		// returns the new head
		return node
	}
	head := reverse(node.Next, node)
	node.Next = prevNode

	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return reverse(head, nil)
}

// Iterative
func reverseList2(head *ListNode) *ListNode {
	var prevHead *ListNode

	for head != nil {
		nextHead := head.Next
		// reverse the node
		head.Next = prevHead
		// adjust the pointers
		prevHead, head = head, nextHead
	}
	return prevHead
}
