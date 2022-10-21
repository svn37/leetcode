/**
 * Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
 * Example:
 *
 * Input: 3
 * Output:
 * [
 *   [1,null,3,2],
 *   [3,2,null,1],
 *   [3,1,null,null,2],
 *   [2,1,3],
 *   [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *    1         3     3      2      1
 *     \       /     /      / \      \
 *      3     2     1      1   3      2
 *     /     /       \                 \
 *    2     1         2                 3
 *
 *
 * Constraints:
 *
 * 	0 <= n <= 8
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1. Iterative
func clone(node *TreeNode, offset int) *TreeNode {
	if node == nil {
		return nil
	}

	return &TreeNode{
		Val:   node.Val + offset,
		Left:  clone(node.Left, offset),
		Right: clone(node.Right, offset),
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	res := make([][]*TreeNode, n+1)
	res[0] = []*TreeNode{nil}

	for i := 1; i <= n; i++ {
		res[i] = make([]*TreeNode, 0)
		for j := 0; j < i; j++ {
			for _, lnode := range res[j] {
				for _, rnode := range res[i-j-1] {
					res[i] = append(res[i], &TreeNode{
						Val:   j + 1,
						Left:  lnode,
						Right: clone(rnode, j+1),
					})
				}
			}
		}
	}
	return res[n]
}

// Method 2. Recursive
func genTrees(start, end int, memo map[[2]int][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	if start == end {
		return []*TreeNode{{
			Val: start,
		}}
	}

	pos := [2]int{start, end}
	if cache, ok := memo[pos]; ok {
		return cache
	}

	nodes := make([]*TreeNode, 0, end-start+1)

	for i := start; i <= end; i++ {
		left := genTrees(start, i-1, memo)
		right := genTrees(i+1, end, memo)

		for _, lnode := range left {
			for _, rnode := range right {
				nodes = append(nodes, &TreeNode{
					Val:   i,
					Left:  lnode,
					Right: rnode,
				})
			}
		}
	}
	memo[pos] = nodes

	return nodes
}

func generateTrees2(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	memo := make(map[[2]int][]*TreeNode)
	return genTrees(1, n, memo)
}
