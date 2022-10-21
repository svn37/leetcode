/**
 * We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.
 *
 * Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.
 *
 * Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).
 *
 * Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: <span id="example-input-1-1">[2,1,5]</span>
 * Output: <span id="example-output-1">[5,5,0]</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: <span id="example-input-2-1">[2,7,4,3,5]</span>
 * Output: <span id="example-output-2">[7,0,5,5,0]</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: <span id="example-input-3-1">[1,7,5,1,9,2,5,1]</span>
 * Output: <span id="example-output-3">[7,9,9,9,0,5,0,0]</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	<span>1 <= node.val <= 10^9</span><span> for each node in the linked list.</span>
 * 	The given list has length in the range [0, 10000].
 * </ol>
 * </div>
 * </div>
 * </div>
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// without converting to array, my original solution
func nextLarger(head *ListNode, m map[*ListNode]*ListNode) *ListNode {
	if head.Next == nil || head.Next.Val > head.Val {
		m[head] = head.Next
		return head.Next
	}

	nextlarger := nextLarger(head.Next, m)
	m[head] = nextlarger
	return nextlarger
}

func nextLargerNodes(head *ListNode) []int {
	m := make(map[*ListNode]*ListNode)

	node := head
	for node != nil {
		node = nextLarger(node, m)
	}

	res := []int{}
	for head != nil {
		candidate := head
		for {
			candidate = m[candidate]
			if candidate == nil {
				res = append(res, 0)
				break
			}

			if candidate.Val > head.Val {
				res = append(res, candidate.Val)
				break
			}
		}
		head = head.Next
	}
	return res
}

// convert to array, maintain a monotonically decreasing stack
type Stack struct {
	storage []int
}

func (s *Stack) Push(i int) {
	s.storage = append(s.storage, i)
}

func (s *Stack) Pop() int {
	i := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return i
}

func (s *Stack) Top() int {
	return s.storage[len(s.storage)-1]
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func nextLargerNodes2(head *ListNode) []int {
	res := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		res = append(res, node.Val)
	}
	stack := &Stack{}
	for i := len(res) - 1; i >= 0; i-- {
		val := res[i]
		for !stack.Empty() && stack.Top() <= res[i] {
			stack.Pop()
		}
		if stack.Empty() {
			res[i] = 0
		} else {
			res[i] = stack.Top()
		}
		stack.Push(val)
	}
	return res
}
