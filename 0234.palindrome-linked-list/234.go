/**
 * Given a singly linked list, determine if it is a palindrome.
 *
 * Example 1:
 *
 *
 * Input: 1->2
 * Output: false
 *
 * Example 2:
 *
 *
 * Input: 1->2->2->1
 * Output: true
 *
 * Follow up:<br />
 * Could you do it in O(n) time and O(1) space?
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// This method ends up modifying the original list.
// -- find the middle of the list by using slow and fast pointers
// -- reverse the second half of the list
// -- compare if both lists contain the same values
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := slow.Next
	var prevNode *ListNode
	// divide the list into two parts
	slow.Next = nil

	for node != nil {
		nextNode := node.Next
		// reverse link
		node.Next = prevNode

		prevNode = node
		node = nextNode
	}

	// head and prevNode are the start of two linked lists
	// which must be contains the same values (except for the last character)
	for head != nil && prevNode != nil {
		if head.Val != prevNode.Val {
			return false
		}
		head = head.Next
		prevNode = prevNode.Next
	}
	return true
}
