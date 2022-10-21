/**
 * Given a string, sort it in decreasing order based on the frequency of characters.
 *
 * Example 1:
 *
 * Input:
 * "tree"
 *
 * Output:
 * "eert"
 *
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * "cccaaa"
 *
 * Output:
 * "cccaaa"
 *
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * "Aabb"
 *
 * Output:
 * "bbAa"
 *
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 *
 *
 */

package leetcode

import "bytes"

type Pair struct {
	count int
	char  rune
}

type Heap struct {
	storage []*Pair
}

func (h *Heap) ExtractMax() *Pair {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *Pair) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].count > h.storage[parent].count {
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
	largest := i

	if l < len(h.storage) && h.storage[l].count > h.storage[largest].count {
		largest = l
	}

	if r < len(h.storage) && h.storage[r].count > h.storage[largest].count {
		largest = r
	}

	if largest != i {
		h.storage[largest], h.storage[i] = h.storage[i], h.storage[largest]
		h.heapify(largest)
	}
}

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
	}

	h := &Heap{}
	for char, count := range m {
		h.Push(&Pair{count, char})
	}

	var b bytes.Buffer

	for !h.Empty() {
		maxVal := h.ExtractMax()
		for maxVal.count != 0 {
			b.WriteRune(maxVal.char)
			maxVal.count--
		}
	}

	return b.String()
}
