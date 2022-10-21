/**
 * Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.
 * Example 1:<br />
 * Given tree s:
 *
 *      3
 *     / \
 *    4   5
 *   / \
 *  1   2
 * Given tree t:
 *
 *    4
 *   / \
 *  1   2
 * Return true, because t has the same structure and node values with a subtree of s.
 *
 * Example 2:<br />
 * Given tree s:
 *
 *      3
 *     / \
 *    4   5
 *   / \
 *  1   2
 *     /
 *    0
 * Given tree t:
 *
 *    4
 *   / \
 *  1   2
 * Return false.
 *
 *
 */

package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSame(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// can be solved the same way "652 find duplicate subtrees" is solved
func postorder(root *TreeNode, m map[string]struct{}, addSubtrees bool) string {
	if root == nil {
		return "#"
	}

	serial := fmt.Sprintf("%d %s %s", root.Val, postorder(root.Left, m, addSubtrees), postorder(root.Right, m, addSubtrees))
	if addSubtrees {
		m[serial] = struct{}{}
	}
	return serial
}

func isSubtree2(s *TreeNode, t *TreeNode) bool {
	m := make(map[string]struct{})
	postorder(s, m, true)
	_, ok := m[postorder(t, m, false)]
	return ok
}
