/**
 * Given a binary tree, flatten it to a linked list in-place.
 *
 * For example, given the following tree:
 *
 *
 *     1
 *    / \
 *   2   5
 *  / \   \
 * 3   4   6
 *
 *
 * The flattened tree should look like:
 *
 *
 * 1
 *  \
 *   2
 *    \
 *     3
 *      \
 *       4
 *        \
 *         5
 *          \
 *           6
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorder(root *TreeNode, prevNode **TreeNode) {
	if root == nil {
		return
	}

	postorder(root.Right, prevNode)
	postorder(root.Left, prevNode)

	if *prevNode != nil {
		root.Left = nil
		root.Right = *prevNode
	}
	*prevNode = root
}

func flatten(root *TreeNode) {
	var prevNode *TreeNode
	postorder(root, &prevNode)
}
