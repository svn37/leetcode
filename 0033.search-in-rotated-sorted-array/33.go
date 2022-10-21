/**
 * You are given an integer array nums sorted in ascending order, and an integer target.
 * Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * If target is found in the array return its index, otherwise, return -1.
 *
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 *
 * Constraints:
 *
 * 	1 <= nums.length <= 5000
 * 	-10^4 <= nums[i] <= 10^4
 * 	All values of nums are unique.
 * 	nums is guranteed to be rotated at some pivot.
 * 	-10^4 <= target <= 10^4
 *
 */

package leetcode

import "math"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		var num int
		// nums[mid] and target are on the same sorted side
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
			// otherwise we don't need this side, use infinity for comparison
		} else if target < nums[0] {
			num = math.MinInt64
		} else {
			num = math.MaxInt64
		}

		if num == target {
			return mid
		} else if num < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
