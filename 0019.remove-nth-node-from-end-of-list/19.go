/**
 * Given the head of a linked list, remove the n^th node from the end of the list and return its head.
 * Follow up: Could you do this in one pass?
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg" style="width: 542px; height: 222px;" />
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 * Example 2:
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 * Example 3:
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the list is sz.
 * 	1 <= sz <= 30
 * 	0 <= Node.val <= 100
 * 	1 <= n <= sz
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func recur(head *ListNode, n int) (*ListNode, int) {
	if head == nil {
		return head, 1
	}
	node, k := recur(head.Next, n)
	if k == n+1 {
		head.Next = node.Next
	}
	return head, k + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head, k := recur(head, n)
	if k == n+1 {
		return head.Next
	}
	return head
}

// Slow, fast pointers
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n != 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
