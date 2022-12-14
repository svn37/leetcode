/**
 * Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.
 *
 * You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.
 *
 * Example 1:
 *
 *
 * Input:
 * 	Tree 1                     Tree 2
 *           1                         2
 *          / \                       / \
 *         3   2                     1   3
 *        /                           \   \
 *       5                             4   7
 * Output:
 * Merged tree:
 * 	     3
 * 	    / \
 * 	   4   5
 * 	  / \   \
 * 	 5   4   7
 *
 *
 *
 *
 * Note: The merging process must start from the root nodes of both trees.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	var newVal int
	var left1, right1, left2, right2 *TreeNode
	if t1 != nil {
		newVal += t1.Val
		left1 = t1.Left
		right1 = t1.Right
	}

	if t2 != nil {
		newVal += t2.Val
		left2 = t2.Left
		right2 = t2.Right
	}

	return &TreeNode{
		Val:   newVal,
		Left:  merge(left1, left2),
		Right: merge(right1, right2),
	}
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return merge(t1, t2)
}
