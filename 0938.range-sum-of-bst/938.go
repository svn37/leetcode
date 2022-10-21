/**
 * Given the root node of a binary search tree, return the sum of values of all nodes with a value in the range [low, high].
 *
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg" style="width: 400px; height: 222px;" />
 *
 * Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
 * Output: 32
 *
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg" style="width: 400px; height: 335px;" />
 *
 * Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
 * Output: 23
 *
 *
 *
 * Constraints:
 *
 *
 * 	The number of nodes in the tree is in the range [1, 2 * 10^4].
 * 	1 <= Node.val <= 10^5
 * 	1 <= low <= high <= 10^5
 * 	All Node.val are unique.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sum(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val > R {
		return sum(root.Left, L, R)
	}
	if root.Val < L {
		return sum(root.Right, L, R)
	}
	return root.Val + sum(root.Left, L, R) + sum(root.Right, L, R)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	return sum(root, L, R)
}
