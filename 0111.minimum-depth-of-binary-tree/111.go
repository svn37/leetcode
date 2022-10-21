/**
 * Given a binary tree, find its minimum depth.
 * The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
 * Note: A leaf is a node with no children.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/12/ex_depth.jpg" style="width: 432px; height: 302px;" />
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 2
 *
 * Example 2:
 *
 * Input: root = [2,null,3,null,4,null,5,null,6]
 * Output: 5
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 10^5].
 * 	-1000 <= Node.val <= 1000
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// the same as usual min, also checks if l and r == 0
func findMin(l, r int) int {
	if l == 0 {
		return r
	} else if r == 0 {
		return l
	} else if l < r {
		return l
	} else {
		return r
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + findMin(minDepth(root.Left), minDepth(root.Right))
}
