/**
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 *
 * Assume a BST is defined as follows:
 *
 *
 * 	The left subtree of a node contains only nodes with keys less than the node's key.
 * 	The right subtree of a node contains only nodes with keys greater than the node's key.
 * 	Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * Example 1:
 *
 *
 *     2
 *    / \
 *   1   3
 *
 * Input: [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 *     5
 *    / \
 *   1   4
 *      / \
 *     3   6
 *
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
 *
 *
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isvalid(root *TreeNode, prev *int) bool {
	if root == nil {
		return true
	}

	left := isvalid(root.Left, prev)
	if *prev != math.MinInt64 && root.Val <= *prev {
		return false
	}
	*prev = root.Val

	right := isvalid(root.Right, prev)
	return left && right
}

func isValidBST(root *TreeNode) bool {
	prev := math.MinInt64
	return isvalid(root, &prev)
}
