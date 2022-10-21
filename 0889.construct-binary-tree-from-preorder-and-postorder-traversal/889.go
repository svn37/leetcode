/**
 * Return any binary tree that matches the given preorder and postorder traversals.
 *
 * Values in the traversals pre and post are distinct positive integers.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: pre = <span id="example-input-1-1">[1,2,4,5,3,6,7]</span>, post = <span id="example-input-1-2">[4,5,2,6,7,3,1]</span>
 * Output: <span id="example-output-1">[1,2,3,4,5,6,7]</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 *
 * 	1 <= pre.length == post.length <= 30
 * 	pre[] and post[] are both permutations of 1, 2, ..., pre.length.
 * 	It is guaranteed an answer exists. If there exists multiple answers, you can return any of them.
 *
 * </div>
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(pre []int, post []int, preIndex, postIndex *int) *TreeNode {
	root := &TreeNode{
		Val: pre[*preIndex],
	}
	(*preIndex)++

	if root.Val != post[*postIndex] {
		root.Left = construct(pre, post, preIndex, postIndex)
	}

	if root.Val != post[*postIndex] {
		root.Right = construct(pre, post, preIndex, postIndex)
	}

	(*postIndex)++
	return root
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	preIndex, postIndex := 0, 0
	return construct(pre, post, &preIndex, &postIndex)
}
