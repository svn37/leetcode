/**
 * We have a collection of stones, each stone has a positive integer weight.
 * Each turn, we choose the two heaviest stones and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:
 *
 * 	If x == y, both stones are totally destroyed;
 * 	If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
 *
 * At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)
 *
 * Example 1:
 *
 * Input: [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
 *
 * Note:
 * <ol>
 * 	1 <= stones.length <= 30
 * 	1 <= stones[i] <= 1000
 * </ol>
 *
 */

package leetcode

type Heap struct {
	storage []int
}

func NewHeap(a []int) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMax() int {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val int) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i] > h.storage[parent] {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) HasOne() bool {
	return len(h.storage) == 1
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
	largest := i

	if l < len(h.storage) && h.storage[l] > h.storage[largest] {
		largest = l
	}

	if r < len(h.storage) && h.storage[r] > h.storage[largest] {
		largest = r
	}

	if largest != i {
		h.storage[largest], h.storage[i] = h.storage[i], h.storage[largest]
		h.heapify(largest)
	}
}

// alternative is sorting
func lastStoneWeight(stones []int) int {
	h := NewHeap(stones)
	for {
		if h.HasOne() {
			return h.ExtractMax()
		}
		maxVal1 := h.ExtractMax()
		maxVal2 := h.ExtractMax()
		if maxVal1 != maxVal2 {
			diff := maxVal1 - maxVal2
			if diff < 0 {
				diff *= (-1)
			}
			h.Push(diff)
		} else if h.Empty() {
			return 0
		}
	}
}
