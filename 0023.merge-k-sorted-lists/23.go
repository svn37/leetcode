/**
 * You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 * Example 1:
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 *
 * Example 2:
 *
 * Input: lists = []
 * Output: []
 *
 * Example 3:
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 * Constraints:
 *
 * 	k == lists.length
 * 	0 <= k <= 10^4
 * 	0 <= lists[i].length <= 500
 * 	-10^4 <= lists[i][j] <= 10^4
 * 	lists[i] is sorted in ascending order.
 * 	The sum of lists[i].length won't exceed 10^4.
 *
 */

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap struct {
	storage []*ListNode
}

func NewHeap(a []*ListNode) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMin() *ListNode {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *ListNode) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].Val < h.storage[parent].Val {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) Empty() bool {
	return len(h.storage) == 0
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 != 0 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func (h *Heap) leftchild(i int) int {
	return 2*i + 1
}

func (h *Heap) rightchild(i int) int {
	return 2*i + 2
}

func (h *Heap) heapify(i int) {
	l := h.leftchild(i)
	r := h.rightchild(i)
	smallest := i

	if l < len(h.storage) && h.storage[l].Val < h.storage[smallest].Val {
		smallest = l
	}

	if r < len(h.storage) && h.storage[r].Val < h.storage[smallest].Val {
		smallest = r
	}

	if smallest != i {
		h.storage[smallest], h.storage[i] = h.storage[i], h.storage[smallest]
		h.heapify(smallest)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := &Heap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(list)
		}
	}

	head := &ListNode{} // dummy node
	node := head

	for !heap.Empty() {
		minVal := heap.ExtractMin()
		if minVal.Next != nil {
			heap.Push(minVal.Next)
		}

		node.Next = &ListNode{
			Val: minVal.Val,
		}
		node = node.Next
	}

	return head.Next
}
