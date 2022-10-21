/**
 *
 * Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
 *
 *
 *
 * Example:<br />
 * Given a binary tree <br />
 *
 *           1
 *          / \
 *         2   3
 *        / \
 *       4   5
 *
 *
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 *
 * Note:
 * The length of path between two nodes is represented by the number of edges between them.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode, maxVal *int) int {
	if root == nil {
		return 0
	}
	L := depth(root.Left, maxVal)
	R := depth(root.Right, maxVal)
	*maxVal = max(*maxVal, L+R)
	return max(L, R) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxVal := 0
	depth(root, &maxVal)
	return maxVal
}
