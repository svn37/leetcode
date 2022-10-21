/**
 * Given the root of a binary tree, return the length of the longest path, where each node in the path has the same value. This path may or may not pass through the root.
 * The length of the path between two nodes is represented by the number of edges between them.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/ex1.jpg" style="width: 571px; height: 302px;" />
 * Input: root = [5,4,5,1,1,5]
 * Output: 2
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/ex2.jpg" style="width: 571px; height: 302px;" />
 * Input: root = [1,4,5,4,4,5]
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 10^4].
 * 	-1000 <= Node.val <= 1000
 * 	The depth of the tree will not exceed 1000.
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

func helper(root *TreeNode, prev int, maxVal *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, root.Val, maxVal)
	right := helper(root.Right, root.Val, maxVal)

	*maxVal = max(*maxVal, left+right)
	if root.Val == prev {
		return max(left, right) + 1
	}
	return 0
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := 0
	helper(root, root.Val, &maxVal)
	return maxVal
}
