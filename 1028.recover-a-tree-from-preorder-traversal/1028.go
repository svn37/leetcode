/**
 * We run a preorder depth-first search (DFS) on the root of a binary tree.
 * At each node in this traversal, we output D dashes (where D is the depth of this node), then we output the value of this node.  If the depth of a node is D, the depth of its immediate child is D + 1.  The depth of the root node is 0.
 * If a node has only one child, that child is guaranteed to be the left child.
 * Given the output S of this traversal, recover the tree and return its root.
 *
 * Example 1:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/04/08/recover-a-tree-from-preorder-traversal.png" style="width: 320px; height: 200px;" />
 * Input: S = "1-2--3--4-5--6--7"
 * Output: [1,2,5,3,4,6,7]
 *
 * Example 2:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114101-pm.png" style="width: 256px; height: 250px;" />
 * Input: S = "1-2--3---4-5--6---7"
 * Output: [1,2,5,3,null,6,null,4,null,7]
 *
 * Example 3:
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114955-pm.png" style="width: 276px; height: 250px;" />
 * Input: S = "1-401--349---90--88"
 * Output: [1,401,null,349,88,90]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the original tree is in the range [1, 1000].
 * 	1 <= Node.val <= 10^9
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	nodes []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() *TreeNode {
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node
}

func (s *Stack) Peek() *TreeNode {
	return s.nodes[len(s.nodes)-1]
}

func (s *Stack) Empty() bool {
	return len(s.nodes) == 0
}

func recoverFromPreorder(S string) *TreeNode {
	s := &Stack{}
	s.Push(&TreeNode{}) // dummy node

	val := 0
	d := 0        // d is the current number of consecutive '-'s
	curDepth := 0 // curDepth is the current depth of the stack

	for i := 0; i <= len(S); i++ {
		if i == len(S) || S[i] == '-' {
			if val > 0 {
				node := &TreeNode{
					Val: val,
				}
				for curDepth != d {
					s.Pop()
					curDepth--
				}

				root := s.Peek()
				if root.Left == nil {
					root.Left = node
				} else {
					root.Right = node
				}

				s.Push(node)
				curDepth++

				val = 0
				d = 0
			}

			d++
		} else {
			val = val*10 + int(S[i]-'0')
		}
	}

	var root *TreeNode
	for !s.Empty() {
		root = s.Pop()
	}

	return root.Left
}
