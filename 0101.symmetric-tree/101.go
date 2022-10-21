/**
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 *
 *     1
 *    / \
 *   2   2
 *  / \ / \
 * 3  4 4  3
 *
 *
 * But the following [1,2,2,null,3,null,3] is not:
 *
 *     1
 *    / \
 *   2   2
 *    \   \
 *    3    3
 *
 *
 * Follow up: Solve it both recursively and iteratively.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func issymmetric(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if (n1 != nil && n2 == nil) || (n1 == nil && n2 != nil) || (n1.Val != n2.Val) {
		return false
	}

	return issymmetric(n1.Left, n2.Right) && issymmetric(n1.Right, n2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || issymmetric(root.Left, root.Right)
}
