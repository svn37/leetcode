/**
 * Given the head of a linked list, return the list after sorting it in ascending order.
 * Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/14/sort_list_1.jpg" style="width: 450px; height: 194px;" />
 * Input: head = [4,2,1,3]
 * Output: [1,2,3,4]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/14/sort_list_2.jpg" style="width: 550px; height: 184px;" />
 * Input: head = [-1,5,3,4,0]
 * Output: [-1,0,3,4,5]
 *
 * Example 3:
 *
 * Input: head = []
 * Output: []
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the list is in the range [0, 5 * 10^4].
 * 	-10^5 <= Node.val <= 10^5
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Method 1. Recursive mergesort (implicit stack space, thus not O(1) space complexity)
func merge(left, right *ListNode) *ListNode {
	head := &ListNode{} // dummy value
	node := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			node = left
			left = left.Next
		} else {
			node.Next = right
			node = right
			right = right.Next
		}
	}

	if left != nil {
		node.Next = left
	} else {
		node.Next = right
	}
	return head.Next
}

func mergeSort(left *ListNode, length int) *ListNode {
	if length < 2 {
		return left
	}

	mid := length / 2

	// find the start of the right half
	node := left
	for i := 1; i < mid; i++ {
		node = node.Next
	}

	// divide the two parts
	right := node.Next
	node.Next = nil

	return merge(mergeSort(left, mid), mergeSort(right, length-mid))
}

func sortList(head *ListNode) *ListNode {
	// calculate the length of the linked list in advance
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}
	return mergeSort(head, l)
}

// Method 2. Bottom-up merge sort (no recursion, O(1) space complexity).
// standard merge procedure
func merge2(left, right, head *ListNode) *ListNode {
	node := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			node = left
			left = left.Next
		} else {
			node.Next = right
			node = right
			right = right.Next
		}
	}

	if left != nil {
		node.Next = left
	} else {
		node.Next = right
	}

	for node.Next != nil {
		node = node.Next
	}
	return node
}

// helper function to cut the start of the list of length n
func split(head *ListNode, n int) *ListNode {
	for i := 1; head != nil && i < n; i++ {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	// the head of the list after the cut
	right := head.Next
	head.Next = nil
	return right
}

// bottom-up mergesort first merges pairs of adjacent arrays of 1 elements
// then merges pairs of adjacent arrays of 2 elements
// then merges pairs of adjacent arrays of 4 elements
// and so on, until the whole array is merged
func sortList2(head *ListNode) *ListNode {
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}

	head = &ListNode{
		Next: head,
	} // dummy node

	for n := 1; n < l; n *= 2 {
		cur, tail := head.Next, head
		for cur != nil {
			// two halves of length n
			left := cur
			right := split(left, n)
			// update cur value for the next iteration
			cur = split(right, n)
			// merge left and right while using tail (basically, left)
			// as the working area. update working area start for the next iteration
			tail = merge2(left, right, tail)
		}
	}
	return head.Next
}
