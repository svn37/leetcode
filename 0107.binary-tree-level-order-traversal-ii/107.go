/**
 * Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:<br />
 * Given binary tree [3,9,20,null,null,15,7],<br />
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 *
 *
 * return its bottom-up level order traversal as:<br />
 *
 * [
 *   [15,7],
 *   [9,20],
 *   [3]
 * ]
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recurse(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, []int{})
	}

	recurse(node.Left, level+1, res)
	(*res)[level] = append((*res)[level], node.Val)
	recurse(node.Right, level+1, res)
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	recurse(root, 0, &res)

	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
