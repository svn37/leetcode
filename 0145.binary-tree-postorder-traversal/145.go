/**
 * Given the root of a binary tree, return the postorder traversal of its nodes' values.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/28/pre1.jpg" style="width: 202px; height: 317px;" />
 * Input: root = [1,null,2,3]
 * Output: [3,2,1]
 *
 * Example 2:
 *
 * Input: root = []
 * Output: []
 *
 * Example 3:
 *
 * Input: root = [1]
 * Output: [1]
 *
 * Example 4:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/28/pre3.jpg" style="width: 202px; height: 197px;" />
 * Input: root = [1,2]
 * Output: [2,1]
 *
 * Example 5:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/08/28/pre2.jpg" style="width: 202px; height: 197px;" />
 * Input: root = [1,null,2]
 * Output: [2,1]
 *
 *
 * Constraints:
 *
 * 	The number of the nodes in the tree is in the range [0, 100].
 * 	-100 <= Node.val <= 100
 *
 *
 * Follow up:
 * Recursive solution is trivial, could you do it iteratively?
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

func (s *Stack) Top() *TreeNode {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}

	stack := &Stack{}
	stack.Push(root)
	stack.Push(root)

	for !stack.Empty() {
		node := stack.Pop()
		if !stack.Empty() && stack.Top() == node {
			if node.Right != nil {
				stack.Push(node.Right)
				stack.Push(node.Right)
			}

			if node.Left != nil {
				stack.Push(node.Left)
				stack.Push(node.Left)
			}
		} else {
			res = append(res, node.Val)
		}
	}
	return res
}
