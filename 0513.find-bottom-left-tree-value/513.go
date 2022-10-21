/**
 *
 * Given a binary tree, find the leftmost value in the last row of the tree.
 *
 *
 * Example 1:<br />
 *
 * Input:
 *
 *     2
 *    / \
 *   1   3
 *
 * Output:
 * 1
 *
 *
 *
 *   Example 2: <br>
 *
 * Input:
 *
 *         1
 *        / \
 *       2   3
 *      /   / \
 *     4   5   6
 *        /
 *       7
 *
 * Output:
 * 7
 *
 *
 *
 * Note:
 * You may assume the tree (i.e., the given root node) is not NULL.
 *
 */

package leetcode

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

func (q *Queue) Empty() bool {
	return len(q.nodes) == 0
}

func findBottomLeftValue(root *TreeNode) int {
	queue := &Queue{}
	queue.Push(root)

	for !queue.Empty() {
		root = queue.Poll()
		if root.Right != nil {
			queue.Push(root.Right)
		}
		if root.Left != nil {
			queue.Push(root.Left)
		}
	}
	return root.Val
}
