/**
 * Given a binary tree, write a function to get the maximum width of the given tree. The maximum width of a tree is the maximum width among all levels.
 * The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.
 * It is guaranteed that the answer will in the range of 32-bit signed integer.
 * Example 1:
 *
 * Input:
 *            1
 *          /   \
 *         3     2
 *        / \     \
 *       5   3     9
 * Output: 4
 * Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
 *
 * Example 2:
 *
 * Input:
 *           1
 *          /
 *         3
 *        / \
 *       5   3
 * Output: 2
 * Explanation: The maximum width existing in the third level with the length 2 (5,3).
 *
 * Example 3:
 *
 * Input:
 *           1
 *          / \
 *         3   2
 *        /
 *       5
 * Output: 2
 * Explanation: The maximum width existing in the second level with the length 2 (3,2).
 *
 * Example 4:
 *
 * Input:
 *           1
 *          / \
 *         3   2
 *        /     \
 *       5       9
 *      /         \
 *     6           7
 * Output: 8
 * Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).
 *
 *
 * Constraints:
 *
 * 	The given binary tree will have between 1 and 3000 nodes.
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func widthBST(root *TreeNode, level, pos int, widths *[][2]int) {
	if root == nil {
		return
	}

	if len(*widths) == level {
		*widths = append(*widths, [2]int{pos, pos})
	} else {
		(*widths)[level][0] = min(pos, (*widths)[level][0])
		(*widths)[level][1] = max(pos, (*widths)[level][1])
	}

	widthBST(root.Left, level+1, pos*2-1, widths)
	widthBST(root.Right, level+1, pos*2, widths)
}

func widthOfBinaryTree(root *TreeNode) int {
	widths := [][2]int{}
	widthBST(root, 0, 1, &widths)

	maxWidth := 0
	for _, width := range widths {
		maxWidth = max(maxWidth, width[1]-width[0]+1)
	}

	return maxWidth
}
