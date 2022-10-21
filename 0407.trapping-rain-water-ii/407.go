/**
 * Given an m x n matrix of positive integers representing the height of each unit cell in a 2D elevation map, compute the volume of water it is able to trap after raining.
 * Example:
 *
 * Given the following 3x6 height map:
 * [
 *   [1,4,3,1,3,2],
 *   [3,2,1,3,2,4],
 *   [2,3,3,2,3,1]
 * ]
 * Return 4.
 *
 * <img src="https://assets.leetcode.com/uploads/2018/10/13/rainwater_empty.png" style="width: 100%; max-width: 500px;" />
 * The above image represents the elevation map [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] before the rain.
 *
 * <img src="https://assets.leetcode.com/uploads/2018/10/13/rainwater_fill.png" style="width: 100%; max-width: 500px;" />
 * After the rain, water is trapped between the blocks. The total volume of water trapped is 4.
 *
 * Constraints:
 *
 * 	1 <= m, n <= 110
 * 	0 <= heightMap[i][j] <= 20000
 *
 */

package leetcode

type Heap struct {
	storage []*Cell
}

func NewHeap(a []*Cell) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMin() *Cell {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *Cell) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].height < h.storage[parent].height {
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

	if l < len(h.storage) && h.storage[l].height < h.storage[smallest].height {
		smallest = l
	}

	if r < len(h.storage) && h.storage[r].height < h.storage[smallest].height {
		smallest = r
	}

	if smallest != i {
		h.storage[smallest], h.storage[i] = h.storage[i], h.storage[smallest]
		h.heapify(smallest)
	}
}

type Cell struct {
	height, x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m < 3 {
		return 0
	}

	n := len(heightMap[0])
	if n < 3 {
		return 0
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	heap := &Heap{}

	// add cells on the border to the priority queue
	for i := 0; i < m; i++ {
		heap.Push(&Cell{heightMap[i][0], i, 0})
		visited[i][0] = true

		heap.Push(&Cell{heightMap[i][n-1], i, n - 1})
		visited[i][n-1] = true
	}

	for j := 1; j < n-1; j++ {
		heap.Push(&Cell{heightMap[0][j], 0, j})
		visited[0][j] = true

		heap.Push(&Cell{heightMap[m-1][j], m - 1, j})
		visited[m-1][j] = true
	}

	volume := 0
	dirs := []int{0, 1, 0, -1, 0}
	for !heap.Empty() {
		cell := heap.ExtractMin()
		for i := 0; i < 4; i++ {
			dx := cell.x + dirs[i]
			dy := cell.y + dirs[i+1]

			if dx >= 0 && dx < m && dy >= 0 && dy < n && !visited[dx][dy] {
				volume += max(0, cell.height-heightMap[dx][dy])

				heap.Push(&Cell{max(cell.height, heightMap[dx][dy]), dx, dy})
				visited[dx][dy] = true
			}
		}
	}

	return volume
}
