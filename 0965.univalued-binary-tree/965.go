/**
 * A binary tree is univalued if every node in the tree has the same value.
 *
 * Return true if and only if the given tree is univalued.
 *
 *
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png" style="width: 265px; height: 172px;" />
 *
 * Input: <span id="example-input-1-1">[1,1,1,1,1,null,1]</span>
 * Output: <span id="example-output-1">true</span>
 *
 *
 * <div>
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png" style="width: 198px; height: 169px;" />
 *
 * Input: <span id="example-input-2-1">[2,2,2,5,2]</span>
 * Output: <span id="example-output-2">false</span>
 *
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The number of nodes in the given tree will be in the range [1, 100].
 * 	Each node's value will be an integer in the range [0, 99].
 * </ol>
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func uni(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	}
	return uni(root.Left, val) && uni(root.Right, val)
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return uni(root, root.Val)
}
