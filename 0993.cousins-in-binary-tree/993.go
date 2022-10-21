/**
 * In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
 * Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
 * We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
 * Return true if and only if the nodes corresponding to the values x and y are cousins.
 *
 * Example 1:<br />
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/02/12/q1248-01.png" style="width: 180px; height: 160px;" />
 *
 * Input: root = <span id="example-input-1-1">[1,2,3,4]</span>, x = <span id="example-input-1-2">4</span>, y = <span id="example-input-1-3">3</span>
 * Output: <span id="example-output-1">false</span>
 *
 * <div>
 * Example 2:<br />
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/02/12/q1248-02.png" style="width: 201px; height: 160px;" />
 *
 * Input: root = <span id="example-input-2-1">[1,2,3,null,4,null,5]</span>, x = <span id="example-input-2-2">5</span>, y = <span id="example-input-2-3">4</span>
 * Output: <span id="example-output-2">true</span>
 *
 * <div>
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/02/13/q1248-03.png" style="width: 156px; height: 160px;" />
 *
 * Input: root = <span id="example-input-3-1">[1,2,3,null,4]</span>, x = 2, y = 3
 * Output: <span id="example-output-3">false</span>
 * </div>
 * </div>
 *
 * Constraints:
 *
 * 	The number of nodes in the tree will be between 2 and 100.
 * 	Each node has a unique integer value from 1 to 100.
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
	storage []*TreeNode
}

func (q *Queue) Push(node *TreeNode) {
	q.storage = append(q.storage, node)
}

func (q *Queue) Poll() *TreeNode {
	node := q.storage[0]
	q.storage = q.storage[1:]
	return node
}

func (q *Queue) Len() int {
	return len(q.storage)
}

func (q *Queue) Empty() bool {
	return len(q.storage) == 0
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := &Queue{}
	queue.Push(root)

	// level where one cousin was found
	level := math.MaxInt64
	curLevel := 0
	for !queue.Empty() || level < curLevel {
		for size := queue.Len(); size > 0; size-- {
			node := queue.Poll()
			if node == nil {
				continue
			}

			if node.Left != nil && node.Right != nil &&
				(node.Left.Val == x && node.Right.Val == y ||
					node.Right.Val == x && node.Left.Val == y) {
				return false
			}

			if node.Val == x || node.Val == y {
				if level == math.MaxInt64 {
					level = curLevel
				} else {
					return level == curLevel
				}
			}

			queue.Push(node.Left)
			queue.Push(node.Right)
		}
		curLevel++
	}
	return false
}
