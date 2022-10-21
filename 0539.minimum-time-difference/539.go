/**
 * Given a list of 24-hour clock time points in "HH:MM" format, return the minimum minutes difference between any two time-points in the list.
 *
 * Example 1:
 * Input: timePoints = ["23:59","00:00"]
 * Output: 1
 * Example 2:
 * Input: timePoints = ["00:00","23:59","00:00"]
 * Output: 0
 *
 * Constraints:
 *
 * 	2 <= timePoints <= 2 * 10^4
 * 	timePoints[i] is in the format "HH:MM".
 *
 */

package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func convertToTimestamp(s string) int {
	hours := int(s[0]-'0')*10 + int(s[1]-'0')
	minutes := int(s[3]-'0')*10 + int(s[4]-'0')
	return hours*60 + minutes
}

func findMinDifference(timePoints []string) int {
	time := make([]bool, 24*60+1)
	for i := range timePoints {
		idx := convertToTimestamp(timePoints[i])
		// met the same timestamp twice
		if time[idx] {
			return 0
		}
		time[idx] = true
	}

	minVal := 24 * 60
	first, prev := -1, -1
	for i := range time {
		if time[i] {
			if prev < 0 {
				first = i
			} else {
				minVal = min(minVal, i-prev)
			}
			prev = i
		}
	}
	// check wrap around time diff
	if first >= 0 {
		minVal = min(minVal, first+24*60-prev)
	}
	return minVal
}
