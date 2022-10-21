/**
 * You are given the root of a binary search tree (BST), where exactly two nodes of the tree were swapped by mistake. Recover the tree without changing its structure.
 * Follow up: A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/28/recover1.jpg" style="width: 422px; height: 302px;" />
 * Input: root = [1,3,null,null,2]
 * Output: [3,1,null,null,2]
 * Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/10/28/recover2.jpg" style="width: 581px; height: 302px;" />
 * Input: root = [3,1,4,null,null,2]
 * Output: [2,1,4,null,null,3]
 * Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [2, 1000].
 * 	-2^31 <= Node.val <= 2^31 - 1
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swap(first, second *TreeNode) {
	temp := first.Val
	first.Val = second.Val
	second.Val = temp
}

func checkSwapped(node *TreeNode, prev, first, second **TreeNode) {
	if *prev != nil && *first == nil && (*prev).Val > node.Val {
		*first = *prev
	}

	if *prev != nil && *first != nil && (*prev).Val > node.Val {
		*second = node
	}

	*prev = node
}

// Morris inorder traversal.
func traverse(node *TreeNode, prev, first, second **TreeNode) {
	for node != nil {
		if node.Left == nil {
			checkSwapped(node, prev, first, second)

			node = node.Right
		} else {
			temp := node.Left
			for temp.Right != node && temp.Right != nil {
				temp = temp.Right
			}

			if temp.Right == node {
				// threading already exists
				temp.Right = nil

				checkSwapped(node, prev, first, second)

				node = node.Right
			} else {
				// create threading
				temp.Right = node
				node = node.Left
			}
		}
	}
}

func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	traverse(root, &prev, &first, &second)
	swap(first, second)
}
