/**
 * The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.
 *
 * Determine the maximum amount of money the thief can rob tonight without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3,null,3,null,1]
 *
 *      <font color="red">3</font>
 *     / \
 *    2   3
 *     \   \
 *      <font color="red">3   1
 * </font>
 * Output: 7
 * Explanation: Maximum amount of money the thief can rob = <font color="red" style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";">3</font><span style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";"> + </span><font color="red" style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";">3</font><span style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";"> + </span><font color="red" style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";">1</font><span style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";"> = </span><b style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";">7<span style="font-family: sans-serif, Arial, Verdana, "Trebuchet MS";">.</span>
 *
 * Example 2:
 *
 *
 * Input: [3,4,5,1,3,null,1]
 *
 *      3
 *     / \
 *    <font color="red">4</font>   <font color="red">5</font>
 *   / \   \
 *  1   3   1
 *
 * Output: 9
 * Explanation: Maximum amount of money the thief can rob = <font color="red">4</font> + <font color="red">5</font> = 9.
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

func robSub(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	leftArr := robSub(root.Left)
	rightArr := robSub(root.Right)

	// if current house is not robbed
	// we are free to choose whatever we want from subtrees
	left := max(leftArr[0], leftArr[1]) + max(rightArr[0], rightArr[1])
	// if cur is robbed
	// can't choose leftArr[1] and rightArr[1], because it'll alert the police
	right := root.Val + leftArr[0] + rightArr[0]

	return [2]int{left, right}
}

func rob(root *TreeNode) int {
	arr := robSub(root)
	return max(arr[0], arr[1])
}
