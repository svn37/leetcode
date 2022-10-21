/**
 * Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
 *
 * Calling next() will return the next smallest number in the BST.
 *
 *
 *
 *
 *
 *
 * Example:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png" style="width: 189px; height: 178px;" />
 *
 *
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // return 3
 * iterator.next();    // return 7
 * iterator.hasNext(); // return true
 * iterator.next();    // return 9
 * iterator.hasNext(); // return true
 * iterator.next();    // return 15
 * iterator.hasNext(); // return true
 * iterator.next();    // return 20
 * iterator.hasNext(); // return false
 *
 *
 *
 *
 * Note:
 *
 *
 * 	next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
 * 	You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.
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
	items []*TreeNode
}

func (s *Stack) Push(val *TreeNode) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() *TreeNode {
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

type BSTIterator struct {
	stack *Stack
}

func addLeftNodes(stack *Stack, node *TreeNode) {
	for node != nil {
		stack.Push(node)
		node = node.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	stack := &Stack{}
	addLeftNodes(stack, root)
	return BSTIterator{
		stack: stack,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	smallest := this.stack.Pop()
	addLeftNodes(this.stack, smallest.Right)
	return smallest.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.stack.Empty()
}
