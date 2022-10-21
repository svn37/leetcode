/**
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 &le; m &le; n &le; length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 *
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// four pointers
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	head = &ListNode{
		Next: head,
	} // dummy node

	left := head
	var start *ListNode
	var end *ListNode

	i := 1
	node := head.Next

	for node != nil {
		nextNode := node.Next

		if i == m-1 {
			left = node
		} else if i == m {
			end = node
		}

		if i >= m {
			node.Next = start
			start = node
		}

		if i == n {
			left.Next = start
			end.Next = nextNode
			break
		}

		node = nextNode
		i++
	}
	return head.Next
}
