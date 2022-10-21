/**
 *
 * Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:
 * <ol>
 * The root is the maximum number in the array.
 * The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
 * The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
 * </ol>
 *
 *
 *
 * Construct the maximum tree by the given array and output the root node of this tree.
 *
 *
 * Example 1:<br />
 *
 * Input: [3,2,1,6,0,5]
 * Output: return the tree root node representing the following tree:
 *
 *       6
 *     /   \
 *    3     5
 *     \    /
 *      2  0
 *        \
 *         1
 *
 *
 *
 * Note:<br>
 * <ol>
 * The size of the given array will be in the range [1,1000].
 * </ol>
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

func (s *Stack) PeekLeft() *TreeNode {
	return s.storage[0]
}

func (s *Stack) PeekRight() *TreeNode {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Pop() *TreeNode {
	node := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return node
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

// Method 1. Use monotonic stack.
func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := &Stack{}
	for i := range nums {
		node := &TreeNode{
			Val: nums[i],
		}
		for !stack.Empty() && stack.PeekRight().Val < node.Val {
			node.Left = stack.Pop()
		}

		if !stack.Empty() {
			stack.PeekRight().Right = node
		}
		stack.Push(node)
	}
	return stack.PeekLeft()
}

// Method 2. Simple recursive, search for the largest element in each array.
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIdx := 0
	for i := range nums {
		if nums[maxIdx] <= nums[i] {
			maxIdx = i
		}
	}
	return &TreeNode{
		Val:   nums[maxIdx],
		Left:  constructMaximumBinaryTree2(nums[:maxIdx]),
		Right: constructMaximumBinaryTree2(nums[maxIdx+1:]),
	}
}
