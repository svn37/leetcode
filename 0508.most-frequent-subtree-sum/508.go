/**
 *
 * Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.
 *
 *
 * Examples 1<br>
 * Input:
 *
 *   5
 *  /  \
 * 2   -3
 *
 * return [2, -3, 4], since all the values happen only once, return all of them in any order.
 *
 *
 * Examples 2<br>
 * Input:
 *
 *   5
 *  /  \
 * 2   -5
 *
 * return [2], since 2 happens twice, however -5 only occur once.
 *
 *
 * Note:
 * You may assume the sum of values in any subtree is in the range of 32-bit signed integer.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countFrequentTreeSum(root *TreeNode, m map[int]int, val *[]int, maxCount *int) int {
	if root == nil {
		return 0
	}
	sum := root.Val + countFrequentTreeSum(root.Left, m, val, maxCount) + countFrequentTreeSum(root.Right, m, val, maxCount)
	m[sum]++
	if *maxCount < m[sum] {
		*val = []int{sum}
		*maxCount = m[sum]
	} else if *maxCount == m[sum] {
		*val = append(*val, sum)
	}
	return sum
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)
	val, maxCount := []int{}, 0
	countFrequentTreeSum(root, m, &val, &maxCount)
	return val
}
