/**
 * Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
 * Return the linked list sorted as well.
 * Example 1:
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 * Example 2:
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
 *
 */

package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	head = &ListNode{math.MinInt64, head} // dummy node

	prev := head
	node := head.Next
	for node != nil {
		if node.Next != nil && node.Next.Val == node.Val {
			for node.Next != nil && node.Next.Val == node.Val {
				node = node.Next
			}
			prev.Next = node.Next
		} else {
			prev = node
		}
		node = node.Next
	}
	return head.Next
}
