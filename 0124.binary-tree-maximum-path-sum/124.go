/**
 * Given a non-empty binary tree, find the maximum path sum.
 * For this problem, a path is defined as any node sequence from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg" style="width: 322px; height: 182px;" />
 * Input: root = [1,2,3]
 * Output: 6
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg" />
 * Input: root = [-10,9,20,null,null,15,7]
 * Output: 42
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 3 * 10^4].
 * 	-1000 <= Node.val <= 1000
 *
 */

package leetcode

import "math"

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

func maxPath(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	leftSum := max(0, maxPath(root.Left, maxSum))
	rightSum := max(0, maxPath(root.Right, maxSum))

	*maxSum = max(*maxSum, root.Val+leftSum+rightSum)

	return root.Val + max(leftSum, rightSum)
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	maxSum = max(maxSum, maxPath(root, &maxSum))
	return maxSum
}
