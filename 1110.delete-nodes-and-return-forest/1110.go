/**
 * Given the root of a binary tree, each node in the tree has a distinct value.
 *
 * After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).
 *
 * Return the roots of the trees in the remaining forest.  You may return the result in any order.
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png" style="width: 237px; height: 150px;" />
 *
 *
 * Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
 * Output: [[1,2,null,4],[6],[7]]
 *
 *
 *
 * Constraints:
 *
 *
 * 	The number of nodes in the given tree is at most 1000.
 * 	Each node has a distinct value between 1 and 1000.
 * 	to_delete.length <= 1000
 * 	to_delete contains distinct values between 1 and 1000.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, isRoot bool, values map[int]struct{}, roots *[]*TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	_, deleted := values[node.Val]
	if !deleted && isRoot {
		*roots = append(*roots, node)
	}
	node.Left = helper(node.Left, deleted, values, roots)
	node.Right = helper(node.Right, deleted, values, roots)
	if deleted {
		return nil
	}
	return node
}

func delNodes(root *TreeNode, values []int) []*TreeNode {
	set := make(map[int]struct{})
	for _, value := range values {
		set[value] = struct{}{}
	}

	roots := make([]*TreeNode, 0)
	helper(root, true, set, &roots)
	return roots
}
