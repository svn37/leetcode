/**
 * Given a binary search tree with non-negative values, find the minimum <a href="https://en.wikipedia.org/wiki/Absolute_difference">absolute difference</a> between values of any two nodes.
 * Example:
 *
 * Input:
 *    1
 *     \
 *      3
 *     /
 *    2
 * Output:
 * 1
 * Explanation:
 * The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
 *
 *
 * Note:
 *
 * 	There are at least two nodes in this BST.
 * 	This question is the same as 783: <a href="https://leetcode.com/problems/minimum-distance-between-bst-nodes/">https://leetcode.com/problems/minimum-distance-between-bst-nodes/</a>
 *
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, prevVal, minDiff *int) {
	if root == nil {
		return
	}

	helper(root.Left, prevVal, minDiff)
	if *prevVal != math.MaxInt64 && root.Val-*prevVal < *minDiff {
		*minDiff = root.Val - *prevVal
	}
	*prevVal = root.Val
	helper(root.Right, prevVal, minDiff)
}

func getMinimumDifference(root *TreeNode) int {
	prevVal, minDiff := math.MaxInt64, math.MaxInt64

	helper(root, &prevVal, &minDiff)
	return minDiff
}
