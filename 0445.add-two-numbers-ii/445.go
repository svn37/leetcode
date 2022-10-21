/**
 * You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Follow up:<br />
 * What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
 *
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	// stack1 is larger stack
	if len(stack1) <= len(stack2) {
		stack1, stack2 = stack2, stack1
	}

	node := &ListNode{}
	var sum int

	for i, j := len(stack1)-1, len(stack2)-1; i >= 0; i, j = i-1, j-1 {
		sum += stack1[i]
		if j >= 0 {
			sum += stack2[j]
		}
		node.Val = sum % 10
		sum /= 10

		if i > 0 || sum > 0 {
			node = &ListNode{
				Val:  sum,
				Next: node,
			}
		}
	}

	return node
}
