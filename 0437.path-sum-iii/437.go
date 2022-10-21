/**
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 *       10
 *      /  \
 *     5   -3
 *    / \    \
 *   3   2   11
 *  / \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// slow recursive method
func pathsum(root *TreeNode, curSum int) int {
	if root == nil {
		return 0
	}

	sum := pathsum(root.Left, curSum-root.Val) + pathsum(root.Right, curSum-root.Val)
	if root.Val == curSum {
		sum++
	}
	return sum
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathsum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// optimized prefixSum map method
func helper(root *TreeNode, cur, sum int, m map[int]int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	res := m[cur-sum]
	m[cur]++

	res += helper(root.Left, cur, sum, m) + helper(root.Right, cur, sum, m)
	m[cur]--
	return res
}

func pathSum2(root *TreeNode, sum int) int {
	// prefix sum map -- sum to frequency
	m := make(map[int]int)
	m[0] = 1 // so that m[cur-sum] from root gave the result
	return helper(root, 0, sum, m)
}
