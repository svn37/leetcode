/**
 * Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
 *
 *
 * For example:<br />
 * Given binary tree [3,9,20,null,null,15,7],<br />
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 *
 *
 * return its level order traversal as:<br />
 *
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
 *
 *
 */

package leetcode

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

func levelOrder(root *TreeNode) [][]int {
	queue := &Queue{}
	queue.Push(root)

	res := [][]int{}

	for !queue.Empty() {
		level := make([]int, 0)
		size := queue.Len()

		for size != 0 {
			node := queue.Poll()
			if node != nil {
				level = append(level, node.Val)

				queue.Push(node.Left)
				queue.Push(node.Right)
			}

			size--
		}

		if len(level) > 0 {
			res = append(res, level)
		}
	}
	return res
}
