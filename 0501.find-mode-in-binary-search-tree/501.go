/**
 * Given a binary search tree (BST) with duplicates, find all the <a href="https://en.wikipedia.org/wiki/Mode_(statistics)" target="_blank">mode(s)</a> (the most frequently occurred element) in the given BST.
 *
 * Assume a BST is defined as follows:
 *
 *
 * 	The left subtree of a node contains only nodes with keys less than or equal to the node's key.
 * 	The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
 * 	Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * For example:<br />
 * Given BST [1,null,2,2],
 *
 *
 *    1
 *     \
 *      2
 *     /
 *    2
 *
 *
 *
 *
 * return [2].
 *
 * Note: If a tree has more than one mode, you can return them in any order.
 *
 * Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
 *
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Simple recursive traversal
func inorder(root *TreeNode, curFreq, maxFreq, prevVal *int, modes *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, curFreq, maxFreq, prevVal, modes)

	if *prevVal != math.MinInt64 && *prevVal == root.Val {
		*curFreq++
	} else {
		*curFreq = 1
	}
	*prevVal = root.Val

	if *curFreq >= *maxFreq {
		if *curFreq > *maxFreq {
			*modes = (*modes)[:0]
			*maxFreq++
		}
		*modes = append(*modes, root.Val)
	}

	inorder(root.Right, curFreq, maxFreq, prevVal, modes)
}

// true O(1) space Morris traversal
func handleValue(node *TreeNode, curFreq, maxFreq, prevVal *int, modes *[]int) {
	if *prevVal != math.MinInt64 && *prevVal == node.Val {
		*curFreq++
	} else {
		*curFreq = 1
	}
	*prevVal = node.Val

	if *curFreq >= *maxFreq {
		if *curFreq > *maxFreq {
			*modes = (*modes)[:0]
			*maxFreq++
		}
		*modes = append(*modes, node.Val)
	}
}

// Morris inorder traversal.
func inorder2(node *TreeNode, curFreq, maxFreq, prevVal *int, modes *[]int) {
	for node != nil {
		if node.Left == nil {
			handleValue(node, curFreq, maxFreq, prevVal, modes)
			node = node.Right
		} else {
			temp := node.Left
			for temp.Right != node && temp.Right != nil {
				temp = temp.Right
			}

			if temp.Right == node {
				// threading already exists
				temp.Right = nil
				handleValue(node, curFreq, maxFreq, prevVal, modes)
				node = node.Right
			} else {
				// create threading
				temp.Right = node
				node = node.Left
			}
		}
	}
}

func findMode(root *TreeNode) []int {
	// global variables
	modes := []int{}
	prevVal := math.MinInt64
	curFreq, maxFreq := 0, 0

	inorder(root, &curFreq, &maxFreq, &prevVal, &modes)
	return modes
}
