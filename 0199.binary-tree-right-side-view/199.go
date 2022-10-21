/**
 * Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
 *
 * Example:
 *
 *
 * Input: [1,2,3,null,5,null,4]
 * Output: [1, 3, 4]
 * Explanation:
 *
 *    1            <---
 *  /   \
 * 2     3         <---
 *  \     \
 *   5     4       <---
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative, BFS
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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}

	queue := &Queue{}
	queue.Push(root)

	for !queue.Empty() {
		var rightmost *TreeNode
		for size := queue.Len(); size != 0; size-- {
			node := queue.Poll()
			if node == nil {
				continue
			}
			rightmost = node
			queue.Push(node.Left)
			queue.Push(node.Right)
		}
		if rightmost != nil {
			res = append(res, rightmost.Val)
		}
	}
	return res
}

// Recursive
func addNodes(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, root.Val)
	}

	addNodes(root.Right, level+1, res)
	addNodes(root.Left, level+1, res)
}

func rightSideView2(root *TreeNode) []int {
	res := []int{}
	addNodes(root, 0, &res)
	return res
}
