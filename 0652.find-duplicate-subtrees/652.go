/**
 * Given the root of a binary tree, return all duplicate subtrees.
 * For each kind of duplicate subtrees, you only need to return the root node of any one of them.
 * Two trees are duplicate if they have the same structure with the same node values.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/16/e1.jpg" style="width: 450px; height: 354px;" />
 * Input: root = [1,2,3,4,null,2,4,null,null,4]
 * Output: [[2,4],[4]]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/16/e2.jpg" style="width: 321px; height: 201px;" />
 * Input: root = [2,1,1]
 * Output: [[1]]
 *
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/16/e33.jpg" style="width: 450px; height: 303px;" />
 * Input: root = [2,2,2,3,null,3,null]
 * Output: [[2,3],[3]]
 *
 *
 * Constraints:
 *
 * 	The number of the nodes in the tree will be in the range [1, 10^4]
 * 	-200 <= Node.val <= 200
 *
 */

package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1,2,#,#,3,4,#,#,5,#,#
// it's postorder indeed.
// Even though cur.val is at the beginning of the string, it doesn't make the algorithm preorder.
// The order of value in the string doesn't really matter, it's the order that we process subtrees that matters.
func postorder(root *TreeNode, m map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	serial := fmt.Sprintf("%d %s %s", root.Val, postorder(root.Left, m, res), postorder(root.Right, m, res))
	// map[string]struct{} doesn't fit here, because we need to return duplicate exactly once.
	if m[serial] == 1 {
		*res = append(*res, root)
	}
	m[serial]++

	return serial
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	postorder(root, make(map[string]int), &res)
	return res
}
