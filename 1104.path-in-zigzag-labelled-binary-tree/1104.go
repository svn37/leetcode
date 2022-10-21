/**
 * In an infinite binary tree where every node has two children, the nodes are labelled in row order.
 * In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/06/24/tree.png" style="width: 300px; height: 138px;" />
 * Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.
 *
 * Example 1:
 *
 * Input: label = 14
 * Output: [1,3,4,14]
 *
 * Example 2:
 *
 * Input: label = 26
 * Output: [1,2,6,10,26]
 *
 *
 * Constraints:
 *
 * 	1 <= label <= 10^6
 *
 */

package leetcode

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func findParent(label int) int {
	if label%2 == 0 {
		return label / 2
	}
	return (label - 1) / 2
}

func pathInZigZagTree(label int) []int {
	res := []int{}

	// find the level
	pow := 1
	maxLabel := 0
	level := 0
	for maxLabel < label {
		maxLabel += pow
		level++
		pow *= 2
	}

	parent := label
	for parent > 0 {
		res = append(res, parent)
		parent = findParent(parent)

		pow /= 2
		level--
		// invert parent
		parent = pow - 1 - (parent - pow/2)
	}

	reverse(res)

	return res
}
