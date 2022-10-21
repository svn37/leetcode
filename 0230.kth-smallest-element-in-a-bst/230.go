/**
 * Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
 *
 * Example 1:
 *
 * Input: root = [3,1,4,null,2], k = 1
 *    3
 *   / \
 *  1   4
 *   \
 *    2
 * Output: 1
 * Example 2:
 *
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 *        5
 *       / \
 *      3   6
 *     / \
 *    2   4
 *   /
 *  1
 * Output: 3
 *
 * Follow up:<br />
 * What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
 *
 * Constraints:
 *
 * 	The number of elements of the BST is between 1 to 10^4.
 * 	You may assume k is always valid, 1 &le; k &le; BST's total elements.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kth(root *TreeNode, count *int, k int) int {
	if root == nil {
		return 0
	}

	left := kth(root.Left, count, k)
	if left != 0 {
		return left
	}

	*count++

	if *count == k {
		return root.Val
	}

	return kth(root.Right, count, k)
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	return kth(root, &count, k)
}
