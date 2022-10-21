/**
 * Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.
 *
 * If there is no non-empty subarray with sum at least K, return -1.
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
 * Input: A = <span id="example-input-1-1">[1]</span>, K = <span id="example-input-1-2">1</span>
 * Output: <span id="example-output-1">1</span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: A = <span id="example-input-2-1">[1,2]</span>, K = <span id="example-input-2-2">4</span>
 * Output: <span id="example-output-2">-1</span>
 *
 *
 * <div>
 * Example 3:
 *
 *
 * Input: A = <span id="example-input-3-1">[2,-1,2]</span>, K = <span id="example-input-3-2">3</span>
 * Output: <span id="example-output-3">3</span>
 *
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	1 <= A.length <= 50000
 * 	-10 ^ 5 <= A[i] <= 10 ^ 5
 * 	1 <= K <= 10 ^ 9
 * </ol>
 * </div>
 * </div>
 * </div>
 *
 */

package leetcode

type Queue struct {
	items []int
}

func (q *Queue) Push(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) PeekFront() int {
	return q.items[0]
}

func (q *Queue) PeekBack() int {
	return q.items[len(q.items)-1]
}

func (q *Queue) Poll() int {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Pop() int {
	val := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return val
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestSubarray(A []int, K int) int {
	queue := &Queue{}
	minVal := len(A) + 1

	prefixSum := make([]int, len(A)+1)
	for i := range A {
		prefixSum[i+1] = prefixSum[i] + A[i]
	}

	// iterate over the prefix sum array
	for i := 0; i < len(A)+1; i++ {
		// deal with negative numbers
		// basically, it will sometimes delete all elements from the queue
		// input = [12, -50, 50]
		// []
		// [0]
		// [] -> delete all elements
		// [2] -> i=3, peekFront=2 (NOT 0 as it would have been)

		for !queue.Empty() && prefixSum[i] <= prefixSum[queue.PeekBack()] {
			queue.Pop()
		}

		// try to minimise the length of the contiguous subarray
		for !queue.Empty() && prefixSum[i]-prefixSum[queue.PeekFront()] >= K {
			minVal = min(minVal, i-queue.Poll())
		}

		queue.Push(i)
	}

	if minVal == len(A)+1 {
		return -1
	}
	return minVal
}
