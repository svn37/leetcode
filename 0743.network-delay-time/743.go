/**
 * There are N network nodes, labelled 1 to N.
 *
 * Given times, a list of travel times as directed edges times[i] = (u, v, w), where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.
 *
 * Now, we send a signal from a certain node K. How long will it take for all nodes to receive the signal? If it is impossible, return -1.
 *
 *
 *
 * Example 1:
 *
 * <img alt="" src="https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png" style="width: 200px; height: 220px;" />
 *
 *
 * Input: times = <span id="example-input-1-1">[[2,1,1],[2,3,1],[3,4,1]]</span>, N = <span id="example-input-1-2">4</span>, K = <span id="example-input-1-3">2</span>
 * Output: <span id="example-output-1">2</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	N will be in the range [1, 100].
 * 	K will be in the range [1, N].
 * 	The length of times will be in the range [1, 6000].
 * 	All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 0 <= w <= 100.
 * </ol>
 *
 */

package leetcode

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Queue struct {
	storage []int
}

func (q *Queue) Enqueue(i int) {
	q.storage = append(q.storage, i)
}

func (q *Queue) Dequeue() int {
	i := q.storage[0]
	q.storage = q.storage[1:]
	return i
}

func (q *Queue) Empty() bool {
	return len(q.storage) == 0
}

type wnode struct {
	v int
	w int
}

func buildGraph(wedgeList [][]int, N int) [][]wnode {
	g := make([][]wnode, N)
	for i := range wedgeList {
		g[wedgeList[i][0]-1] = append(g[wedgeList[i][0]-1], wnode{
			v: wedgeList[i][1] - 1,
			w: wedgeList[i][2],
		})
	}
	return g
}

// SPFA algorithm
func networkDelayTime(times [][]int, N int, K int) int {
	dist := make([]int, N)
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	K--
	dist[K] = 0

	q := &Queue{}
	q.Enqueue(K)

	g := buildGraph(times, N)

	for !q.Empty() {
		u := q.Dequeue()
		for i := range g[u] {
			if dist[g[u][i].v] > g[u][i].w+dist[u] {
				dist[g[u][i].v] = g[u][i].w + dist[u]
				q.Enqueue(g[u][i].v)
			}
		}
	}

	maxDelay := math.MinInt64
	for i := range dist {
		maxDelay = max(maxDelay, dist[i])
	}
	if maxDelay == math.MaxInt64 {
		return -1
	}
	return maxDelay
}
