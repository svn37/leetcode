/**
 * Given a binary tree, determine if it is a complete binary tree.
 *
 * <u>Definition of a complete binary tree from <a href="http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees" target="_blank">Wikipedia</a>:</u><br />
 * In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2^h nodes inclusive at the last level h.
 *
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-1.png" style="width: 180px; height: 145px;" />
 *
 *
 * Input: <span id="example-input-1-1">[1,2,3,4,5,6]</span>
 * Output: <span id="example-output-1">true</span>
 * <span>Explanation: </span>Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
 *
 *
 * <div>
 * Example 2:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-2.png" style="width: 200px; height: 145px;" />
 *
 *
 * Input: <span id="example-input-2-1">[1,2,3,4,5,null,7]</span>
 * Output: <span id="example-output-2">false</span>
 * Explanation: The node with value 7 isn't as far left as possible.<span>
 * </span>
 *
 * <div> </div>
 * </div>
 *
 * Note:
 *
 * <ol>
 * 	The tree will have between 1 and 100 nodes.
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

func (q *Queue) Empty() bool {
	return len(q.storage) == 0
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := &Queue{}
	queue.Push(root)

	for !queue.Empty() {
		node := queue.Poll()
		if node == nil {
			// check that only nils remain in the queue
			for !queue.Empty() {
				node := queue.Poll()
				if node != nil {
					return false
				}
			}
			return true
		}

		queue.Push(node.Left)
		queue.Push(node.Right)
	}
	return true
}
