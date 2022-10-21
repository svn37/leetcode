/**
 * Given a binary tree root and a linked list with head as the first node.
 * Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary tree otherwise return False.
 * In this context downward path means a path that starts at some node and goes downwards.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/02/12/sample_1_1720.png" style="width: 220px; height: 280px;" />
 *
 * Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: true
 * Explanation: Nodes in blue form a subpath in the binary Tree.
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2020/02/12/sample_2_1720.png" style="width: 220px; height: 280px;" />
 *
 * Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: true
 *
 * Example 3:
 *
 * Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
 * Output: false
 * Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.
 *
 *
 * Constraints:
 *
 * 	1 <= node.val <= 100 for each node in the linked list and binary tree.
 * 	The given linked list will contain between 1 and 100 nodes.
 * 	The given binary tree will contain between 1 and 2500 nodes.
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return head.Val == root.Val && (dfs(head.Next, root.Left) || dfs(head.Next, root.Right))
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

/*
 * func helper(root *TreeNode, path, target string, lastVal int) bool {
 *     if root == nil {
 *         return false
 *     }
 *     newPath := fmt.Sprintf("%s-%d", path, root.Val)
 *     if root.Val == lastVal {
 *         ok := strings.HasSuffix(newPath, target)
 *         if ok {
 *             return true
 *         }
 *     }
 *     return helper(root.Left, newPath, target, lastVal) || helper(root.Right, newPath, target, lastVal)
 * }
 *
 * func isSubPath(head *ListNode, root *TreeNode) bool {
 *     var b bytes.Buffer
 *     var lastVal int
 *     for head != nil {
 *         b.WriteString(fmt.Sprintf("-%d", head.Val))
 *         if head.Next == nil {
 *             lastVal = head.Val
 *         }
 *         head = head.Next
 *     }
 *     return helper(root, "", b.String(), lastVal)
 * }
 */
