/**
 * Given a complete binary tree, count the number of nodes.
 *
 * Note:
 *
 * <u>Definition of a complete binary tree from <a href="http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees" target="_blank">Wikipedia</a>:</u><br />
 * In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2^h nodes inclusive at the last level h.
 *
 * Example:
 *
 *
 * Input:
 *     1
 *    / \
 *   2   3
 *  / \  /
 * 4  5 6
 *
 * Output: 6
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var height uint32
	left, right := root, root
	for right != nil {
		left = left.Left
		right = right.Right
		height++
	}
	// all levels are fully filled
	if left == nil {
		return (1 << height) - 1
	}
	// this is basically the whole solution
	// however, each of these calls might return the correct answer
	// without recursing further, which is a gigantic optimization
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
