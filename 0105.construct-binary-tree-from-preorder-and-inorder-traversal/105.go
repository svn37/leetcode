/**
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 *
 * Note:<br />
 * You may assume that duplicates do not exist in the tree.
 *
 * For example, given
 *
 *
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
 *
 * Return the following binary tree:
 *
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// original solution with map
func build(inStart, inEnd int, preorder []int, preStart *int, m map[int]int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	(*preStart)++
	root := &TreeNode{
		Val: preorder[*preStart],
	}
	if inStart == inEnd {
		return root
	}

	inIndex := m[root.Val]
	root.Left = build(inStart, inIndex-1, preorder, preStart, m)
	root.Right = build(inIndex+1, inEnd, preorder, preStart, m)

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// mapping from element in inorder set to its index, instead of linear search
	m := make(map[int]int)
	for i := range inorder {
		m[inorder[i]] = i
	}

	preStart := -1
	return build(0, len(inorder)-1, preorder, &preStart, m)
}

// correct solution
func helper(preStart, inStart, inEnd int, preorder []int, inorder []int) *TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}
	val := preorder[preStart]
	inIndex := 0 // index of current root in inorder
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == val {
			inIndex = i
		}
	}
	return &TreeNode{
		Val:   val,
		Left:  helper(preStart+1, inStart, inIndex-1, preorder, inorder),
		Right: helper(preStart+inIndex-inStart+1, inIndex+1, inEnd, preorder, inorder),
	}
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	return helper(0, 0, len(inorder)-1, preorder, inorder)
}
