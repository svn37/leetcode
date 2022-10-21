/**
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 *
 * Note:<br />
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 *
 * Return the following binary tree:
 *
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(inorder []int, is, ie int, postorder []int, ps, pe int, m map[int]int) *TreeNode {
	if ie < is || pe < ps {
		return nil
	}

	root := &TreeNode{
		Val: postorder[pe],
	}

	ri := m[postorder[pe]]

	root.Left = build(inorder, is, ri-1, postorder, ps, ps+ri-is-1, m)
	root.Right = build(inorder, ri+1, ie, postorder, ps+ri-is, pe-1, m)

	return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i := range inorder {
		m[inorder[i]] = i
	}

	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, m)
}
