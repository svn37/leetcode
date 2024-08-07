/**
 * Given the root of a binary search tree and the lowest and highest boundaries as low and high, trim the tree so that all its elements lies in [low, high]. Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a unique answer.
 * Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.
 *
 * Example 1:
 * (https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg)
 * Input: root = [1,0,2], low = 1, high = 2
 * Output: [1,null,2]
 *
 * Example 2:
 * (https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg)
 * Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
 * Output: [3,2,null,1]
 *
 * Example 3:
 *
 * Input: root = [1], low = 1, high = 2
 * Output: [1]
 *
 * Example 4:
 *
 * Input: root = [1,null,2], low = 1, high = 3
 * Output: [1,null,2]
 *
 * Example 5:
 *
 * Input: root = [1,null,2], low = 2, high = 4
 * Output: [2]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree in the range [1, 10^4].
 * 	0 <= Node.val <= 10^4
 * 	The value of each node in the tree is unique.
 * 	root is guaranteed to be a valid binary search tree.
 * 	0 <= low <= high <= 10^4
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > R {
		return trimBST(root.Left, L, R)
	} else if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)

	return root
}
