/**
 * We are given a binary tree (with root node root), a target node, and an integer value K.
 *
 * Return a list of the values of all nodes that have a distance K from the target node.  The answer can be returned in any order.
 *
 *
 *
 * <ol>
 * </ol>
 *
 * <div>
 * Example 1:
 *
 *
 * Input: root = <span id="example-input-1-1">[3,5,1,6,2,0,8,null,null,7,4]</span>, target = <span id="example-input-1-2">5</span>, K = <span id="example-input-1-3">2</span>
 *
 * Output: <span id="example-output-1">[7,4,1]</span>
 *
 * Explanation:
 * The nodes that are a distance 2 from the target node (with value 5)
 * have values 7, 4, and 1.
 *
 * <img alt="" src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" style="width: 280px; height: 240px;" />
 *
 * Note that the inputs "root" and "target" are actually TreeNodes.
 * The descriptions of the inputs above are just serializations of these objects.
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	The given tree is non-empty.
 * 	Each node in the tree has unique values 0 <= node.val <= 500.
 * 	The target node is a node in the tree.
 * 	0 <= K <= 1000.
 * </ol>
 * </div>
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addKnodes(root *TreeNode, K int, kNodes *[]int) {
	if root == nil {
		return
	}

	if K == 0 {
		*kNodes = append(*kNodes, root.Val)
		return
	}

	addKnodes(root.Left, K-1, kNodes)
	addKnodes(root.Right, K-1, kNodes)
}

func dK(root *TreeNode, target *TreeNode, K int, kNodes *[]int) int {
	if root == nil {
		return -1
	}

	if root == target {
		addKnodes(root, K, kNodes)
		return K - 1
	}

	left := dK(root.Left, target, K, kNodes)
	if left > 0 {
		addKnodes(root.Right, left-1, kNodes)
		return left - 1
	} else if left == 0 {
		*kNodes = append(*kNodes, root.Val)
		return -1
	}

	right := dK(root.Right, target, K, kNodes)
	if right > 0 {
		addKnodes(root.Left, right-1, kNodes)
		return right - 1
	} else if right == 0 {
		*kNodes = append(*kNodes, root.Val)
		return -1
	}

	return -1
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	kNodes := []int{}
	dK(root, target, K, &kNodes)
	return kNodes
}
