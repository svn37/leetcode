/**
 * Find the sum of all left leaves in a given binary tree.
 *
 * Example:
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 * There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumleaves(root *TreeNode, sum *int, left bool) {
	if root == nil {
		return
	}

	if left && (root.Left == nil && root.Right == nil) {
		*sum += root.Val
		return
	}

	sumleaves(root.Left, sum, true)
	sumleaves(root.Right, sum, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	sumleaves(root, &sum, false)
	return sum
}
