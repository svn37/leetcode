/**
 * Given a binary search tree, return a balanced binary search tree with the same node values.
 *
 * A binary search tree is balanced if and only if the depth of the two subtrees of every node never differ by more than 1.
 *
 * If there is more than one answer, return any of them.
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/08/22/1515_ex1.png" style="width: 250px; height: 248px;" /><img alt="" src="https://assets.leetcode.com/uploads/2019/08/22/1515_ex1_out.png" style="width: 200px; height: 200px;" />
 *
 *
 * Input: root = [1,null,2,null,3,null,4,null,null]
 * Output: [2,1,3,null,null,null,4]
 * Explanation: This is not the only correct answer, [3,1,4,null,2,null,null] is also correct.
 *
 *
 *
 * Constraints:
 *
 *
 * 	The number of nodes in the tree is between 1 and 10^4.
 * 	The tree nodes will have distinct values between 1 and 10^5.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertToSortedArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	convertToSortedArray(root.Left, arr)
	*arr = append(*arr, root.Val)
	convertToSortedArray(root.Right, arr)
}

func buildTreeFromSortedArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	return &TreeNode{
		Val:   arr[mid],
		Left:  buildTreeFromSortedArray(arr[:mid]),
		Right: buildTreeFromSortedArray(arr[mid+1:]),
	}
}

func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	convertToSortedArray(root, &arr)
	return buildTreeFromSortedArray(arr)
}
