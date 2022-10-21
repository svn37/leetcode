/**
 * Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
 * <img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png" style="width: 400px; height: 336px;" />
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
 * Two binary trees are considered leaf-similar if their leaf value sequence is the same.
 * Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" style="width: 750px; height: 297px;" />
 * Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
 * Output: true
 *
 * Example 2:
 *
 * Input: root1 = [1], root2 = [1]
 * Output: true
 *
 * Example 3:
 *
 * Input: root1 = [1], root2 = [2]
 * Output: false
 *
 * Example 4:
 *
 * Input: root1 = [1,2], root2 = [2,2]
 * Output: true
 *
 * Example 5:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" style="width: 450px; height: 165px;" />
 * Input: root1 = [1,2,3], root2 = [1,3,2]
 * Output: false
 *
 *
 * Constraints:
 *
 * 	The number of nodes in each tree will be in the range [1, 200].
 * 	Both of the given trees will have values in the range [0, 200].
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode, fn func(int) bool) bool {
	if root.Left == nil && root.Right == nil {
		return fn(root.Val)
	}

	if root.Left != nil && !findLeaves(root.Left, fn) {
		return false
	}
	if root.Right != nil && !findLeaves(root.Right, fn) {
		return false
	}
	return true
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves := []int{}
	findLeaves(root1, func(val int) bool {
		leaves = append(leaves, val)
		return true
	})

	return findLeaves(root2, func(val int) bool {
		if len(leaves) == 0 || leaves[0] != val {
			return false
		}
		leaves = leaves[1:]
		return true
	})
}
