/**
 * Given the root of a binary tree, then value v and depth d, you need to add a row of nodes with value v at the given depth d. The root node is at depth 1.
 *
 * The adding rule is: given a positive integer depth d, for each NOT null tree nodes N in depth d-1, create two tree nodes with value v as N's left subtree root and right subtree root. And N's original left subtree should be the left subtree of the new left subtree root, its original right subtree should be the right subtree of the new right subtree root. If depth d is 1 that means there is no depth d-1 at all, then create a tree node with value v as the new root of the whole original tree, and the original tree is the new root's left subtree.
 *
 * Example 1:<br />
 *
 * Input:
 * A binary tree as following:
 *        4
 *      /   \
 *     2     6
 *    / \   /
 *   3   1 5
 *
 * v = 1
 *
 * d = 2
 *
 * Output:
 *        4
 *       / \
 *      1   1
 *     /     \
 *    2       6
 *   / \     /
 *  3   1   5
 *
 *
 *
 *
 *
 * Example 2:<br />
 *
 * Input:
 * A binary tree as following:
 *       4
 *      /
 *     2
 *    / \
 *   3   1
 *
 * v = 1
 *
 * d = 3
 *
 * Output:
 *       4
 *      /
 *     2
 *    / \
 *   1   1
 *  /     \
 * 3       1
 *
 *
 *
 * Note:<br>
 * <ol>
 * The given d is in range [1, maximum depth of the given tree + 1].
 * The given binary tree has at least one tree node.
 * </ol>
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The given d is in range [1, maximum depth of the given tree + 1]
// The idea is to:
// use 1 to indicate attach to left node, and 0 to indicate attach to right node
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 0 || d == 1 {
		newRoot := &TreeNode{
			Val: v,
		}
		if d == 0 {
			newRoot.Right = root
		} else {
			newRoot.Left = root
		}
		root = newRoot
	} else if root != nil {
		if d > 2 {
			root.Left = addOneRow(root.Left, v, d-1)
			root.Right = addOneRow(root.Right, v, d-1)
		} else {
			root.Left = addOneRow(root.Left, v, 1)
			root.Right = addOneRow(root.Right, v, 0)
		}
	}

	return root
}
