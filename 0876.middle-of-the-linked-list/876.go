/**
 * Given a non-empty, singly linked list with head node head, return a middle node of linked list.
 *
 * If there are two middle nodes, return the second middle node.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[1,2,3,4,5]</span>
 * Output: Node 3 from this list (Serialization: <span id="example-output-1">[3,4,5]</span>)
 * The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
 * Note that we returned a ListNode object ans, such that:
 * ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[1,2,3,4,5,6]</span>
 * Output: Node 4 from this list (Serialization: <span id="example-output-2">[4,5,6]</span>)
 * Since the list has two middle nodes with values 3 and 4, we return the second one.
 *
 *
 *
 *
 * Note:
 *
 *
 * 	The number of nodes in the given list will be between 1 and 100.
 *
 * </div>
 * </div>
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
	}
	return head
}
