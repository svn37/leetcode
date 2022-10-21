/**
 * Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given the below binary tree and sum = 22,
 *
 *
 *       5
 *      / \
 *     4   8
 *    /   / \
 *   11  13  4
 *  /  \    / \
 * 7    2  5   1
 *
 *
 * Return:
 *
 *
 * [
 *    [5,4,11,2],
 *    [5,8,4,5]
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

func backtrack(root *TreeNode, path []int, sum int, res *[][]int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	sum -= root.Val

	if sum == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, append(path[:0:0], path...))
	} else {
		backtrack(root.Left, path, sum, res)
		backtrack(root.Right, path, sum, res)
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	backtrack(root, []int{}, sum, &res)

	return res
}
