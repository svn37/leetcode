/**
 * Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
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
 * return its zigzag level order traversal as:<br />
 *
 * [
 *   [3],
 *   [20,9],
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

type Stack struct {
	storage []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.storage = append(s.storage, node)
}

func (s *Stack) Pop() *TreeNode {
	node := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return node
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	stack := &Stack{}
	temp := &Stack{}

	stack.Push(root)
	level := make([]int, 0)
	var zigZag bool

	for !stack.Empty() {
		node := stack.Pop()
		if node != nil {
			level = append(level, node.Val)

			if zigZag {
				temp.Push(node.Right)
				temp.Push(node.Left)
			} else {
				temp.Push(node.Left)
				temp.Push(node.Right)
			}
		}

		if stack.Empty() {
			zigZag = !zigZag
			stack, temp = temp, stack

			if len(level) > 0 {
				res = append(res, append(level[:0:0], level...))
				level = level[:0]
			}
		}
	}
	return res
}
