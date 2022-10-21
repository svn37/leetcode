/**
 * Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).
 *
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/21/largest_e1.jpg" style="width: 450px; height: 258px;" />
 * Input: root = [1,3,2,5,3,null,9]
 * Output: [1,3,9]
 *
 * Example 2:
 *
 * Input: root = [1,2,3]
 * Output: [1,3]
 *
 * Example 3:
 *
 * Input: root = [1]
 * Output: [1]
 *
 * Example 4:
 *
 * Input: root = [1,null,2]
 * Output: [1,2]
 *
 * Example 5:
 *
 * Input: root = []
 * Output: []
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree will be in the range [0, 10^4].
 * 	-2^31 <= Node.val <= 2^31 - 1
 *
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	nodes []*TreeNode
}

func (q *Queue) Push(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Poll() *TreeNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *Queue) Len() int {
	return len(q.nodes)
}

func (q *Queue) Empty() bool {
	return len(q.nodes) == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestValues(root *TreeNode) []int {
	queue := &Queue{}
	queue.Push(root)

	maxVals := []int{}
	for !queue.Empty() {
		maxVal := math.MinInt64
		for size := queue.Len(); size > 0; size-- {
			root = queue.Poll()
			if root == nil {
				continue
			}
			maxVal = max(maxVal, root.Val)
			queue.Push(root.Left)
			queue.Push(root.Right)
		}
		if maxVal != math.MinInt64 {
			maxVals = append(maxVals, maxVal)
		}
	}
	return maxVals
}
