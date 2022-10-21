/**
 * Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
 *
 * Example 1:<br />
 *
 * Input:
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * Output: [3, 14.5, 11]
 * Explanation:
 * The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
 *
 *
 *
 * Note:<br>
 * <ol>
 * The range of node's value is in the range of 32-bit signed integer.
 * </ol>
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

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}

	queue := &Queue{}
	queue.Push(root)

	for !queue.Empty() {
		sum, count := 0, 0.0
		for l := queue.Len(); l > 0; l-- {
			node := queue.Poll()
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}

			sum += node.Val
			count++
		}
		res = append(res, float64(sum)/count)
	}
	return res
}
