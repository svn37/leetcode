/**
 * Given a linked list, swap every two adjacent nodes and return its head.
 * You may not modify the values in the list's nodes. Only nodes itself may be changed.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/swap_ex1.jpg" style="width: 422px; height: 222px;" />
 * Input: head = [1,2,3,4]
 * Output: [2,1,4,3]
 *
 * Example 2:
 *
 * Input: head = []
 * Output: []
 *
 * Example 3:
 *
 * Input: head = [1]
 * Output: [1]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the list is in the range [0, 100].
 * 	0 <= Node.val <= 100
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swappairs(head, prev *ListNode, count int) *ListNode {
	if head == nil {
		return nil
	}

	prevNext := head.Next
	if count%2 == 0 {
		head.Next = prev
		return swappairs(prevNext, nil, count+1)
	} else {
		if prevNext == nil {
			return head
		}
		head.Next = swappairs(head.Next, head, count+1)
		return prevNext
	}
}

func swapPairs(head *ListNode) *ListNode {
	return swappairs(head, nil, 1)
}
