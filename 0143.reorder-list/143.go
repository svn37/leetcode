/**
 * Given a singly linked list L: L0&rarr;L1&rarr;&hellip;&rarr;Ln-1&rarr;Ln,<br />
 * reorder it to: L0&rarr;Ln&rarr;L1&rarr;Ln-1&rarr;L2&rarr;Ln-2&rarr;&hellip;
 *
 * You may not modify the values in the list's nodes, only nodes itself may be changed.
 *
 * Example 1:
 *
 *
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 *
 * Example 2:
 *
 *
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
 *
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func reorder(node, head *ListNode) *ListNode {
	if node == nil {
		return head
	}
	// recurse until the end of the list
	// prevNode is the new parent of node.Next
	// in the reordered list. If the parent was
	// already set, returns nil.
	prevNode := reorder(node.Next, head)
	if prevNode == nil {
		return nil
	}
	if prevNode == node {
		prevNode.Next = nil
		return nil
	}
	if prevNode.Next == node {
		node.Next = nil
		return nil
	}
	nextNode := prevNode.Next
	prevNode.Next = node
	node.Next = nextNode
	return nextNode
}

func reorderList(head *ListNode) {
	reorder(head, head)
}

// Iterative
func reorderList2(head *ListNode) {
	if head == nil {
		return
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

	// merge the two halves
	node1 := head     // start of the first list
	node2 := prevNode // start of the second list
	for node1 != nil && node2 != nil {
		nextNode1 := node1.Next
		nextNode2 := node2.Next

		node1.Next = node2
		node2.Next = nextNode1

		node1 = nextNode1
		node2 = nextNode2
	}
}
