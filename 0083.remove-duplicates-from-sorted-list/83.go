/**
 * Given a sorted linked list, delete all duplicates such that each element appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
 *
 *
 */

package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	head = &ListNode{math.MinInt64, head}

	prev := head
	node := head.Next
	for node != nil {
		if node.Val != prev.Val {
			prev.Next = node
			prev = node
		}
		node = node.Next
	}
	prev.Next = node

	return head.Next
}
