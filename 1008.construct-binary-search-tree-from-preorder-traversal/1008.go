/**
 * Return the root node of a binary search tree that matches the given preorder traversal.
 * (Recall that a binary search tree is a binary tree where for every <font face="monospace">node</font>, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
 * It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.
 * Example 1:
 *
 * Input: <span id="example-input-1-1">[8,5,1,7,10,12]</span>
 * Output: <span id="example-output-1">[8,5,10,1,7,null,12]
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/03/06/1266.png" style="height: 200px; width: 306px;" /></span>
 *
 *
 * Constraints:
 *
 * 	1 <= preorder.length <= 100
 * 	1 <= preorder[i] <= 10^8
 * 	The values of preorder are distinct.
 *
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	storage []*TreeNode
}

func (s *Stack) Push(val *TreeNode) {
	s.storage = append(s.storage, val)
}

func (s *Stack) Top() *TreeNode {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Pop() *TreeNode {
	val := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func bstFromPreorder(preorder []int) *TreeNode {
	stack := &Stack{}
	root := &TreeNode{
		Val: math.MaxInt64,
	} // sentinel value
	stack.Push(root)

	for _, val := range preorder {
		node := &TreeNode{
			Val: val,
		}
		if val > stack.Top().Val {
			var last *TreeNode
			for !stack.Empty() && val > stack.Top().Val {
				last = stack.Pop()
			}
			last.Right = node
		} else {
			stack.Top().Left = node
		}
		stack.Push(node)
	}
	return root.Left
}
