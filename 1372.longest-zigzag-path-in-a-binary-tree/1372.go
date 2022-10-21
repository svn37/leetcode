/**
 * Given a binary tree root, a ZigZag path for a binary tree is defined as follow:
 *
 *
 * 	Choose any node in the binary tree and a direction (right or left).
 * 	If the current direction is right then move to the right child of the current node otherwise move to the left child.
 * 	Change the direction from right to left or right to left.
 * 	Repeat the second and third step until you can't move in the tree.
 *
 *
 * Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
 *
 * Return the longest ZigZag path contained in that tree.
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png" style="width: 151px; height: 283px;" />
 *
 *
 * Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
 * Output: 3
 * Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
 *
 *
 * Example 2:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png" style="width: 120px; height: 253px;" />
 *
 *
 * Input: root = [1,1,1,null,1,null,null,1,1,null,1]
 * Output: 4
 * Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 	Each tree has at most 50000 nodes..
 * 	Each node's value is between [1, 100].
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// intead of dir we could use -curLen or smth like that, doesn't matter
func dfs(root *TreeNode, curLen int, dir bool, maxVal *int) {
	if root == nil {
		return
	}
	*maxVal = max(*maxVal, curLen)
	if dir {
		// left
		dfs(root.Left, 1, true, maxVal)
		dfs(root.Right, curLen+1, false, maxVal)
	} else {
		dfs(root.Left, curLen+1, true, maxVal)
		dfs(root.Right, 1, false, maxVal)
	}
}

func longestZigZag(root *TreeNode) int {
	maxVal := 0
	dfs(root.Left, 1, true, &maxVal)
	dfs(root.Right, 1, false, &maxVal)
	return maxVal
}
