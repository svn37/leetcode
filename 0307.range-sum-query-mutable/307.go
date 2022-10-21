/**
 * Given an integer array nums, find the sum of the elements between indices i and j (i &le; j), inclusive.
 * The update(i, val) function modifies nums by updating the element at index i to val.
 * Example:
 *
 * Given nums = [1, 3, 5]
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 *
 *
 * Constraints:
 *
 * 	The array is only modifiable by the update function.
 * 	You may assume the number of calls to update and sumRange function is distributed evenly.
 * 	0 <= i <= j <= nums.length - 1
 *
 */

package leetcode

// segment tree node
type Node struct {
	start, end  int
	left, right *Node
	sum         int
}

type NumArray struct {
	root *Node
}

func Constructor(nums []int) NumArray {
	root := buildTree(nums, 0, len(nums)-1)
	return NumArray{root}
}

func buildTree(nums []int, start, end int) *Node {
	if start > end {
		return nil
	}
	node := &Node{
		start: start,
		end:   end,
	}
	if start == end {
		node.sum = nums[start]
	} else {
		// overflow trick
		mid := start + (end-start)/2
		node.left = buildTree(nums, start, mid)
		node.right = buildTree(nums, mid+1, end)
		node.sum = node.left.sum + node.right.sum
	}
	return node
}

func update(node *Node, i int, val int) {
	if node.start == node.end {
		node.sum = val
	} else {
		mid := node.start + (node.end-node.start)/2
		if i <= mid {
			update(node.left, i, val)
		} else {
			update(node.right, i, val)
		}
		node.sum = node.left.sum + node.right.sum
	}
}

func sumRange(node *Node, start, end int) int {
	if node.start == start && node.end == end {
		return node.sum
	} else {
		mid := node.start + (node.end-node.start)/2
		if mid >= end {
			return sumRange(node.left, start, end)
		} else if mid+1 <= start {
			return sumRange(node.right, start, end)
		} else {
			return sumRange(node.left, start, mid) + sumRange(node.right, mid+1, end)
		}
	}
}

func (this *NumArray) Update(i int, val int) {
	update(this.root, i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	return sumRange(this.root, i, j)
}
