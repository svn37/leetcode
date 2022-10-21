/**
 * Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
 * Return the smallest level X such that the sum of all the values of nodes at level X is maximal.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" style="width: 200px; height: 175px;" />
 * Input: root = [1,7,0,7,-8,null,null]
 * Output: 2
 * Explanation:
 * Level 1 sum = 1.
 * Level 2 sum = 7 + 0 = 7.
 * Level 3 sum = 7 + -8 = -1.
 * So we return the level with the maximum sum which is level 2.
 *
 * Example 2:
 *
 * Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [1, 10^4].
 * 	-10^5 <= Node.val <= 10^5
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
func helper(root *TreeNode, curLevel int, sums *[]int) {
	if root == nil {
		return
	}

	if curLevel == len(*sums) {
		*sums = append(*sums, 0)
	}
	(*sums)[curLevel] += root.Val

	helper(root.Left, curLevel+1, sums)
	helper(root.Right, curLevel+1, sums)
}

func maxLevelSum(root *TreeNode) int {
	sums := make([]int, 1)
	helper(root, 1, &sums)

	smallestLevel, maxSum := 0, 0
	for level, sum := range sums {
		if sum > maxSum {
			maxSum = sum
			smallestLevel = level
		}
	}
	return smallestLevel
}

// Iterative BFS search
type Queue struct {
	nodes []*TreeNode
}

func (q *Queue) Push(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Poll() *TreeNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *Queue) Len() int {
	return len(q.nodes)
}

func (q *Queue) Empty() bool {
	return len(q.nodes) == 0
}

func maxLevelSum2(root *TreeNode) (maxLevel int) {
	queue := &Queue{}
	queue.Push(root)

	smallestLevel := 0
	maxSum, curSum := 0, 0

	for level := 1; !queue.Empty(); level++ {
		for length := queue.Len(); length > 0; length-- {
			node := queue.Poll()
			if node != nil {
				curSum += node.Val

				queue.Push(node.Left)
				queue.Push(node.Right)
			}
		}
		if curSum > maxSum {
			maxSum = curSum
			smallestLevel = level
		}
		curSum = 0
	}
	return smallestLevel
}
