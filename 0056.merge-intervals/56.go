/**
 * Given a collection of intervals, merge all overlapping intervals.
 * Example 1:
 *
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
 *
 * Example 2:
 *
 * Input: intervals = [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 * NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
 *
 * Constraints:
 *
 * 	intervals[i][0] <= intervals[i][1]
 *
 */

package leetcode

import (
	"math"
	"math/rand"
)

func qsort(a [][]int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Intn(len(a) - 1)

	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if a[right][0] > a[i][0] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	// sort intervals by start value
	qsort(intervals)

	res := [][]int{}
	maxVal := math.MinInt64
	for _, interval := range intervals {
		if interval[0] <= maxVal {
			// overlapping interval, update last interval end value
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		} else {
			res = append(res, interval)
		}

		if interval[1] > maxVal {
			maxVal = interval[1]
		}
	}

	return res
}
