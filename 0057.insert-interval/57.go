/**
 * Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
 * You may assume that the intervals were initially sorted according to their start times.
 *
 * Example 1:
 *
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 *
 * Example 2:
 *
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
 * Example 3:
 *
 * Input: intervals = [], newInterval = [5,7]
 * Output: [[5,7]]
 *
 * Example 4:
 *
 * Input: intervals = [[1,5]], newInterval = [2,3]
 * Output: [[1,5]]
 *
 * Example 5:
 *
 * Input: intervals = [[1,5]], newInterval = [2,7]
 * Output: [[1,7]]
 *
 *
 * Constraints:
 *
 * 	0 <= intervals.length <= 10^4
 * 	intervals[i].length == 2
 * 	0 <= intervals[i][0] <= intervals[i][1] <= 10^5
 * 	intervals is sorted by intervals[i][0] in ascending order.
 * 	newInterval.length == 2
 * 	0 <= newInterval[0] <= newInterval[1] <= 10^5
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Collect the intervals strictly left or right of the new interval, then merge the new one with
// the middle ones (if any) before inserting it between left and right ones.
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	start := newInterval[0]
	end := newInterval[1]

	left := make([][]int, 0)
	right := make([][]int, 0)

	for _, interval := range intervals {
		if interval[1] < start {
			left = append(left, interval)
		} else {
			start = min(interval[0], start)
		}

		if interval[0] > end {
			right = append(right, interval)
		} else {
			end = max(interval[1], end)
		}
	}

	return append(left, append([][]int{{start, end}}, right...)...)
}
