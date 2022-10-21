/**
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 *    1
 *  /   \
 * 2     3
 *  \
 *   5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 *
 */

package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPaths(root *TreeNode, curPath string, res *[]string) {
	path := fmt.Sprintf("%s->%d", curPath, root.Val)

	// if both are nil
	if root.Left == root.Right {
		*res = append(*res, path[2:])
		return
	}

	if root.Left != nil {
		allPaths(root.Left, path, res)
	}

	if root.Right != nil {
		allPaths(root.Right, path, res)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root != nil {
		allPaths(root, "", &res)
	}
	return res
}
