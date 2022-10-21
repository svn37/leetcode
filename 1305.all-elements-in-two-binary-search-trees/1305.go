/**
 * Given two binary search trees root1 and root2.
 * Return a list containing all the integers from both trees sorted in ascending order.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png" style="width: 457px; height: 207px;" />
 * Input: root1 = [2,1,4], root2 = [1,0,3]
 * Output: [0,1,1,2,3,4]
 *
 * Example 2:
 *
 * Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
 * Output: [-10,0,0,1,2,5,7,10]
 *
 * Example 3:
 *
 * Input: root1 = [], root2 = [5,1,7,0,2]
 * Output: [0,1,2,5,7]
 *
 * Example 4:
 *
 * Input: root1 = [0,-10,10], root2 = []
 * Output: [-10,0,10]
 *
 * Example 5:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png" style="width: 352px; height: 197px;" />
 * Input: root1 = [1,null,8], root2 = [8,1]
 * Output: [1,1,8,8]
 *
 *
 * Constraints:
 *
 * 	Each tree has at most 5000 nodes.
 * 	Each node's value is between [-10^5, 10^5].
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// could be done with iterator, similar to "173. Binary Search Tree Iterator"
// in fact, it could be argued that separate full traversals and sorting is better,
// because timsort in python performs better when there are sorted runs inside the data.
func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	inorder(root1, &arr1)
	inorder(root2, &arr2)

	res := make([]int, len(arr1)+len(arr2))
	i, a, b := 0, 0, 0
	for i < len(res) && a < len(arr1) && b < len(arr2) {
		if arr1[a] < arr2[b] {
			res[i] = arr1[a]
			a++
		} else {
			res[i] = arr2[b]
			b++
		}
		i++
	}
	if a < len(arr1) {
		copy(res[i:], arr1[a:])
	} else {
		copy(res[i:], arr2[b:])
	}

	return res
}
