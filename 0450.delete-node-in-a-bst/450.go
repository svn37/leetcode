/**
 * Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
 * Basically, the deletion can be divided into two stages:
 * <ol>
 * 	Search for a node to remove.
 * 	If the node is found, delete the node.
 * </ol>
 * Follow up: Can you solve it with time complexity O(height of tree)?
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" style="width: 800px; height: 214px;" />
 * Input: root = [5,3,6,2,4,null,7], key = 3
 * Output: [5,4,6,2,null,null,7]
 * Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
 * One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
 * Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/09/04/del_node_supp.jpg" style="width: 350px; height: 255px;" />
 *
 * Example 2:
 *
 * Input: root = [5,3,6,2,4,null,7], key = 0
 * Output: [5,3,6,2,4,null,7]
 * Explanation: The tree does not contain a node with value = 0.
 *
 * Example 3:
 *
 * Input: root = [], key = 0
 * Output: []
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 10^4].
 * 	-10^5 <= Node.val <= 10^5
 * 	Each node has a unique value.
 * 	root is a valid binary search tree.
 * 	-10^5 <= key <= 10^5
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative
func deleteRootNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	next := findMin(root.Right)
	next.Left = root.Left
	return root.Right
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	node := root
	var prev *TreeNode
	for node != nil && node.Val != key {
		prev = node
		if key < node.Val {
			node = node.Left
		} else if key > node.Val {
			node = node.Right
		}
	}

	if prev == nil {
		return deleteRootNode(node)
	}
	if prev.Left == node {
		prev.Left = deleteRootNode(node)
	} else {
		prev.Right = deleteRootNode(node)
	}
	return root
}

// Recursive
func minValueNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		// find the inorder successor
		node := minValueNode(root.Right)
		val := node.Val
		deleteNode2(root, val)
		root.Val = val
		return root
	} else {
		if root.Val > key {
			root.Left = deleteNode2(root.Left, key)
		} else {
			root.Right = deleteNode2(root.Right, key)
		}
		return root
	}
}
