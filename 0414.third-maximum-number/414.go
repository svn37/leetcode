/**
 * Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
 *
 * Example 1:<br />
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:<br />
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
 *
 *
 *
 * Example 3:<br />
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */

package leetcode

import "math"

func thirdMax(nums []int) int {
	max3 := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}
	for i := range nums {
		if nums[i] == max3[0] || nums[i] == max3[1] || nums[i] == max3[2] {
			continue
		}

		if nums[i] > max3[0] {
			max3[0], max3[1], max3[2] = nums[i], max3[0], max3[1]
		} else if nums[i] > max3[1] {
			max3[1], max3[2] = nums[i], max3[1]
		} else if nums[i] > max3[2] {
			max3[2] = nums[i]
		}
	}
	// if it does not exist, return the maximum number
	if max3[2] == math.MinInt64 {
		return max3[0]
	}
	return max3[2]
}
