/**
 * Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.
 *
 * After doing so, return the head of the final linked list.  You may return any such answer.
 *
 * (Note that in the examples below, all sequences are serializations of ListNode objects.)
 * Example 1:
 *
 * Input: head = [1,2,-3,3,1]
 * Output: [3,1]
 * Note: The answer [1,2,1] would also be accepted.
 *
 * Example 2:
 *
 * Input: head = [1,2,3,-3,4]
 * Output: [1,2,4]
 *
 * Example 3:
 *
 * Input: head = [1,2,3,-3,-2]
 * Output: [1]
 *
 *
 * Constraints:
 *
 * 	The given linked list will contain between 1 and 1000 nodes.
 * 	Each node in the linked list has -1000 <= node.val <= 1000.
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)

	dummyNode := &ListNode{
		Next: head,
	}
	m[0] = dummyNode

	sum := 0
	for head != nil {
		sum += head.Val
		if node, ok := m[sum]; ok {
			cur := node.Next
			p := sum
			for cur != head {
				p += cur.Val
				delete(m, p)
				cur = cur.Next
			}
			node.Next = head.Next
		} else {
			m[sum] = head
		}
		head = head.Next
	}

	return dummyNode.Next
}
