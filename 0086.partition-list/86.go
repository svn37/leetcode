/**
 * Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
 *
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	head1 := &ListNode{}
	head2 := &ListNode{}

	lesser, greater := head1, head2

	for head != nil {
		if head.Val < x {
			lesser.Next = head
			lesser = head
		} else {
			greater.Next = head
			greater = head
		}
		head = head.Next
	}
	lesser.Next, greater.Next = head2.Next, nil
	return head1.Next
}
