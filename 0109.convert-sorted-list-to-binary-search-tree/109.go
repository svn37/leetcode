/**
 * Given the head of a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
 * For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/17/linked.jpg" style="width: 600px; height: 466px;" />
 * Input: head = [-10,-3,0,5,9]
 * Output: [0,-3,9,-10,null,5]
 * Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
 *
 * Example 2:
 *
 * Input: head = []
 * Output: []
 *
 * Example 3:
 *
 * Input: head = [0]
 * Output: [0]
 *
 * Example 4:
 *
 * Input: head = [1,3]
 * Output: [3,1]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in head is in the range [0, 2 * 10^4].
 * 	-10^5 <= Node.val <= 10^5
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// O(n) time with global variable
func helper(start, end int, listNode **ListNode) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2

	left := helper(start, mid-1, listNode)
	treeNode := &TreeNode{
		Val: (*listNode).Val,
	}
	*listNode = (*listNode).Next
	right := helper(mid+1, end, listNode)

	treeNode.Left = left
	treeNode.Right = right

	return treeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	size := 0
	for node := head; node != nil; node = node.Next {
		size++
	}
	return helper(0, size-1, &head)
}

// original solution, finding the middle node each recursive call
func addNode(head, head2 *ListNode) *TreeNode {
	if head == head2 {
		return nil
	}

	slow, fast := head, head
	for fast != head2 && fast.Next != head2 {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  addNode(head, slow),
		Right: addNode(slow.Next, head2),
	}
}

func sortedListToBST2(head *ListNode) *TreeNode {
	return addNode(head, nil)
}
