/**
 * Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
 * k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
 * Follow up:
 *
 * 	Could you solve the problem in O(1) extra memory space?
 * 	You may not alter the values in the list's nodes, only nodes itself may be changed.
 *
 *
 * Example 1:
 * (https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [2,1,4,3,5]
 *
 * Example 2:
 * (https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)
 * Input: head = [1,2,3,4,5], k = 3
 * Output: [3,2,1,4,5]
 *
 * Example 3:
 *
 * Input: head = [1,2,3,4,5], k = 1
 * Output: [1,2,3,4,5]
 *
 * Example 4:
 *
 * Input: head = [1], k = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the list is in the range sz.
 * 	1 <= sz <= 5000
 * 	0 <= Node.val <= 1000
 * 	1 <= k <= sz
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(begin, end *ListNode) *ListNode {
	var prev *ListNode
	node := begin.Next
	endNode := node
	for node != end {
		temp := node
		node = node.Next
		temp.Next = prev
		prev = temp
	}
	begin.Next = prev
	endNode.Next = end
	return endNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	dummyNode := &ListNode{
		Next: head,
	}

	begin := dummyNode
	i := 0
	for head != nil {
		i++
		if i%k == 0 {
			begin = reverse(begin, head.Next)
			head = begin.Next
		} else {
			head = head.Next
		}
	}

	return dummyNode.Next
}
